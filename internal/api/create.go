package api

import (
	"context"

	desc "github.com/mchekalov/auth/pkg/user_v1"

	"github.com/mchekalov/auth/internal/converter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create creates a new user
func (i *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	output, err := i.authService.Create(ctx, converter.FromAPIToServUserInfo(request))
	if err != nil {

		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return converter.FromServToAPIID(output), nil
}
