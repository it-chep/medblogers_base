package speciality

// Option .
type Option func(s *Speciality)

// WithID .
func WithID(id SpecialityID) Option {
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

// WithDoctorsCount .
func WithDoctorsCount(doctorsCount int64) Option {
	return func(s *Speciality) {
		s.doctorsCount = doctorsCount
	}
}

// WithPrimarySpecialityID .
func WithPrimarySpecialityID(primarySpecialityID *int64) Option {
	return func(s *Speciality) {
		s.primarySpecialityID = primarySpecialityID
	}
}

// WithIsOnlyAdditional .
func WithIsOnlyAdditional(isOnlyAdditional *bool) Option {
	return func(s *Speciality) {
		s.isOnlyAdditional = isOnlyAdditional
	}
}
