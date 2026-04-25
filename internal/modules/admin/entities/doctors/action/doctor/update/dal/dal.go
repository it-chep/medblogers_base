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
		inst_url = $3,
		vk_url = $4,
		dzen_url = $5,
		tg_url = $6,
		tg_channel_url = $7,
		youtube_url = $8,
		tiktok_url = $9,
		prodoctorov = $10,
		city_id = $11,
		speciallity_id = $12,
		main_blog_theme = $13,
		is_kf_doctor = $14,
		birth_date = $15,
		cooperation_type = $16,
		medical_directions = $17,
		marketing_preferences = $18,
		email = $19
   	where id = $1
	`

	args := []interface{}{
		doctorID,                 // $1
		req.Name,                 // $2
		req.InstURL,              // $3
		req.VkURL,                // $4
		req.DzenURL,              // $5
		req.TgURL,                // $6
		req.TgChannelURL,         // $7
		req.YouTubeURL,           // $8
		req.TikTokURL,            // $9
		req.SiteLink,             // $10
		req.MainCityID,           // $11
		req.MainSpecialityID,     // $12
		req.MainBlogTheme,        // $13
		req.IsKFDoctor,           // $14
		req.BirthDate,            // $15
		req.CooperationTypeID,    // $16
		req.MedicalDirections,    // $17
		req.MarketingPreferences, // $18
		req.Email,                // $19
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
