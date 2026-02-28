package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/mbc_history/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetMBCHistory(ctx context.Context, doctorID int64) ([]dto.MBCHistoryItem, error) {
	sql := `
		select mbc_count, occurred_at
		from mbc_operation
		where doctor_id = $1 and occurred_at >= NOW() - interval '1 year'
		order by occurred_at desc 
	`

	var items []dao.MBCHistoryItemDAO
	err := pgxscan.Select(ctx, r.db, &items, sql, doctorID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.MBCHistoryItem, 0, len(items))
	for _, item := range items {
		result = append(result, dto.MBCHistoryItem{
			MBCCount:   item.MBCCount,
			OccurredAt: item.OccurredAt,
		})
	}

	return result, nil
}
