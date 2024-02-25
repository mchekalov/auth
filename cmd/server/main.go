package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mchekalov/auth/config"
	desc "github.com/mchekalov/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	pgx *pgx.Conn
}

func (s *server) Create(ctx context.Context, in *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create new User %v", in.Info.Name)
	var userID int64

	row := s.pgx.QueryRow(ctx, "INSERT INTO auth_user (user_name, email, user_role) VALUES ($1, $2, $3) RETURNING id",
		in.Info.Name, in.Info.Email, in.Info.Role)

	err := row.Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: userID,
	}, nil
}

func (s *server) Get(ctx context.Context, in *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User ID %v", in.GetId())

	row := s.pgx.QueryRow(ctx, "SELECT id, user_name, email, user_role, created_at, updated_at FROM auth_user WHERE id=$1", in.GetId())

	var id int64
	var name, email string
	var role int32
	var createdAt time.Time
	var updatedAt sql.NullTime

	err := row.Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id: id,
			Info: &desc.UserInfo{
				Name:  name,
				Email: email,
				Role:  desc.Role(role),
			},
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtTime,
		},
	}, nil
}

func (s *server) Update(ctx context.Context, in *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user %v", in.GetWrap())

	_, err := s.pgx.Exec(ctx, "UPDATE auth_user SET user_name=$1, email=$2, updated_at=CURRENT_DATE WHERE id=$3",
		in.Wrap.Name, in.Wrap.Email, in.Wrap.Id)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *server) Delete(ctx context.Context, in *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user %v", in.Id)

	_, err := s.pgx.Exec(ctx, "DELETE FROM auth_user WHERE id=$1", in.Id)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func main() {
	ctx := context.Background()

	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("fatal to listen %v", err)
	}

	// get postgres db connect

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load postgres config: %v", err)
	}

	con, err := pgx.Connect(ctx, pgConfig.DsnString())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err = con.Close(ctx); err != nil {
			log.Printf("Error when closing connection: %v", err)
		}

	}()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{pgx: con})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve; %v", err)
	}
}
