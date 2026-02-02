package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
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

// GetAdditionalSpecialities получение доп специальностей фрилансера
func (r *Repository) GetAdditionalSpecialities(ctx context.Context, freelancerID int64) ([]*speciality.Speciality, error) {
	sql := `
		select s.id as "id", s.name as "name" 
		from freelancer_speciality_m2m m2m 
		    join freelancers_speciality s on m2m.speciality_id = s.id
			join freelancer f on m2m.freelancer_id = f.id
		where m2m.freelancer_id = $1 and f.speciality_id != m2m.speciality_id
	`

	var specialities dao.SpecialitiesDAO
	err := pgxscan.Select(ctx, r.db, &specialities, sql, freelancerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return specialities.ToDomain(), nil
}
