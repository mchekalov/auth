package auth

import (
	"context"
	"fmt"

	"github.com/mchekalov/auth/internal/client/db"
	"github.com/mchekalov/auth/internal/model"
	"github.com/mchekalov/auth/internal/repository"
	"github.com/mchekalov/auth/internal/repository/auth/converter"
	modelRepo "github.com/mchekalov/auth/internal/repository/auth/model"

	"github.com/Masterminds/squirrel"
)

const (
	tableName       = "auth_user"
	userIDColumn    = "id"
	userNameColumn  = "user_name"
	emailColumn     = "email"
	userRoleColumn  = "user_role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
	sq squirrel.StatementBuilderType
}

// NewRepository create new instance for repo object
func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}

func (r *repo) Create(ctx context.Context, ui *model.UserInfo) (*model.UserID, error) {
	userInfo := converter.FromServToRepoUserInfo(ui)

	builder := r.sq.Insert(tableName).
		Columns(userNameColumn, emailColumn, userRoleColumn).
		Values(userInfo.Name, userInfo.Email, userInfo.Role).
		Suffix(fmt.Sprintf("RETURNING %v", userIDColumn))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var id modelRepo.UserID
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id.UserID)
	if err != nil {
		return nil, err
	}

	return converter.FromRepoToServUserID(&id), nil
}

func (r *repo) Get(ctx context.Context, id *model.UserID) (*model.User, error) {
	userID := converter.FromServToRepoUserID(id)

	builder := r.sq.Select(userIDColumn, userNameColumn, emailColumn, userRoleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(squirrel.Eq{userIDColumn: userID.UserID}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.FromRepoToServUser(&user), nil
}

func (r *repo) Update(ctx context.Context, info *model.UpdateInfo) error {
	infoRepo := converter.FromServToRepoInfo(info)

	builder := r.sq.Update(tableName).
		Set(userNameColumn, infoRepo.Name).
		Set(emailColumn, infoRepo.Email).
		Set(updatedAtColumn, "CURRENT_DATE").
		Where(squirrel.Eq{userIDColumn: infoRepo.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Update",
		QueryRaw: query,
	}

	r.db.DB().QueryRowContext(ctx, q, args...)

	return nil
}

func (r *repo) Delete(ctx context.Context, id *model.UserID) error {
	userID := converter.FromServToRepoUserID(id)

	builder := r.sq.Delete(tableName).
		Where(squirrel.Eq{userIDColumn: userID.UserID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	r.db.DB().QueryRowContext(ctx, q, args...)

	return nil
}
