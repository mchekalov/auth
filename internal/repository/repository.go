package repository

import (
	"context"

	"github.com/mchekalov/auth/internal/model"
)

// AuthRepository defines an interface for interacting with the repository layer
// to perform operations related to auth entities.
type AuthRepository interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (*model.UserID, error)
	Get(ctx context.Context, userID *model.UserID) (*model.User, error)
	Update(ctx context.Context, updateInfo *model.UpdateInfo) error
	Delete(ctx context.Context, userID *model.UserID) error
}
