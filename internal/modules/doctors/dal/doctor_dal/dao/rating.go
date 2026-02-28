package dao

import "database/sql"

type RatingItemDAO struct {
	Slug           string         `db:"slug"`
	Name           string         `db:"name"`
	S3Image        sql.NullString `db:"s3_image"`
	CityID         sql.NullInt64  `db:"city_id"`
	CityName       sql.NullString `db:"city_name"`
	SpecialityID   sql.NullInt64  `db:"speciality_id"`
	SpecialityName sql.NullString `db:"speciality_name"`
	MBCCoins       int64          `db:"mbc_coins"`
}
