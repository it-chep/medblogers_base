package get_all_specialities

import (
	"context"
)

// Action список специальностей
type Action struct{}

// New .
func New() *Action {
	return &Action{}
}

// Do выполнение
func (a Action) Do(ctx context.Context) error {
	return nil
}
