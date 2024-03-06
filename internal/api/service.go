package api

import (
	"github.com/mchekalov/auth/internal/service"
	desc "github.com/mchekalov/auth/pkg/user_v1"
)

// Implementation represents the implementation of the auth API server.
type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

// NewImplementation creates a new instance of the auth API server implementation.
func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
