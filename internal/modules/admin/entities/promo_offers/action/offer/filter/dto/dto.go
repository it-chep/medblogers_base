package dto

type Request struct {
	BrandIDs []int64
	IsActive *bool
	Page     int64
	Limit    int64
}
