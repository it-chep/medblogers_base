package dto

import (
	"time"
)

type CreateMMRequest struct {
	MMDatetime time.Time
	Name       string
	MMLink     string
}
