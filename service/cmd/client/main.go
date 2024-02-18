package main

import (
	"context"
	"fmt"
	"log"
	"time"

	desc "github.com/mchekalov/auth/service/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server %v", err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Printf("Error when closing: %v", err)
		}
	}()

	c := desc.NewUserV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: 333})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf("User:\n %v", r.GetUser())
}
