package auth

import (
	"context"

	"github.com/mchekalov/auth/internal/model"
	repository "github.com/mchekalov/auth/internal/repository"
	"github.com/mchekalov/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

// NewService creates instance of service layer
func NewService(authRepository repository.AuthRepository) service.AuthService {
	return &serv{authRepository: authRepository}
}

func (s *serv) Create(ctx context.Context, in *model.UserInfo) (*model.UserID, error) {
	output, err := s.authRepository.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *serv) Get(ctx context.Context, userID *model.UserID) (*model.User, error) {
	output, err := s.authRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *serv) Update(ctx context.Context, in *model.UpdateInfo) error {
	err := s.authRepository.Update(ctx, in)
	if err != nil {
		return err
	}

	return nil

}

func (s *serv) Delete(ctx context.Context, in *model.UserID) error {
	err := s.authRepository.Delete(ctx, in)
	if err != nil {
		return err
	}

	return nil

}
