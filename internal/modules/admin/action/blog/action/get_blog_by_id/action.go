package get_blog_by_id

import (
	"medblogers_base/internal/modules/admin/action/blog/action/get_blog_by_id/dal"
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
