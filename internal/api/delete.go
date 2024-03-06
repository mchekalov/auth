package api

import (
	"context"

	"github.com/mchekalov/auth/internal/converter"
	desc "github.com/mchekalov/auth/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a user in API layer
func (i *Implementation) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.authService.Delete(ctx, converter.FromAPIToServID(request))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	return &emptypb.Empty{}, nil
}
