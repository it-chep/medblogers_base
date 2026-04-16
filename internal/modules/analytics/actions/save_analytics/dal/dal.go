package dal

import (
	"context"
	"medblogers_base/internal/modules/analytics/actions/save_analytics/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository .
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) IsCookieUserExists(ctx context.Context, cookieID uuid.UUID) (bool, error) {
	sql := `
		select exists(
			select 1
			from cookie_users
			where cookie_id = $1
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, sql, cookieID).Scan(&exists)
	return exists, err
}

func (r *Repository) IsAnalyticsExistsForLast7Days(ctx context.Context, req dto.CreateAnalyticsRequest) (bool, error) {
	sql := `
		select exists(
			select 1
			from utm_analytics
			where created_at >= now() - interval '7 days'
			  and utm_source = $1
			  and utm_medium = $2
			  and utm_campaign = $3
			  and utm_term = $4
			  and utm_content = $5
			  and domain_name = $6
			  and cookie_id = $7
			  and company is not distinct from $8
			  and event is not distinct from $9
		)
	`

	var exists bool
	err := r.db.QueryRow(
		ctx,
		sql,
		req.UtmSource,
		req.UtmMedium,
		req.UtmCampaign,
		req.UtmTerm,
		req.UtmContent,
		req.DomainName,
		req.CookieID,
		req.Company,
		req.Event,
	).Scan(&exists)

	return exists, err
}

func (r *Repository) CreateAnalytics(ctx context.Context, req dto.CreateAnalyticsRequest) error {
	sql := `
		insert into utm_analytics (
			uuid,
			domain_name,
			company,
			event,
			cookie_id,
			utm_source,
			utm_medium,
			utm_campaign,
			utm_term,
			utm_content
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	args := []interface{}{
		req.ID,
		req.DomainName,
		req.Company,
		req.Event,
		req.CookieID,
		req.UtmSource,
		req.UtmMedium,
		req.UtmCampaign,
		req.UtmTerm,
		req.UtmContent,
	}

	_, err := r.db.Exec(ctx, sql, args...)

	return err
}
