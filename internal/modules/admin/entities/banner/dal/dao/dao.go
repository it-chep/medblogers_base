package dao

import (
	"database/sql"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/admin/entities/banner/dto"
)

type Banner struct {
	ID              int64          `db:"id"`
	Name            sql.NullString `db:"name"`
	IsActive        bool           `db:"is_active"`
	OrderingNumber  int64          `db:"ordering_number"`
	DesktopImage    sql.NullString `db:"desktop_image"`
	DesktopFileType sql.NullString `db:"desktop_file_type"`
	MobileImage     sql.NullString `db:"mobile_image"`
	MobileFileType  sql.NullString `db:"mobile_file_type"`
	BannerLink      sql.NullString `db:"banner_link"`
}

func (b Banner) ToDTO() dto.Banner {
	return dto.Banner{
		ID:              b.ID,
		Name:            b.Name.String,
		IsActive:        b.IsActive,
		OrderingNumber:  b.OrderingNumber,
		DesktopImage:    b.DesktopImage.String,
		DesktopFileType: b.DesktopFileType.String,
		MobileImage:     b.MobileImage.String,
		MobileFileType:  b.MobileFileType.String,
		BannerLink:      b.BannerLink.String,
	}
}

type Banners []Banner

func (b Banners) ToDTO() []dto.Banner {
	return lo.Map(b, func(item Banner, _ int) dto.Banner {
		return item.ToDTO()
	})
}
