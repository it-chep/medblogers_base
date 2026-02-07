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

type GetcourseUserDAO struct {
	ID        int64          `db:"id"`
	SbID      sql.NullInt64  `db:"sb_id"`
	GkID      sql.NullInt64  `db:"gk_id"`
	Name      sql.NullString `db:"name"`
	EndDate   sql.NullTime   `db:"end_date"`
	DaysCount sql.NullInt64  `db:"days_count"`
}

type GetcourseUsers []GetcourseUserDAO

func (u GetcourseUsers) GetSbIDs() []int64 {
	res := make([]int64, 0, len(u))
	for _, user := range u {
		if user.SbID.Valid {
			res = append(res, user.SbID.Int64)
		}
	}
	return res
}
