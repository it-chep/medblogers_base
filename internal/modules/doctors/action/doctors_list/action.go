package doctors_list

import "context"

// Action список докторов
type Action struct{}

// New .
func New() *Action {
	return &Action{}
}

func (a Action) Do(ctx context.Context) {
	// Получаем докторов по лимиту
	// обогащаем подписчиков в миниатюры
}
