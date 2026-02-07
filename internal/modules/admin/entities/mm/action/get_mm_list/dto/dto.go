package dto

import (
	"database/sql"
	"time"
)

type MM struct {
	ID         int64          `db:"id"`
	MMDatetime sql.NullTime   `db:"mm_datetime"`
	Name       sql.NullString `db:"name"`
	State      sql.NullString `db:"state"`
	MMLink     sql.NullString `db:"mm_link"`
	CreatedAt  time.Time      `db:"created_at"`
	IsActive   sql.NullBool   `json:"is_active"`
}
