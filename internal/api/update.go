package api

import (
	"context"

	"github.com/mchekalov/auth/internal/converter"
	desc "github.com/mchekalov/auth/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates a user in API layer
func (i *Implementation) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.authService.Update(ctx, converter.FromAPIToServUpdate(request))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &emptypb.Empty{}, nil
}
