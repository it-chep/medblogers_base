package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/doctors/dal/vip_card_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/vip_card"
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

// GetVipCardInfo получение данных вип карточки доктора
func (r *Repository) GetVipCardInfo(ctx context.Context, doctorID int64) (*vip_card.VipCard, error) {
	sql := `
		select can_barter, can_buy_advertising, can_sell_advertising, short_message, advertising_price_from, blog_info
		from vip_card
		where doctor_id = $1 and is_active is true
	`

	var cardDao dao.VipCardDao
	err := pgxscan.Get(ctx, r.db, &cardDao, sql, doctorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return cardDao.ToDomain(), nil
}
