package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/promo_offers/action/filter_settings/dto"
	"medblogers_base/internal/modules/promo_offers/dal/dao"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) from promo_offer
	`

	var count int64
	if err := r.db.QueryRow(ctx, sql).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetCooperationTypeCounts(ctx context.Context) ([]dto.CountItem, error) {
	sql := `
		select cooperation_type_id as id, count(*) as offers_count
		from promo_offer
		where cooperation_type_id is not null
		group by cooperation_type_id
	`

	var rows []dao.FilterCountDAO
	err := pgxscan.Select(ctx, r.db, &rows, sql)
	if err != nil {
		return nil, err
	}

	result := make([]dto.CountItem, 0, len(rows))
	for _, row := range rows {
		result = append(result, dto.CountItem{
			ID:          row.ID.Int64,
			OffersCount: row.OffersCount,
		})
	}

	return result, nil
}
