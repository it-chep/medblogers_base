package freelancer

import (
	"encoding/json"
	"time"
)

func (f *Freelancer) Json() ([]byte, error) {
	data := map[string]interface{}{
		"id":                    f.id,
		"is_active":             f.isActive,
		"name":                  f.name,
		"slug":                  f.slug,
		"tg_url":                f.tgURL,
		"city_id":               f.cityID,
		"speciality_id":         f.specialityID,
		"s3_image":              f.s3Image,
		"cooperation_type":      f.cooperationType.id,
		"agency_representative": f.agencyRepresentative,
		"price_category":        f.priceCategory,
		"start_working":         f.startWorking.Format(time.DateTime),
		"portfolio_link":        f.portfolioLink,
	}

	return json.Marshal(data)
}
