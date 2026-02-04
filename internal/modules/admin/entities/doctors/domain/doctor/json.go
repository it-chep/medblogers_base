package doctor

import (
	"encoding/json"
	"time"
)

func (d *Doctor) Json() ([]byte, error) {
	data := map[string]interface{}{
		"id":                    int64(d.medblogersID),
		"is_active":             d.isActive,
		"name":                  d.name,
		"slug":                  d.slug,
		"email":                 d.email,
		"inst_url":              d.instURL,
		"vk_url":                d.vkURL,
		"dzen_url":              d.dzenURL,
		"tg_url":                d.tgURL,
		"tg_channel_url":        d.tgChannelURL,
		"youtube_url":           d.youtubeURL,
		"tiktok_url":            d.tiktokURL,
		"site_link":             d.siteLink,
		"city_id":               int64(d.cityID),
		"speciality_id":         int64(d.specialityID),
		"medical_direction":     d.medicalDirection,
		"main_blog_theme":       d.mainBlogTheme,
		"s3_image":              d.s3Image,
		"birth_date":            d.birthDate.Format("2006-01-02"),
		"cooperation_type":      d.cooperationType.id,
		"is_kf_doctor":          d.isKFDoctor,
		"created_at":            d.createdAt.Format(time.RFC3339),
		"marketing_preferences": d.marketingPreferences,
	}

	return json.Marshal(data)
}
