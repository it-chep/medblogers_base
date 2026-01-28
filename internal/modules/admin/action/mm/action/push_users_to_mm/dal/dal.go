package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/action/mm/action/push_users_to_mm/dto"
	"medblogers_base/internal/pkg/postgres"
)

const (
	StateActive   = 1
	StateInactive = 2
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

func (r *Repository) GetNearestMM(ctx context.Context) (dto.MM, error) {
	sql := `
		select * 
		from mm 
		where is_active is true 
		  and state = $1 
		  and mm_datetime between now() - interval '5 minutes' and now()
  	`

	var mm dto.MM
	err := pgxscan.Get(ctx, r.db, &mm, sql, StateActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.MM{}, nil
		}
		return dto.MM{}, err
	}

	return mm, nil
}

// SetMMPassed ставит, что ММ прошло
func (r *Repository) SetMMPassed(ctx context.Context, mmID int64) error {
	sql := `
		update mm set is_active = false, state = $1 where id = $2
	`

	args := []interface{}{
		StateInactive,
		mmID,
	}
	_, err := r.db.Exec(ctx, sql, args...)
	return err
}

// GetUserToNotificate получаем пользователей для рассылки у кого больше 6 месяцев клуба в общем
func (r *Repository) GetUserToNotificate(ctx context.Context) (dto.GetcourseUsers, error) {
	sql := `select * from getcourse_users where days_count >= 180 and sb_id is not null and end_date >= now()`

	var users dto.GetcourseUsers
	err := pgxscan.Select(ctx, r.db, &users, sql)
	return users, err
}

// CreateNewsletter создание рассылки
func (r *Repository) CreateNewsletter(ctx context.Context, newsletterUUID uuid.UUID, plannedUsers []int64) error {
	sql := `insert into newsletter (newsletter_uuid, planned_sb_ids, event_type) values ($1, $2, 'mm_notification')`

	_, err := r.db.Exec(ctx, sql, newsletterUUID, plannedUsers)
	return err
}

// SendNewsletter сохранение факта отправки рассылки
func (r *Repository) SendNewsletter(ctx context.Context, newsletterUUID uuid.UUID, sbID int64) error {
	sql := `insert into sent_newsletter (newsletter_uuid, sb_id) values ($1, $2)`

	_, err := r.db.Exec(ctx, sql, newsletterUUID, sbID)
	return err
}
