package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/doctor"
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

func (r *Repository) GetRecommendations(ctx context.Context, freelancerID int64) ([]int64, error) {
	sql := `
		select doctor_id from freelancer_recommendation where freelancer_id = $1
	`
	// todo cортировочку
	var recommendations []int64
	err := pgxscan.Select(ctx, r.db, &recommendations, sql, freelancerID)
	if err != nil {
		return nil, err
	}
	return recommendations, nil
}

func (r *Repository) GetRecommendationInfoByIDs(ctx context.Context, doctorIDs []int64) ([]*doctor.Doctor, error) {
	sql := `
		select id, name from docstar_site_doctor where id = any($1)
	`

	var recommendations dao.Recommendations
	err := pgxscan.Select(ctx, r.db, &recommendations, sql, doctorIDs)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return recommendations.ToDomain(), nil
}
