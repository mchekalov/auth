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

func TestUpdate(t *testing.T) {
	t.Parallel()
	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository

	type args struct {
		ctx context.Context
		req *model.UpdateInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		idres = gofakeit.IntRange(0, 2)

		serviceErr = fmt.Errorf("repo error")

		req = &model.UpdateInfo{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  int64(idres),
		}
	)

	tests := []struct {
		name               string
		args               args
		err                error
		authRepositoryMock authRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: serviceErr,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authRepositoryMock := tt.authRepositoryMock(mc)
			service := auth.NewService(authRepositoryMock)

			err := service.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
