package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/samber/lo"

	"medblogers_base/internal/modules/admin/entities/promo_offers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetSocialNetworks(ctx context.Context) (dictionary.SocialNetworks, error) {
	sql := `
		select id, name, slug
		from social_networks
		order by id desc
	`

	var rows []dao.SocialNetworkDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql); err != nil {
		return nil, err
	}

	return lo.Map(rows, func(item dao.SocialNetworkDAO, _ int) *dictionary.SocialNetwork {
		return item.ToDomain()
	}), nil
}
