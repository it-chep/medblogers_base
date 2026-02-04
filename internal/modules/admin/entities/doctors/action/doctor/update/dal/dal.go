package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// UpdateDoctor обновление врача
func (r *Repository) UpdateDoctor(ctx context.Context, doctorID int64, req dto.UpdateRequest) error {
	sql := `
	update docstar_site_doctor set 
		name = $2,
		slug = $3,
		inst_url = $4,
		vk_url = $5,
		dzen_url = $6,
		tg_url = $7,
		tg_channel_url = $8,
		youtube_url = $9,
		tiktok_url = $10,
		prodoctorov = $11,
		city_id = $12,
		speciallity_id = $13,
		main_blog_theme = $14,
		is_kf_doctor = $15,
		birth_date = $16,
		cooperation_type = $17,
		medical_directions = $18,
		marketing_preferences = $19
   	where id = $1
	`

	args := []interface{}{
		doctorID,
		req.Name,
		req.Slug,
		req.InstURL,
		req.VkURL,
		req.DzenURL,
		req.TgURL,
		req.TgChannelURL,
		req.YouTubeURL,
		req.TikTokURL,
		req.SiteLink,
		req.MainCityID,
		req.MainSpecialityID,
		req.MainBlogTheme,
		req.IsKFDoctor,
		req.BirthDate,
		req.CooperationTypeID,
		req.MedicalDirections,
		req.MarketingPreferences,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
