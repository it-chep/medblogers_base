package doctor

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
		s.cityID = mainCityID
	}
}

// WithMainSpecialityID .
func WithMainSpecialityID(mainSpecialityID int64) Option {
	return func(s *Doctor) {
		s.specialityID = mainSpecialityID
	}
}
