package dao

import (
	"database/sql"
	"medblogers_base/internal/modules/freelancers/domain/doctor"
)

// DoctorMiniatureDAO .
type DoctorMiniatureDAO struct {
	ID           int64          `db:"id"`
	Name         string         `db:"name"`
	Slug         string         `db:"slug"`
	CityID       int64          `db:"city_id"`
	SpecialityID int64          `db:"speciallity_id"`
	S3Image      sql.NullString `db:"s3_image"`
}

type DoctorMiniatureDAOs []DoctorMiniatureDAO

func (ds DoctorMiniatureDAOs) ToDomain() []*doctor.Doctor {
	domains := make([]*doctor.Doctor, 0, len(ds))
	for _, doc := range ds {
		domains = append(domains, doc.ToDomain())
	}
	return domains
}

// ToDomain конвертирует DAO в доменный объект
func (d DoctorMiniatureDAO) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(d.ID),
		doctor.WithSlug(d.Slug),
		doctor.WithName(d.Name),
		doctor.WithMainSpecialityID(d.SpecialityID),
		doctor.WithMainCityID(d.CityID),
		doctor.WithS3Image(doctor.S3Key(d.S3Image.String)),
	)
}
