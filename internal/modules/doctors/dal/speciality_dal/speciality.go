package speciality_dal

import (
	"context"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/speciality"

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

// GetSpecialitiesWithDoctorsCount получение списка специальностей с количеством докторов
func (r Repository) GetSpecialitiesWithDoctorsCount(ctx context.Context) ([]*speciality.Speciality, error) {
	logger.Message(ctx, "[DAL] Запрос специальностей")

	sql := `
	with active_doctors_in_specialities as (select ds.speciallity_id,
												   ds.doctor_id
											from docstar_site_doctor_additional_specialties ds
													 join docstar_site_doctor d on ds.doctor_id = d.id
											where d.is_active = true)
	select s.id                         as id,
		   s.name                       as name,
		   count(distinct ad.doctor_id) as doctors_count
	from docstar_site_speciallity s
			 left join active_doctors_in_specialities ad on s.id = ad.speciallity_id
	group by s.id, s.name
	having count(distinct ad.doctor_id) > 0
	order by s.name
	`

	var specialitiesDAO []specialityDAO.SpecialityDAOWithDoctorsCount
	if err := pgxscan.Select(ctx, r.db, &specialitiesDAO, sql); err != nil {
		return nil, err
	}

	specialities := make([]*speciality.Speciality, 0, len(specialitiesDAO))
	for _, dao := range specialitiesDAO {
		specialities = append(specialities, dao.ToDomain())
	}

	return specialities, nil
}

// GetAllSpecialities все специальности
func (r Repository) GetAllSpecialities(ctx context.Context) ([]*speciality.Speciality, error) {
	sql := `
		select s.id                      as id,
			   s.name                    as name
		from docstar_site_speciallity s
		group by s.id, s.name
		order by s.name
	`

	var specialitiesDAO []specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specialitiesDAO, sql); err != nil {
		return nil, err
	}

	specialities := make([]*speciality.Speciality, 0, len(specialitiesDAO))
	for _, dao := range specialitiesDAO {
		specialities = append(specialities, dao.ToDomain())
	}

	return specialities, nil
}
