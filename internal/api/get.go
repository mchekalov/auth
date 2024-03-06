package api

import (
	"context"

	"github.com/mchekalov/auth/internal/converter"

	desc "github.com/mchekalov/auth/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get get a user full information in API layer.
func (i *Implementation) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.authService.Get(ctx, converter.FromAPIToServIDGet(request))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	return converter.FromServToAPIUser(user), nil
}
