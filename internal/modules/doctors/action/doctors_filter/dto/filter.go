package dto

type Sort int8

const (
	SubscribersDESC = iota
	SubscribersASC
	//NameASC
	//NameDESC
)

func StringToSort(s string) Sort {
	switch s {
	case "subscribers_desc":
		return SubscribersDESC
	case "subscribers_asc":
		return SubscribersASC
	//case "name_desc":
	//	return NameDESC
	//case "name_asc":
	//	return NameASC
	default:
		return SubscribersDESC
	}
}

func (s Sort) String() string {
	switch s {
	case SubscribersDESC:
		return "desc"
	case SubscribersASC:
		return "asc"
	//case NameDESC:
	//	return "desc"
	//case NameASC:
	//	return "asc"
	default:
		return "desc"
	}
}

type Filter struct {
	// Фильтры относящиеся к подписчикам
	MaxSubscribers int64
	MinSubscribers int64
	SocialMedia    []string

	// Дефолтные фильтры
	Page         int64
	Cities       []int64
	Specialities []int64

	// Сортировка
	Sort Sort
}
