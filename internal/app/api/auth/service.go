package auth

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/auth"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
)

type Implementation struct {
	desc.UnimplementedAuthServiceServer

	auth   *auth.Module
	config *config.Config
}

// NewAuthService return new instance of Implementation.
func NewAuthService(auth *auth.Module, config *config.Config) *Implementation {
	return &Implementation{
		auth:   auth,
		config: config,
	}
}
