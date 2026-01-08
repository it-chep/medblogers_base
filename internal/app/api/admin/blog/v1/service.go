package v1

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/auth"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"
)

type Implementation struct {
	desc.UnimplementedBlogsAdminServiceServer

	admin  *admin.Module
	auth   *auth.Module
	config *config.Config
}

// NewAdminService return new instance of Implementation.
func NewAdminService(admin *admin.Module, auth *auth.Module, cfg *config.Config) *Implementation {
	return &Implementation{
		admin:  admin,
		auth:   auth,
		config: cfg,
	}
}
