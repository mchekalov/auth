package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/mchekalov/auth/internal/model"
	"github.com/mchekalov/auth/internal/repository"
	repoMocks "github.com/mchekalov/auth/internal/repository/mocks"
	"github.com/mchekalov/auth/internal/service/auth"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository

	type args struct {
		ctx context.Context
		req *model.UserID
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("repo error")

		req = &model.UserID{
			UserID: id,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               error
		authRepositoryMock authRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: serviceErr,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, req).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authRepositoryMock := tt.authRepositoryMock(mc)
			service := auth.NewService(authRepositoryMock)

			err := service.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, err)
			require.Equal(t, tt.want, err)
		})
	}
}
