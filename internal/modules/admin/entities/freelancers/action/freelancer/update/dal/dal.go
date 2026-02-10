package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/dto"
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

// UpdateFreelancer обновление фрилансера
func (r *Repository) UpdateFreelancer(ctx context.Context, freelancerID int64, req dto.UpdateRequest) error {
	sql := `
	update freelancer set 
		name = $2,
		slug = $3,
		tg_username = $4,
		portfolio_link = $5,
		speciality_id = $6,
		city_id = $7,
		price_category = $8,
		start_working_date = $9,
		agency_representative = $10,
		cooperation_type_id = $11
   	where id = $1
	`

	args := []interface{}{
		freelancerID,
		req.Name,
		req.Slug,
		req.TgURL,
		req.PortfolioLink,
		req.MainSpecialityID,
		req.MainCityID,
		req.PriceCategory,
		req.DateStarted,
		req.AgencyRepresentative,
		req.CooperationTypeID,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
