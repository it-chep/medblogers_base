package doctor_dal

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
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

func (r Repository) GetDoctorsCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) as doctors_count
		from docstar_site_doctor d
		where d.is_active = true
	`

	var count int64
	if err := pgxscan.Get(ctx, r.db, &count, sql); err != nil {
		return 0, err
	}

	return count, nil
}

// GetDoctorInfo получает информацию о докторе
func (r Repository) GetDoctorInfo(ctx context.Context, slug string) (*doctor.Doctor, error) {
	sql := `
		select 
			id, name, slug, 
			inst_url, vk_url, dzen_url, tg_url, youtube_url, prodoctorov, tg_channel_url, tiktok_url, 
			s3_image, is_active, medical_directions, main_blog_theme, 
			city_id, speciallity_id, is_kf_doctor
		from docstar_site_doctor
		where slug = $1
	`

	var doctorDAO dao.DoctorDAO
	err := pgxscan.Get(ctx, r.db, &doctorDAO, sql, slug)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, fmt.Errorf("doctor with slug %s not found", slug)
	case err != nil:
		return nil, fmt.Errorf("failed to get doctor: %w", err)
	}

	return doctorDAO.ToDomain(), nil
}
