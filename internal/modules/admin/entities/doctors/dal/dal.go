package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
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

// GetDoctorByID получение доктора по ID
func (r *Repository) GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	sql := `
		select id,
			name,
			slug,
			email,
			tg_url,
			inst_url,
			dzen_url,
			tg_channel_url,
			youtube_url,
			vk_url,
			prodoctorov,
			tiktok_url,
			main_blog_theme,
			city_id,
			speciallity_id,
			medical_directions,
			date_created,
			birth_date,
			s3_image,
			is_kf_doctor,
			cooperation_type,
			is_active
		from docstar_site_doctor where id = $1`

	var doctorDAO dao.FullDoctorDAO
	err := pgxscan.Get(ctx, r.db, &doctorDAO, sql, doctorID)
	if err != nil {
		return nil, err
	}

	return doctorDAO.ToDomain(), nil
}

// GetDoctors получение всех докторов
func (r *Repository) GetDoctors(ctx context.Context) ([]*doctor.Doctor, error) {
	sql := `
		select id,
			name,
			s3_image,
			cooperation_type,
			is_active
		from docstar_site_doctor
		order by id desc 
	`

	var doctorsDAO dao.DoctorMiniatureList
	err := pgxscan.Select(ctx, r.db, &doctorsDAO, sql)
	if err != nil {
		return nil, err
	}

	return doctorsDAO.ToDomain(), nil
}

func (r *Repository) GetCityByID(ctx context.Context, cityID int64) (*city.City, error) {
	sql := `
		select id, name 
		from docstar_site_city 
		where id = $1
	`

	var cityDAO dao.CityDAO
	err := pgxscan.Get(ctx, r.db, &cityDAO, sql, cityID)
	if err != nil {
		return nil, err
	}

	return cityDAO.ToDomain(), nil
}

func (r *Repository) GetSpecialityByID(ctx context.Context, specialityID int64) (*speciality.Speciality, error) {
	sql := `
		select id, name 
		from docstar_site_speciallity 
		where id = $1
	`

	var specialityDAO dao.SpecialityDAO
	err := pgxscan.Get(ctx, r.db, &specialityDAO, sql, specialityID)
	if err != nil {
		return nil, err
	}

	return specialityDAO.ToDomain(), nil
}
