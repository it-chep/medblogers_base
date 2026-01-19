package manual_notification_mm

import "context"

type Action struct {
}

func New() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context) error {
	return nil
}
