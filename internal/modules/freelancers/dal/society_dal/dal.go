package society_dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/freelancers/dal/society_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Repository соц сети
type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с соцсетями
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetCitiesWithFreelancersCount получение списка городов с количеством фрилансеров
func (r Repository) GetCitiesWithFreelancersCount(ctx context.Context) ([]*social_network.SocialNetwork, error) {
	logger.Message(ctx, "[DAL] Запрос соцсетей")
	sql := `
		with active_freelancers_in_media as (
			select
				fc.social_network_id,
				fc.freelancer_id
			from freelancer_social_networks_m2m fc
					 join freelancer f on fc.freelancer_id = f.id
			where f.is_active = true
		)
		select
			n.id as id,
			n.name as name,
			count(distinct af.freelancer_id) as freelancers_count
		from social_networks n
				 left join active_freelancers_in_media af on n.id = af.social_network_id
		group by n.id, n.name
		having count(distinct af.freelancer_id) > 0
		order by n.name;
	`

	var socialDao []dao.SocialNetworkWithFreelancersCount
	if err := pgxscan.Select(ctx, r.db, &socialDao, sql); err != nil {
		return nil, err
	}

	networks := make([]*social_network.SocialNetwork, 0, len(socialDao))
	for _, sDao := range socialDao {
		networks = append(networks, sDao.ToDomain())
	}

	return networks, nil
}

// GetAllNetworks все соц сети
func (r Repository) GetAllNetworks(ctx context.Context) ([]*social_network.SocialNetwork, error) {
	sql := `
		select c.id   as id,
			   c.name as name
		from social_networks c
		group by c.id, c.name
	`

	var socialDao []dao.SocialNetworkWithFreelancersCount
	if err := pgxscan.Select(ctx, r.db, &socialDao, sql); err != nil {
		return nil, err
	}

	networks := make([]*social_network.SocialNetwork, 0, len(socialDao))
	for _, sDao := range socialDao {
		networks = append(networks, sDao.ToDomain())
	}

	return networks, nil
}
