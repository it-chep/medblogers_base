package get_all_cities

import "context"

// Action список городов
type Action struct{}

// New .
func New() *Action {
	return &Action{}
}

// Do выполнение
func (a Action) Do(ctx context.Context) error {
	return nil
}
