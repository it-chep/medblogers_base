package dto

import "database/sql"

type Banner struct {
	ID              int64          `db:"id"`
	Name            sql.NullString `db:"name"`
	OrderingNumber  int64          `db:"ordering_number"`
	DesktopImage    sql.NullString `db:"desktop_image"`
	DesktopFileType sql.NullString `db:"desktop_file_type"`
	MobileImage     sql.NullString `db:"mobile_image"`
	MobileFileType  sql.NullString `db:"mobile_file_type"`
	BannerLink      sql.NullString `db:"banner_link"`
}
