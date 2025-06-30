package speciality_dal

import (
	"context"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/speciality"

	"github.com/georgysavva/scany/pgxscan"
)

// Repository специальности
type Repository struct {
}

// NewRepository создает новый репозиторий по работе со специальностями
func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) GetSpecialities(ctx context.Context) ([]*speciality.Speciality, error) {
	sql := `
		select s.id                      as speciality_id,
			   s.name                    as speciality_name,
			   count(distinct doctor_id) as doctors_count
		from docstar_site_speciallity s
				 left join (select dc.speciallity_id, dc.doctor_id
							from docstar_site_doctor_additional_specialties dc
									 join docstar_site_doctor d on dc.doctor_id = d.id
							where d.is_active = true) as combined on s.id = combined.speciallity_id
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
