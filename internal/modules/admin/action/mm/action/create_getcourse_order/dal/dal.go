package dal

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/action/mm/action/create_getcourse_order/dto"
	"medblogers_base/internal/pkg/postgres"
	"time"

	"github.com/georgysavva/scany/pgxscan"
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

// CreateGetcourseOrder создает заказ от геткурса
func (r *Repository) CreateGetcourseOrder(ctx context.Context, req dto.GetcourseOrder) error {
	sql := `insert into getcourse_orders (order_id, gk_id, days_count, name) values ($1, $2, $3, $4)`

	args := []interface{}{
		req.OrderID,
		req.GkID,
		req.DaysCount,
		req.Name,
	}
	_, err := r.db.Exec(ctx, sql, args...)
	return err
}

// GetUserByGKID получение пользака
func (r *Repository) GetUserByGKID(ctx context.Context, gkID int64) (dto.GetcourseUserDAO, error) {
	sql := `select * from getcourse_users where gk_id = $1`

	var user dto.GetcourseUserDAO
	err := pgxscan.Get(ctx, r.db, &user, sql, gkID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.GetcourseUserDAO{}, nil
		}
		return dto.GetcourseUserDAO{}, err
	}
	return user, err
}

// CreateGetcourseUser создание пользователя геткурс
func (r *Repository) CreateGetcourseUser(ctx context.Context, req dto.CreateUserRequest) error {
	sql := `insert into getcourse_users (gk_id, name, end_date, days_count) values ($1, $2, $3, $4)`

	args := []interface{}{
		req.GkID,
		req.Name,
		req.EndDate,
		req.DaysCount,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	return err
}

// UpdateUserSubscription обновление подписки
func (r *Repository) UpdateUserSubscription(ctx context.Context, gkID, daysCount int64, endDate time.Time) error {
	sql := `update getcourse_users set days_count = $2, end_date = $3 where gk_id = $1`

	_, err := r.db.Exec(ctx, sql, gkID, daysCount, endDate)
	return err
}
