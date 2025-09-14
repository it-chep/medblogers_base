package dal

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
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

func (r *Repository) CreateFreelancer(ctx context.Context, createDTO dto.CreateRequest) (int64, error) {
	sql := `
		insert into freelancer (email,
						slug,
						name,
						is_worked_with_doctors,
						is_active,
						tg_username,
						portfolio_link,
						speciality_id,
						city_id,
						price_category)
		values ($1, $2, $3, $4, false, $5, $6, $7, $8, $9)
		returning id;
	`

	args := []interface{}{
		createDTO.Email,
		createDTO.Slug,
		createDTO.Name,
		createDTO.HasExperienceWithDoctors,
		createDTO.TgUsername,
		createDTO.PortfolioLink,
		createDTO.MainSpecialityID,
		createDTO.MainCityID,
		createDTO.PriceCategory,
	}

	var freelancerID int64
	err := r.db.QueryRow(ctx, sql, args...).Scan(&freelancerID)
	return freelancerID, err
}

func (r *Repository) CreateSocialNetworks(ctx context.Context, freelancerID int64, networkIDs []int64) error {
	if len(networkIDs) == 0 {
		return nil
	}

	sql := `
	insert into freelancer_social_networks_m2m (freelancer_id, social_network_id)
		select $1, unnest($2::bigint[])
		on conflict (freelancer_id, social_network_id) do nothing`

	_, err := r.db.Exec(ctx, sql, freelancerID, networkIDs)
	return err
}

func (r *Repository) CreateAdditionalCities(ctx context.Context, freelancerID int64, citiesIDs []int64) error {
	if len(citiesIDs) == 0 {
		return nil
	}

	sql := `
	insert into freelancer_city_m2m (freelancer_id, city_id)
		select $1, unnest($2::bigint[])
		on conflict (freelancer_id, city_id) do nothing`

	_, err := r.db.Exec(ctx, sql, freelancerID, citiesIDs)
	return err
}

func (r *Repository) CreateAdditionalSpecialities(ctx context.Context, freelancerID int64, specialitiesIDs []int64) error {
	if len(specialitiesIDs) == 0 {
		return nil
	}

	sql := `
	insert into freelancer_speciality_m2m (freelancer_id, speciality_id)
		select $1, unnest($2::bigint[])
		on conflict (freelancer_id, speciality_id) do nothing`

	_, err := r.db.Exec(ctx, sql, freelancerID, specialitiesIDs)
	return err
}

func (r *Repository) CreatePriceList(ctx context.Context, freelancerID int64, priceList dto.PriceList) error {
	if len(priceList) == 0 {
		return nil
	}

	names := make([]string, len(priceList))
	prices := make([]int64, len(priceList))

	for i, item := range priceList {
		names[i] = item.Name
		prices[i] = item.Price
	}

	sql := `
	insert into freelancers_price_list (freelancer_id, name, price)
         select $1, name, price
         from unnest($2::text[], $3::bigint[]) as t(name, price)
         on conflict (freelancer_id, name, price) do nothing`

	_, err := r.db.Exec(ctx, sql, freelancerID, names, prices)
	return err
}
