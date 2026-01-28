package dto

import (
	"database/sql"
	"strconv"
	"strings"
)

type GetcourseUserDAO struct {
	ID        int64          `db:"id"`
	SbID      sql.NullInt64  `db:"sb_id"`
	GkID      sql.NullInt64  `db:"gk_id"`
	Name      sql.NullString `db:"name"`
	EndDate   sql.NullTime   `db:"end_date"`
	DaysCount sql.NullInt64  `db:"days_count"`
}

type GetcourseUsers []GetcourseUserDAO

func (u GetcourseUsers) StringGkIDs() string {
	if len(u) == 0 {
		return ""
	}

	var builder strings.Builder

	// Предварительно оцениваем размер
	// Предполагаем, что средний ID 8 символов + запятая
	builder.Grow(len(u) * 10)

	first := true
	for _, user := range u {
		if user.GkID.Valid {
			if !first {
				builder.WriteString(", ")
			}
			builder.WriteString(strconv.FormatInt(user.GkID.Int64, 10))
			first = false
		}
	}

	return builder.String()
}
