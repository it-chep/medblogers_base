package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/doctor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/social_network"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/converter"
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

func (r *Repository) GetPriceList(ctx context.Context, freelancerID int64) ([]dto.PriceList, error) {
	sql := `
		select id, name, price from freelancers_price_list where freelancer_id = $1
	`

	var priceList []dao.PriceListDao
	err := pgxscan.Select(ctx, r.db, &priceList, sql, freelancerID)
	if err != nil {
		return nil, err
	}

	return lo.Map(priceList, func(item dao.PriceListDao, _ int) dto.PriceList {
		return dto.PriceList{
			ID:     item.ID,
			Name:   item.Name,
			Amount: converter.NumericToDecimal(item.Price).String(),
		}
	}), nil
}

func (r *Repository) GetRecommendations(ctx context.Context, freelancerID int64) ([]int64, error) {
	sql := `
		select doctor_id from freelancer_recommendation where freelancer_id = $1
	`
	// todo cортировочку
	var recommendations []int64
	err := pgxscan.Select(ctx, r.db, &recommendations, sql, freelancerID)
	if err != nil {
		return nil, err
	}
	return recommendations, nil
}

func (r *Repository) GetRecommendationInfoByIDs(ctx context.Context, doctorIDs []int64) ([]*doctor.Doctor, error) {
	sql := `
		select id, name from docstar_site_doctor where id = any($1)
	`

	var recommendations dao.Recommendations
	err := pgxscan.Select(ctx, r.db, &recommendations, sql, doctorIDs)
	if err != nil {
		return nil, err
	}

	return recommendations.ToDomain(), nil
}

func (r *Repository) GetCities(ctx context.Context, freelancerID int64) ([]*city.City, error) {
	sql := `
		select c.id, c.name 
		from freelancer_city_m2m m2m 
		    join freelancers_city c on m2m.city_id = c.id
		where m2m.freelancer_id = $1
	`

	var cities dao.CitiesDAO
	err := pgxscan.Select(ctx, r.db, &cities, sql, freelancerID)
	if err != nil {
		return nil, err
	}
	return cities.ToDomain(), nil
}

func (r *Repository) GetSpecialities(ctx context.Context, freelancerID int64) ([]*speciality.Speciality, error) {
	sql := `
		select s.id as "id", s.name as "name" 
		from freelancer_speciality_m2m m2m 
		    join freelancers_speciality s on m2m.speciality_id = s.id
		where m2m.freelancer_id = $1
	`

	var specialities dao.SpecialitiesDAO
	err := pgxscan.Select(ctx, r.db, &specialities, sql, freelancerID)
	if err != nil {
		return nil, err
	}
	return specialities.ToDomain(), nil
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
