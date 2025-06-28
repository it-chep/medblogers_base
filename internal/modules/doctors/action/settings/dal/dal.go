package dal

import (
	"context"
)

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) GetDoctorsCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) as doctors_count
		from docstar_site_doctor d
		where d.is_active = true
	`

	var count int64
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &count, sql); err != nil {
		return 0, err
	}

	return count, nil
}
