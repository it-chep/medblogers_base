package get_blogs

import (
	"medblogers_base/internal/modules/admin/action/blog/action/get_blogs/dal"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	dal *dal.Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}
