package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetExpiredVipDoctorIDs(ctx context.Context) ([]int64, error) {
	sql := `
		select distinct v.doctor_id
		from vip_card v
			join docstar_site_doctor d on d.id = v.doctor_id
		where v.end_date <= now() 
		  and d.is_vip is true
	`

	var doctorIDs []int64
	err := pgxscan.Select(ctx, r.db, &doctorIDs, sql)
	if err != nil {
		return nil, err
	}

	return doctorIDs, nil
}
