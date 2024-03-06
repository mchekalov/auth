package converter

import (
	"github.com/mchekalov/auth/internal/model"
	modelRepo "github.com/mchekalov/auth/internal/repository/auth/model"
)

// FromRepoToServUserID converts a repo model to a service model user id.
func FromRepoToServUserID(id *modelRepo.UserID) *model.UserID {
	return &model.UserID{
		UserID: id.UserID,
	}
}

// FromServToRepoUserID converts a service model to a repo model user id.
func FromServToRepoUserID(id *model.UserID) *modelRepo.UserID {
	return &modelRepo.UserID{
		UserID: id.UserID,
	}
}

// FromServToRepoUserInfo converts from service to repo layer user info model.
func FromServToRepoUserInfo(in *model.UserInfo) *modelRepo.UserInfo {
	return &modelRepo.UserInfo{
		Name:  in.Name,
		Email: in.Email,
		Role:  in.Role,
	}
}

// FromRepoToServUser converts a repository User model to a service user model.
func FromRepoToServUser(u *modelRepo.User) *model.User {
	return &model.User{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		Role:            u.Role,
		Password:        u.Password,
		PasswordConfirm: u.PasswordConfirm,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}

// FromServToRepoInfo converts a service info model to a repo model.
func FromServToRepoInfo(m *model.UpdateInfo) *modelRepo.UpdateInfo {
	return &modelRepo.UpdateInfo{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
		Role:  m.Role,
	}
}
