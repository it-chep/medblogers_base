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
	MMLink     sql.NullString `db:"m_link"`
	CreatedBy  time.Time      `db:"created_by"`
	IsActive   sql.NullBool   `json:"is_active"`
}
