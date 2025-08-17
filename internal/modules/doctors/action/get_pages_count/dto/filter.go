package dto

type Filter struct {
	// Фильтры относящиеся к подписчикам
	MaxSubscribers int64
	MinSubscribers int64
	SocialMedia    []string

	// Дефолтные фильтры
	Cities       []int64
	Specialities []int64
}
