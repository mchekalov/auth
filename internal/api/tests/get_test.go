package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/mchekalov/auth/internal/api"
	"github.com/mchekalov/auth/internal/model"
	"github.com/mchekalov/auth/internal/service"
	serviceMocks "github.com/mchekalov/auth/internal/service/mocks"
	desc "github.com/mchekalov/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		pass      = gofakeit.Password(true, true, true, true, false, 10)
		idres     = gofakeit.IntRange(0, 2)
		role      = desc.Role(int32(idres))
		createdat = gofakeit.Date()
		updatedat = gofakeit.Date()

		serviceErr = status.Error(codes.Internal, "failed to delete chat")

		req = &desc.GetRequest{
			Id: id,
		}

		info = &model.UserID{
			UserID: id,
		}

		userID = &model.User{
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

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name:            name,
					Email:           email,
					Password:        pass,
					PasswordConfirm: pass,
					Role:            role,
				},
				CreatedAt: timestamppb.New(createdat),
				UpdatedAt: timestamppb.New(updatedat),
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, info).Return(userID, nil)
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
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, info).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := api.NewImplementation(authServiceMock)

			newID, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
