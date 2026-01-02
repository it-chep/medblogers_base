package dal

import (
	"context"
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

func (r *Repository) DeleteFreelancer(ctx context.Context, id int64) error {
	sql := `
		delete from freelancer where id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}

func (r *Repository) DeleteFreelancerRecommendation(ctx context.Context, id int64) error {
	sql := `
		delete from freelancer_recommendation where freelancer_id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}

func (r *Repository) DeleteFreelancerNetworks(ctx context.Context, id int64) error {
	sql := `
		delete from freelancer_social_networks_m2m where freelancer_id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}

func (r *Repository) DeleteFreelancerCities(ctx context.Context, id int64) error {
	sql := `
		delete from freelancer_city_m2m where freelancer_id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}

func (r *Repository) DeleteFreelancerPriceList(ctx context.Context, id int64) error {
	sql := `
		delete from freelancers_price_list where freelancer_id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}

func (r *Repository) DeleteFreelancerSpecialities(ctx context.Context, id int64) error {
	sql := `
		delete from freelancer_speciality_m2m where freelancer_id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}
