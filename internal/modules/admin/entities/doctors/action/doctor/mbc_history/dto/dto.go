package dto

import "time"

type MBCHistoryItem struct {
	MBCCount   int64
	OccurredAt time.Time
}
