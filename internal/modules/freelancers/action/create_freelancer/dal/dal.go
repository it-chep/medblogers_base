package dal

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/jackc/pgtype"
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
						is_active,
						tg_username,
						portfolio_link,
						speciality_id,
						city_id,
						price_category,
		                agency_representative,
		                has_med_education,
		                start_working_date)
		values ($1, $2, $3, false, $4, $5, $6, $7, $8, $9, $10, $11)
		returning id;
	`

	args := []interface{}{
		createDTO.Email,
		createDTO.Slug,
		createDTO.Name,
		createDTO.TgUsername,
		createDTO.PortfolioLink,
		createDTO.MainSpecialityID,
		createDTO.MainCityID,
		createDTO.PriceCategory,
		createDTO.AgencyRepresentative,
		createDTO.HasMedEducation,
		createDTO.StartWorkingExperience,
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

	names := make([]string, 0, len(priceList))
	prices := make([]int64, 0, len(priceList))
	priceToElements := make([]pgtype.Int8, 0, len(priceList))

	for _, item := range priceList {
		names = append(names, item.Name)
		prices = append(prices, item.Price)

		if item.PriceTo == nil {
			priceToElements = append(priceToElements, pgtype.Int8{Status: pgtype.Null})
			continue
		}

		priceToElements = append(priceToElements, pgtype.Int8{
			Int:    *item.PriceTo,
			Status: pgtype.Present,
		})
	}

	priceToArray := pgtype.Int8Array{
		Elements:   priceToElements,
		Dimensions: []pgtype.ArrayDimension{{Length: int32(len(priceToElements)), LowerBound: 1}},
		Status:     pgtype.Present,
	}

	sql := `
	insert into freelancers_price_list (freelancer_id, name, price, price_to, search_vector)
	select $1,
	       t.name,
	       t.price,
	       t.price_to,
	       to_tsvector('russian', coalesce(t.name, ''))
	from unnest($2::text[], $3::bigint[], $4::bigint[]) as t(name, price, price_to)
	`

	_, err := r.db.Exec(ctx, sql, freelancerID, names, prices, &priceToArray)
	return err
}
