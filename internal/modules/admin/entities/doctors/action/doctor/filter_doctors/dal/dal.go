package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
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

// FilterDoctors фильтрация докторов
func (r *Repository) FilterDoctors(ctx context.Context, specialitiesIDs []int64) ([]*doctor.Doctor, error) {
	sql := `
		select distinct d.id,
			d.name,
			d.s3_image,
			d.cooperation_type,
			d.is_active,
			d.date_created
		from docstar_site_doctor d
			join docstar_site_doctor_additional_specialties das on d.id = das.doctor_id
		where das.speciallity_id = any($1)
		order by id desc 
	`

	var doctorsDAO dao.DoctorMiniatureList
	err := pgxscan.Select(ctx, r.db, &doctorsDAO, sql, specialitiesIDs)
	if err != nil {
		return nil, err
	}

	return doctorsDAO.ToDomain(), nil
}
