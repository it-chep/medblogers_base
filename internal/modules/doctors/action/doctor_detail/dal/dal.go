package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository() *Repository {
	return &Repository{}
}

// GetDoctorInfo получает информацию о докторе
func (r Repository) GetDoctorInfo(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	sql := fmt.Sprintf(`
		select 
			id, name, slug, 
			inst_url, vk_url, dzen_url, tg_url,youtube_url, prodoctorov, tg_channel_url, tiktok_url, 
			s3_image, cooperation_type, is_active, medical_directions, main_blog_theme, 
			city_id, speciallity_id
		from docstar_site_doctor
		where id = $1
	`)

	var doctorDAO dao.DoctorDAO
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &doctorDAO, sql, doctorID); err != nil {
		return nil, err
	}

	return doctorDAO.ToDomain(), nil
}
