package action

import (
	"medblogers_base/internal/modules/auth/action/check_permissions"
	"medblogers_base/internal/modules/auth/action/get_user_info"
	"medblogers_base/internal/modules/auth/action/register"
	"medblogers_base/internal/pkg/postgres"
)

type Aggregator struct {
	GetUserInfo      *get_user_info.Action
	CheckPermissions *check_permissions.Action
	Register         *register.Action
}

func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		GetUserInfo:      get_user_info.New(pool),
		Register:         register.New(pool),
		CheckPermissions: check_permissions.New(pool),
	}
}
