package dal

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
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

func (r *Repository) CreateDoctor(ctx context.Context, createDTO dto.CreateDoctorRequest) (doctor.MedblogersID, error) {
	sql := `
		insert into docstar_site_doctor (
		                                 name, slug, email, inst_url, vk_url, dzen_url, tg_url, 
		                                 main_blog_theme, prodoctorov, city_id, speciallity_id, youtube_url, 
		                                 is_active, date_created, birth_date, tg_channel_url, tiktok_url
		                                 )
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, false, now(), $13, $14, $15)
		on conflict (name, email) do update set 
			slug = excluded.slug,
			inst_url = excluded.inst_url,
			vk_url = excluded.vk_url,
			dzen_url = excluded.dzen_url,
			tg_url = excluded.tg_url,
			main_blog_theme = excluded.main_blog_theme,
			prodoctorov = excluded.prodoctorov,
			city_id = excluded.city_id,
			speciallity_id = excluded.speciallity_id,
			youtube_url = excluded.youtube_url,
			birth_date = excluded.birth_date,
			tg_channel_url = excluded.tg_channel_url,
			tiktok_url = excluded.tiktok_url
		returning id
	`

	args := []any{
		createDTO.FullName,
		createDTO.Slug,
		createDTO.Email,
		createDTO.InstagramUsername,
		createDTO.VKUsername,
		createDTO.DzenUsername,
		createDTO.TelegramUsername,
		createDTO.MainBlogTheme,
		createDTO.SiteLink,
		createDTO.CityID,
		createDTO.SpecialityID,
		createDTO.YoutubeUsername,
		createDTO.BirthDateTime,
		createDTO.TelegramChannel,
		createDTO.TikTokURL,
	}

	var doctorID int64
	err := r.db.QueryRow(ctx, sql, args...).Scan(&doctorID)
	if err != nil {
		return 0, err
	}

	return doctor.MedblogersID(doctorID), nil
}

func (r *Repository) CreateAdditionalCities(ctx context.Context, medblogersID doctor.MedblogersID, citiesIDs []int64) error {
	if len(citiesIDs) == 0 {
		return nil
	}

	sql := `
	insert into docstar_site_doctor_additional_cities (doctor_id, city_id)
		select $1, unnest($2::bigint[])
		on conflict (doctor_id, city_id) do nothing`

	_, err := r.db.Exec(ctx, sql, medblogersID, citiesIDs)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateAdditionalSpecialities(ctx context.Context, medblogersID doctor.MedblogersID, specialitiesIDs []int64) error {
	if len(specialitiesIDs) == 0 {
		return nil
	}

	sql := `
	insert into docstar_site_doctor_additional_specialties (doctor_id, speciallity_id)
		select $1, unnest($2::bigint[])
		on conflict (doctor_id, speciallity_id) do nothing`

	_, err := r.db.Exec(ctx, sql, medblogersID, specialitiesIDs)
	if err != nil {
		return err
	}

	return nil
}
