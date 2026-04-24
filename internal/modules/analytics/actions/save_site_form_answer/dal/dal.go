package dal

import (
	"context"
	"medblogers_base/internal/modules/analytics/actions/save_site_form_answer/dto"
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

func (r *Repository) CreateSiteFormAnswer(ctx context.Context, req dto.CreateSiteFormAnswerRequest) error {
	sql := `
		insert into site_forms_answers (
			form_name,
			answer,
			cookie_id,
			source,
			tg
		)
		values ($1, $2::jsonb, $3, $4, $5)
	`

	args := []interface{}{
		req.FormName,
		req.Answer,
		req.CookieID,
		req.Source,
		req.TG,
	}
	_, err := r.db.Exec(ctx, sql, args...)

	return err
}
