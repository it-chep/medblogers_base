package dal

import (
	"context"
	"fmt"
	cityDAO "medblogers_base/internal/modules/freelancers/dal/city_dal/dao"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_dal/dao"
	specialityDAO "medblogers_base/internal/modules/freelancers/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с фрилансерами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetFreelancerInfo получает информацию о докторе
func (r Repository) GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error) {
	sql := `
		select 
			id, name, s3_image, is_worked_with_doctors, start_working_date
		from freelancer
		where slug = $1
	`

	var freelancerDAO dao.FreelancerSeoInfo
	err := pgxscan.Get(ctx, r.db, &freelancerDAO, sql, slug)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, fmt.Errorf("freelancer with slug %s not found", slug)
	case err != nil:
		return nil, fmt.Errorf("failed to get freelancer: %w", err)
	}

	return freelancerDAO.ToDomain(), nil
}

// GetFreelancerAdditionalCities получение информации о городах фрилансера
func (r Repository) GetFreelancerAdditionalCities(ctx context.Context, freelancerID int64) (map[int64]*city.City, error) {
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

// GetFreelancerAdditionalSpecialities получение информации о специальностях фрилансера
func (r Repository) GetFreelancerAdditionalSpecialities(ctx context.Context, freelancerID int64) (map[int64]*speciality.Speciality, error) {
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
