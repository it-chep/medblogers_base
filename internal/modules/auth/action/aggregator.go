package action

import (
	"medblogers_base/internal/modules/auth/action/check_permissions"
	"medblogers_base/internal/modules/auth/action/get_user_info"
	"medblogers_base/internal/modules/auth/action/register"
)

type Aggregator struct {
	GetUserInfo      *get_user_info.Action
	CheckPermissions *check_permissions.Action
	Register         *register.Action
}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}
