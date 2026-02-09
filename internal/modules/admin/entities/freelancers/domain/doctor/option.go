package doctor

import (
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

// Option .
type Option func(c *Doctor)

// WithID .
func WithID(id int64) Option {
	return func(s *Doctor) {
		s.medblogersID = MedblogersID(id)
	}
}

// WithSlug .
func WithSlug(slug string) Option {
	return func(s *Doctor) {
		s.slug = slug
	}
}

// WithName .
func WithName(name string) Option {
	return func(s *Doctor) {
		s.name = name
	}
}

// WithCityName .
func WithCityName(name string) Option {
	return func(s *Doctor) {
		s.cityName = name
	}
}

// WithSpecialityName .
func WithSpecialityName(name string) Option {
	return func(s *Doctor) {
		s.specialityName = name
	}
}

// WithS3Image .
func WithS3Image(s3key S3Key) Option {
	return func(s *Doctor) {
		s.s3Image = s3key
	}
}

// WithMainCityID .
func WithMainCityID(mainCityID int64) Option {
	return func(s *Doctor) {
		s.cityID = city.CityID(mainCityID)
	}
}

// WithMainSpecialityID .
func WithMainSpecialityID(mainSpecialityID int64) Option {
	return func(s *Doctor) {
		s.specialityID = speciality.SpecialityID(mainSpecialityID)
	}
}

// WithIsActive .
func WithIsActive(isActive bool) Option {
	return func(s *Doctor) {
		s.isActive = isActive
	}
}
