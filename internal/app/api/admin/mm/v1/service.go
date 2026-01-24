package v1

import (
	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/auth"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mm/v1"
)

type Implementation struct {
	desc.UnimplementedMMAdminServiceServer

	admin *admin.Module
	auth  *auth.Module
}

func NewMMService(admin *admin.Module, auth *auth.Module) *Implementation {
	return &Implementation{
		admin: admin,
		auth:  auth,
	}
}
