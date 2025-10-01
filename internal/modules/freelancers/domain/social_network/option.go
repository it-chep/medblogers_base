package social_network

// Option .
type Option func(s *SocialNetwork)

// WithID .
func WithID(id int64) Option {
	return func(s *SocialNetwork) {
		s.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(s *SocialNetwork) {
		s.name = name
	}
}

// WithSlug .
func WithSlug(slug string) Option {
	return func(s *SocialNetwork) {
		s.slug = slug
	}
}

// WithFreelancersCount .
func WithFreelancersCount(freelancersCount int64) Option {
	return func(s *SocialNetwork) {
		s.freelancersCount = freelancersCount
	}
}
