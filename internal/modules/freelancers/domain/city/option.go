package city

// Option .
type Option func(c *City)

// WithID .
func WithID(id int64) Option {
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

// WithFreelancersCount .
func WithFreelancersCount(freelancersCount int64) Option {
	return func(s *City) {
		s.freelancersCount = freelancersCount
	}
}
