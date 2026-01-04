package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/social_network"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetNetworks(ctx context.Context) (social_network.Networks, error) {
	sql := `
		select id, name from social_networks
		`

	var networks dao.NetworksDAO
	err := pgxscan.Select(ctx, r.db, &networks, sql)
	if err != nil {
		return nil, err
	}
	return networks.ToDomain(), err
}
