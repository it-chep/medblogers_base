package speciality

// Option .
type Option func(s *Speciality)

// WithID .
func WithID(id int64) Option {
	return func(s *Speciality) {
		s.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(s *Speciality) {
		s.name = name
	}
}

// WithFreelancersCount .
func WithFreelancersCount(freelancersCount int64) Option {
	return func(s *Speciality) {
		s.freelancersCount = freelancersCount
	}
}
