package dao

import "time"

type MBCHistoryItemDAO struct {
	MBCCount   int64     `db:"mbc_count"`
	OccurredAt time.Time `db:"occurred_at"`
}
