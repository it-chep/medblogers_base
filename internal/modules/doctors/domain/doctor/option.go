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

// WithName .
func WithName(name string) Option {
	return func(s *Doctor) {
		s.name = name
	}
}

// WithS3Image .
func WithS3Image(s3key string) Option {
	return func(s *Doctor) {
		s.s3Image = s3key
	}
}

// WithTgURL .
func WithTgURL(tgURL string) Option {
	return func(s *Doctor) {
		s.tgURL = tgURL
	}
}

// WithTgChannelURL .
func WithTgChannelURL(tgChannelURL string) Option {
	return func(s *Doctor) {
		s.tgChannelURL = tgChannelURL
	}
}

// WithInstURL .
func WithInstURL(instURL string) Option {
	return func(s *Doctor) {
		s.instURL = instURL
	}
}

// WithVkURL .
func WithVkURL(vkURL string) Option {
	return func(s *Doctor) {
		s.vkURL = vkURL
	}
}

// WithTikTokURL .
func WithTikTokURL(tikTokURL string) Option {
	return func(s *Doctor) {
		s.tiktokURL = tikTokURL
	}
}

// WithDzenURL .
func WithDzenURL(dzenURL string) Option {
	return func(s *Doctor) {
		s.dzenURL = dzenURL
	}
}

// WithYoutubeURL .
func WithYoutubeURL(youtubeURL string) Option {
	return func(s *Doctor) {
		s.youtubeURL = youtubeURL
	}
}

// WithSiteLink .
func WithSiteLink(siteLink string) Option {
	return func(s *Doctor) {
		s.siteLink = siteLink
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

// WithMainBlogTheme .
func WithMainBlogTheme(mainBlogTheme string) Option {
	return func(s *Doctor) {
		s.mainBlogTheme = mainBlogTheme
	}
}

// WithMedicalDirection .
func WithMedicalDirection(medicalDirection string) Option {
	return func(s *Doctor) {
		s.medicalDirection = medicalDirection
	}
}
