package v1

import (
	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/auth"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

type Implementation struct {
	desc.UnimplementedDoctorAdminServiceServer

	admin *admin.Module
	auth  *auth.Module
}

func New(admin *admin.Module, auth *auth.Module) *Implementation {
	return &Implementation{
		admin: admin,
		auth:  auth,
	}
}
