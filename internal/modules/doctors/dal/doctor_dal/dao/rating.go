package dao

import "database/sql"

type MBCAggDAO struct {
	DoctorID int64 `db:"doctor_id"`
	MBCCoins int64 `db:"mbc_coins"`
}

type RatingDoctorDAO struct {
	ID           int64          `db:"id"`
	Slug         string         `db:"slug"`
	Name         string         `db:"name"`
	S3Image      sql.NullString `db:"s3_image"`
	CityID       sql.NullInt64  `db:"city_id"`
	SpecialityID sql.NullInt64  `db:"speciallity_id"`
}
