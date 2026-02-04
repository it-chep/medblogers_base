package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_price_list/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/pkg/postgres"
	"strconv"
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

func (r *Repository) GetPriceList(ctx context.Context, freelancerID int64) ([]dto.PriceList, error) {
	sql := `
		select id, name, price from freelancers_price_list where freelancer_id = $1
	`

	var priceList []dao.PriceListDao
	err := pgxscan.Select(ctx, r.db, &priceList, sql, freelancerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return lo.Map(priceList, func(item dao.PriceListDao, _ int) dto.PriceList {
		return dto.PriceList{
			ID:     item.ID,
			Name:   item.Name,
			Amount: strconv.FormatInt(item.Price, 10),
		}
	}), nil
}
