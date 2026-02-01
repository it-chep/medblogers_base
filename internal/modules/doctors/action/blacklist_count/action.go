package blacklist_count

import (
	"context"
	"medblogers_base/internal/modules/doctors/client"
)

// Storage .
type Storage interface {
	GetBlackListCount(ctx context.Context) (int64, error)
}

// Action .
type Action struct {
	gateway Storage
}

// New .
func New(clients *client.Aggregator) *Action {
	return &Action{
		gateway: clients.Subscribers,
	}
}

// Do ..
func (a *Action) Do(ctx context.Context) (int64, error) {
	return a.gateway.GetBlackListCount(ctx)
}
