package repomodel

import (
	"database/sql"
	"time"
)

// UserID represents a user with a unique ID.
type UserID struct {
	UserID int64
}

// UserInfo represents a user information.
type UserInfo struct {
	Name  string
	Email string
	Role  int64
}

// UpdateInfo represents a information for user update.
type UpdateInfo struct {
	ID    int64
	Name  string
	Email string
	Role  int64
}

// User represents all information about user.
type User struct {
	ID              int64
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            int64
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}
