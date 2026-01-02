package freelancer

import "time"

// Option .
type Option func(c *Freelancer)

// WithID .
func WithID(id int64) Option {
	return func(s *Freelancer) {
		s.id = id
	}
}

// WithSlug .
func WithSlug(slug string) Option {
	return func(s *Freelancer) {
		s.slug = slug
	}
}

// WithName .
func WithName(name string) Option {
	return func(s *Freelancer) {
		s.name = name
	}
}

// WithCityName .
func WithCityName(name string) Option {
	return func(s *Freelancer) {
		s.cityName = name
	}
}

// WithSpecialityName .
func WithSpecialityName(name string) Option {
	return func(s *Freelancer) {
		s.specialityName = name
	}
}

// WithS3Image .
func WithS3Image(s3key string) Option {
	return func(s *Freelancer) {
		s.s3Image = s3key
	}
}

// WithIsAgencyRepresentative .
func WithIsAgencyRepresentative(agencyRepresentative bool) Option {
	return func(s *Freelancer) {
		s.agencyRepresentative = agencyRepresentative
	}
}

// WithStartWorkingTime .
func WithStartWorkingTime(startWorkingTime time.Time) Option {
	return func(s *Freelancer) {
		s.startWorking = startWorkingTime
	}
}

// WithTgURL .
func WithTgURL(tgURL string) Option {
	return func(s *Freelancer) {
		s.tgURL = tgURL
	}
}

// WithPortfolioLink .
func WithPortfolioLink(portfolioLink string) Option {
	return func(s *Freelancer) {
		s.portfolioLink = portfolioLink
	}
}

// WithMainCityID .
func WithMainCityID(mainCityID int64) Option {
	return func(s *Freelancer) {
		s.cityID = mainCityID
	}
}

// WithMainSpecialityID .
func WithMainSpecialityID(mainSpecialityID int64) Option {
	return func(s *Freelancer) {
		s.specialityID = mainSpecialityID
	}
}

// WithPriceCategory .
func WithPriceCategory(priceCategory int64) Option {
	return func(s *Freelancer) {
		s.priceCategory = priceCategory
	}
}

func WithIsActive(isActive bool) Option {
	return func(s *Freelancer) {
		s.isActive = isActive
	}
}

func WithCooperationTypeID(cooperationTypeId int64) Option {
	return func(s *Freelancer) {
		s.cooperationType = cooperationTypeId
	}
}
