package freelancer

import "strings"

func (f *Freelancer) GetID() int64 {
	return f.id
}

func (f *Freelancer) IsActive() bool {
	return f.isActive
}

func (f *Freelancer) HasExperienceWithDoctor() bool {
	return f.experienceWithDoctor
}

func (f *Freelancer) GetPriceCategory() int64 {
	return f.priceCategory
}

func (f *Freelancer) GetName() string {
	return f.name
}

func (f *Freelancer) GetSlug() string {
	return f.slug
}

func (f *Freelancer) GetEmail() string {
	return f.email
}

func (f *Freelancer) GetTgURL() string {
	if f.tgURL == "" {
		return ""
	}

	// Удаляем @ в начале, если есть
	tgURL := strings.TrimPrefix(f.tgURL, "@")

	// Если URL не содержит http/https, формируем полный URL
	if !strings.HasPrefix(tgURL, "http") {
		tgURL = "https://t.me/" + tgURL
	}

	return tgURL
}

func (f *Freelancer) GetPortfolioLink() string {
	return f.portfolioLink
}

func (f *Freelancer) GetCityID() int64 {
	return f.cityID
}

func (f *Freelancer) GetCityName() string {
	return f.cityName
}

func (f *Freelancer) GetAdditionalCitiesIDs() []int64 {
	return f.additionalCitiesIDs
}

func (f *Freelancer) GetSpecialityID() int64 {
	return f.specialityID
}

func (f *Freelancer) GetSpecialityName() string {
	return f.specialityName
}

func (f *Freelancer) GetAdditionalSpecialitiesIDs() []int64 {
	return f.additionalSpecialitiesIDs
}

func (f *Freelancer) GetSocialNetworks() []int64 {
	return f.socialNetworks
}

func (f *Freelancer) GetS3Image() string {
	return f.s3Image
}

func (f *Freelancer) GetMainCityID() int64 {
	return f.cityID
}

func (f *Freelancer) GetMainSpecialityID() int64 {
	return f.specialityID
}
