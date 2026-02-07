package dto

import (
	"database/sql"
	"time"
)

// CreateOrderRequest данне от ГК
type CreateOrderRequest struct {
	OrderID  int64
	Position string
	GkID     int64
	UserName string
}

// GetcourseOrder представление для базки
type GetcourseOrder struct {
	Name      string
	DaysCount int64
	GkID      int64
	OrderID   int64
}

type CreateUserRequest struct {
	GkID      int64
	Name      string
	EndDate   time.Time
	DaysCount int64
}

type GetcourseUserDAO struct {
	ID        int64          `db:"id"`
	SbID      sql.NullInt64  `db:"sb_id"`
	GkID      sql.NullInt64  `db:"gk_id"`
	Name      sql.NullString `db:"name"`
	EndDate   sql.NullTime   `db:"end_date"`
	DaysCount sql.NullInt64  `db:"days_count"`
}
