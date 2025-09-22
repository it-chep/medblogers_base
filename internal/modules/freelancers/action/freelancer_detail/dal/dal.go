package dal

import (
	"context"
	cityDAO "medblogers_base/internal/modules/freelancers/dal/city_dal/dao"
	freelancerDao "medblogers_base/internal/modules/freelancers/dal/freelancer_dal/dao"
	priceListDao "medblogers_base/internal/modules/freelancers/dal/price_list/dao"
	socialDao "medblogers_base/internal/modules/freelancers/dal/society_dal/dao"
	specialityDAO "medblogers_base/internal/modules/freelancers/dal/speciality_dal/dao"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/modules/freelancers/domain/price_list"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
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

// GetFreelancerInfo детальная информация о фрилансере
func (r *Repository) GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error) {
	sql := `
		select id, slug, name, is_worked_with_doctors, tg_username, portfolio_link, speciality_id, city_id, price_category, s3_image, has_command, start_working_date
		    from freelancer
		where slug = $1 and is_active = true
	`

	var fDao freelancerDao.FreelancerDetail
	err := pgxscan.Get(ctx, r.db, &fDao, sql, slug)
	if err != nil {
		return nil, err
	}
	return fDao.ToDomain(), nil
}

// GetPriceList получение прайс-листа фрилансера
func (r *Repository) GetPriceList(ctx context.Context, freelancerID int64) (price_list.PriceList, error) {
	sql := `
		select name, price from freelancers_price_list where freelancer_id = $1
	`

	var priceList priceListDao.PriceList
	err := pgxscan.Select(ctx, r.db, &priceList, sql, freelancerID)
	if err != nil {
		return nil, err
	}

	return priceList.ToDomain(), nil
}

// GetSocialNetworks получение соц сетей
func (r *Repository) GetSocialNetworks(ctx context.Context, freelancerID int64) (social_network.Networks, error) {
	sql := `
		select s.id, s.name
        from social_networks s
        inner join freelancer_social_networks_m2m fs ON s.id = fs.social_network_id
        where fs.freelancer_id = $1
        order by s.id
	`

	var socialNetworks socialDao.Networks
	err := pgxscan.Select(ctx, r.db, &socialNetworks, sql, freelancerID)
	if err != nil {
		return nil, err
	}

	return socialNetworks.ToDomain(), nil
}

// GetAdditionalCities получение информации о городах фрилансера
func (r *Repository) GetAdditionalCities(ctx context.Context, freelancerID int64) (map[int64]*city.City, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных городов фрилансера")
	sql := `
		select c.id, c.name
        from freelancers_city c
        inner join freelancer_city_m2m fc ON c.id = fc.city_id
        where fc.freelancer_id = $1
        order by c.name
	`

	var cities []*cityDAO.CityDAO
	if err := pgxscan.Select(ctx, r.db, &cities, sql, freelancerID); err != nil {
		return nil, err
	}

	result := make(map[int64]*city.City, len(cities))
	for _, c := range cities {
		result[c.ID] = c.ToDomain()
	}

	return result, nil
}

// GetAdditionalSpecialities получение информации о специальностях фрилансера
func (r *Repository) GetAdditionalSpecialities(ctx context.Context, freelancerID int64) (map[int64]*speciality.Speciality, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных специальностей фрилансера")
	sql := `
		select s.id, s.name
		from freelancers_speciality s
		inner join freelancer_speciality_m2m fs ON s.id = fs.speciality_id
		where fs.freelancer_id = $1
        order by s.name
	`

	var specialities []*specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, freelancerID); err != nil {
		return nil, err
	}

	result := make(map[int64]*speciality.Speciality, len(specialities))
	for _, s := range specialities {
		result[s.ID] = s.ToDomain()
	}

	return result, nil
}
