package update

import "medblogers_base/internal/pkg/postgres"

type Action struct {
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{}
}
