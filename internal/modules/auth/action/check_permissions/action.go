package check_permissions

import "context"

type Action struct {
}

func New() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context) {}
