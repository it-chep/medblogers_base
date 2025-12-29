package blacklist_check

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/blacklist_check/service/checker"
	"medblogers_base/internal/modules/doctors/client"
)

// Action ...
type Action struct {
	checker *checker.Service
}

// New ...
func New(clients *client.Aggregator) *Action {
	return &Action{
		checker: checker.NewService(clients.Subscribers),
	}
}

// Do ...
func (a *Action) Do(ctx context.Context, telegram string) (bool, error) {
	return a.checker.Check(ctx, telegram)
}
