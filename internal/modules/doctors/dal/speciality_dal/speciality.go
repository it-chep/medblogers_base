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
		select s.id as id,
			   s.name as name,
			   COUNT(distinct d.id) as doctors_count
		from docstar_site_speciallity s
		left join (
			-- Основные города докторов
			select d.id, d.speciallity_id 
			from docstar_site_doctor d
			where is_active = true and speciallity_id is not null
			
			union 
			
			-- Дополнительные города
			select dc.speciallity_id, dc.doctor_id
							from docstar_site_doctor_additional_specialties dc
									 join docstar_site_doctor d on dc.doctor_id = d.id
							where d.is_active = true
		) as combined on s.id = combined.speciallity_id
		left join docstar_site_doctor d ON d.id = combined.id
		group by s.id, s.name
		having count(distinct d.id) > 0
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
