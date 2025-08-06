package city

// Option .
type Option func(c *City)

// WithID .
func WithID(id CityID) Option {
	return func(s *City) {
		s.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(s *City) {
		s.name = name
	}
}

// WithDoctorsCount .
func WithDoctorsCount(doctorsCount int64) Option {
	return func(s *City) {
		s.doctorsCount = doctorsCount
	}
}
