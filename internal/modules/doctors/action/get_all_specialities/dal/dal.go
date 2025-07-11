package dal

import (
	"context"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с городами
func NewRepository() *Repository {
	return &Repository{}
}

// GetAllSpecialities все специальности
func (r Repository) GetAllSpecialities(ctx context.Context) ([]*speciality.Speciality, error) {
	sql := `
		select s.id                      as speciality_id,
			   s.name                    as speciality_name
		from docstar_site_speciallity s
		group by s.id, s.name
		order by s.name
	`

	var specialitiesDAO []specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &specialitiesDAO, sql); err != nil {
		return nil, err
	}

	specialities := make([]*speciality.Speciality, 0, len(specialitiesDAO))
	for _, dao := range specialitiesDAO {
		specialities = append(specialities, dao.ToDomain())
	}

	return specialities, nil
}
