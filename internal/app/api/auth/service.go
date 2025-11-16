package auth

import (
	"medblogers_base/internal/modules/auth"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
)

type Implementation struct {
	desc.UnimplementedAuthServiceServer

	auth *auth.Module
	//config
}

// NewAuthService return new instance of Implementation.
func NewAuthService() *Implementation {
	return &Implementation{}
}
