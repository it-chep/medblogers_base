package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/doctor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
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

func (r *Repository) GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error) {
	sql := `
		select
		    id,
		    email, 
		    slug, 
		    name, 
		    is_active, 
		    tg_username, 
		    portfolio_link, 
		    speciality_id, 
		    city_id, 
		    price_category, 
		    s3_image, 
		    start_working_date, 
		    cooperation_type_id, 
		    agency_representative
		from freelancer where id = $1
	`

	var freelancerDAO dao.FullFreelancerDAO
	err := pgxscan.Get(ctx, r.db, &freelancerDAO, sql, freelancerID)

	return freelancerDAO.ToDomain(), err
}

func (r *Repository) GetFreelancers(ctx context.Context) ([]*freelancer.Freelancer, error) {
	sql := `
		select
		    id, name, is_active 
		from freelancer 
		order by id desc
	`

	var miniatures dao.MiniatureList
	err := pgxscan.Select(ctx, r.db, &miniatures, sql)
	if err != nil {
		return nil, err
	}

	return miniatures.ToDomain(), nil
}

func (r *Repository) GetDoctorToRecommendation(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	sql := `
		select id, name, slug, is_active 
		from docstar_site_doctor 
		where id = $1 and is_active is true
	`

	var recoDoctor dao.RecommendationDoctorDAO
	err := pgxscan.Get(ctx, r.db, &recoDoctor, sql, doctorID)
	if err != nil {
		return nil, err
	}

	return recoDoctor.ToDomain(), nil
}
