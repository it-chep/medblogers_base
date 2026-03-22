package dto

type Banner struct {
	ID              int64
	Name            string
	IsActive        bool
	OrderingNumber  int64
	DesktopImage    string
	DesktopFileType string
	MobileImage     string
	MobileFileType  string
	BannerLink      string
}

type UpdateRequest struct {
	Name           string
	OrderingNumber int64
	BannerLink     string
}
