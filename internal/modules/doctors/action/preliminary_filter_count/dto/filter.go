package dto

type Filter struct {
	// Фильтры относящиеся к подписчикам
	MaxSubscribers int64
	MinSubscribers int64
	SocialMedia    []string

	// Дефолтные фильтры
	Page         int64
	Cities       []int64
	Specialities []int64
}
