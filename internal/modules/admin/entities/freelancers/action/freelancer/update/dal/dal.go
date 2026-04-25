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
		tg_username = $3,
		portfolio_link = $4,
		speciality_id = $5,
		city_id = $6,
		price_category = $7,
		start_working_date = $8,
		agency_representative = $9,
		cooperation_type_id = $10,
		has_med_education = $11
   	where id = $1
	`

	args := []interface{}{
		freelancerID,             // $1
		req.Name,                 // $2
		req.TgURL,                // $3
		req.PortfolioLink,        // $4
		req.MainSpecialityID,     // $5
		req.MainCityID,           // $6
		req.PriceCategory,        // $7
		req.DateStarted,          // $8
		req.AgencyRepresentative, // $9
		req.CooperationTypeID,    // $10
		req.HasMedEducation,      // $11
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
