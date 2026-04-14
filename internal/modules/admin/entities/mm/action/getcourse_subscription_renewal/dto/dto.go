package dto

import (
	"database/sql"
	"time"
)

type Request struct {
	GkID      int64
	DaysCount int64
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
