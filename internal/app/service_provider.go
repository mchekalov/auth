package app

import (
	"context"
	"log"

	"github.com/mchekalov/auth/internal/api"
	"github.com/mchekalov/auth/internal/config"
	"github.com/mchekalov/auth/internal/repository"
	chatrepository "github.com/mchekalov/auth/internal/repository/auth"
	"github.com/mchekalov/auth/internal/service"
	chatservice "github.com/mchekalov/auth/internal/service/auth"
	"github.com/mchekalov/platform_common/pkg/closer"
	"github.com/mchekalov/platform_common/pkg/db"
	"github.com/mchekalov/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	dbClient   db.Client

	authRepository repository.AuthRepository

	authService service.AuthService

	authImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DsnString())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = chatrepository.NewRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = chatservice.NewService(
			s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *api.Implementation {
	if s.authImpl == nil {
		s.authImpl = api.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
