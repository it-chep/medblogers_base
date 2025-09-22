package speciality_dal

import (
	"context"
	"medblogers_base/internal/modules/freelancers/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
)

// Repository специальности
type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе со специальностями
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetSpecialitiesWithFreelancersCount получение списка специальностей с количеством фрилансеров
func (r Repository) GetSpecialitiesWithFreelancersCount(ctx context.Context) ([]*speciality.Speciality, error) {
	logger.Message(ctx, "[DAL] Запрос специальностей")

	sql := `
	with active_freelancers_in_specialities as (select fs.speciality_id,
												   fs.freelancer_id
											from freelancer_speciality_m2m fs
													 join freelancer f on fs.freelancer_id = f.id
											where f.is_active = true)
	select s.id                         as id,
		   s.name                       as name,
		   count(distinct af.freelancer_id) as freelancers_count
	from freelancers_speciality s
			 left join active_freelancers_in_specialities af on s.id = af.speciality_id
	group by s.id, s.name
	having count(distinct af.freelancer_id) > 0
	order by s.name
	`

	var specialitiesDAO []dao.SpecialityDAOWithFreelancersCount
	if err := pgxscan.Select(ctx, r.db, &specialitiesDAO, sql); err != nil {
		return nil, err
	}

	specialities := make([]*speciality.Speciality, 0, len(specialitiesDAO))
	for _, specDao := range specialitiesDAO {
		specialities = append(specialities, specDao.ToDomain())
	}

	return specialities, nil
}

// GetAllSpecialities все специальности
func (r Repository) GetAllSpecialities(ctx context.Context) ([]*speciality.Speciality, error) {
	sql := `
		select s.id                      as id,
			   s.name                    as name
		from freelancers_speciality s
		group by s.id, s.name
		order by s.name
	`

	var specialitiesDAO []dao.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specialitiesDAO, sql); err != nil {
		return nil, err
	}

	specialities := make([]*speciality.Speciality, 0, len(specialitiesDAO))
	for _, specDao := range specialitiesDAO {
		specialities = append(specialities, specDao.ToDomain())
	}

	return specialities, nil
}
