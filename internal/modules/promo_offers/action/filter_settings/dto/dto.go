package dto

type CountItem struct {
	ID          int64
	Name        string
	OffersCount int64
}

type Response struct {
	All              int64
	CooperationTypes []CountItem
}
