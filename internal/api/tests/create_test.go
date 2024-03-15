package tests

import (
	"context"
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
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		pass  = gofakeit.Password(true, true, true, true, false, 10)
		idres = gofakeit.IntRange(0, 2)
		role  = desc.Role(int32(idres))

		serviceErr = status.Error(codes.Internal, "failed to create user")

		req = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Name:            name,
				Email:           email,
				Password:        pass,
				PasswordConfirm: pass,
				Role:            role,
			},
		}

		info = &model.UserInfo{
			Name:  name,
			Email: email,
			Role:  int64(idres),
		}

		userID = &model.UserID{
			UserID: id,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, info).Return(userID, nil)
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
				mock.CreateMock.Expect(ctx, info).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := api.NewImplementation(authServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
