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
func (r *Repository) GetNetworks(ctx context.Context, freelancerID int64) ([]*social_network.SocialNetwork, error) {
	sql := `
		select s.id, s.name, s.slug
		from freelancer_social_networks_m2m m2m 
		    join social_networks s on m2m.social_network_id = s.id
		where m2m.freelancer_id = $1
	`

	var networks dao.NetworksDAO
	err := pgxscan.Select(ctx, r.db, &networks, sql, freelancerID)
	if err != nil {
		return nil, err
	}
	return networks.ToDomain(), nil
}
