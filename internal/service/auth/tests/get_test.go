package tests

import (
	"context"
	"database/sql"
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

func TestGet(t *testing.T) {
	t.Parallel()
	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository

	type args struct {
		ctx context.Context
		req *model.UserID
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		pass      = gofakeit.Password(true, true, true, true, false, 10)
		idres     = gofakeit.IntRange(0, 2)
		createdat = gofakeit.Date()
		updatedat = gofakeit.Date()

		serviceErr = fmt.Errorf("repo error")

		req = &model.UserID{
			UserID: id,
		}

		res = &model.User{
			ID:              id,
			Name:            name,
			Email:           email,
			Password:        pass,
			PasswordConfirm: pass,
			Role:            int64(idres),
			CreatedAt:       createdat,
			UpdatedAt: sql.NullTime{
				Time:  updatedat,
				Valid: true,
			},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		authRepositoryMock authRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.GetMock.Expect(ctx, req).Return(res, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.GetMock.Expect(ctx, req).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authRepositoryMock := tt.authRepositoryMock(mc)
			service := auth.NewService(authRepositoryMock)

			newID, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
