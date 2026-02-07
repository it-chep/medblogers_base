package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
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

// GetAdditionalSpecialities получение доп специальностей доктора
func (r *Repository) GetAdditionalSpecialities(ctx context.Context, doctorID int64) ([]*speciality.Speciality, error) {
	sql := `
		select s.id as "id", s.name as "name"
		from docstar_site_doctor_additional_specialties ads
			join docstar_site_speciallity s on ads.speciallity_id = s.id
			join docstar_site_doctor d on ads.doctor_id = d.id
		where ads.doctor_id = $1 and ads.doctor_id != d.id
	`

	var specialities dao.SpecialitiesDAO
	err := pgxscan.Select(ctx, r.db, &specialities, sql, doctorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return specialities.ToDomain(), err
}
