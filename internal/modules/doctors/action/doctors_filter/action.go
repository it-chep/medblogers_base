package doctors_filter

import "context"

// Action фильтрация докторов
type Action struct{}

// New .
func New() *Action {
	return &Action{}
}

func (a Action) Do(ctx context.Context) {
	//	Если фильтр не подписчиков. Берем докторов, обогажаем миниатюрами
	//	Если фильтр подписчиков. Фильтруем в подписчиках, по ID получаем докторов и добавляем другие фильтры
}
