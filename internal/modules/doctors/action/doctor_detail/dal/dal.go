package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/doctor"

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

// GetDoctorInfo получает информацию о докторе
func (r Repository) GetDoctorInfo(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	sql := `
		select 
			id, name, slug, 
			inst_url, vk_url, dzen_url, tg_url,youtube_url, prodoctorov, tg_channel_url, tiktok_url, 
			s3_image, cooperation_type, is_active, medical_directions, main_blog_theme, 
			city_id, speciallity_id
		from docstar_site_doctor
		where id = $1
	`

	var doctorDAO dao.DoctorDAO
	if err := pgxscan.Select(ctx, r.db, &doctorDAO, sql, doctorID); err != nil {
		return nil, err
	}

	return doctorDAO.ToDomain(), nil
}

//// GetCitiesByIDs получение информации о городах доктора
//func (r Repository) GetCitiesByIDs(ctx context.Context, citiesIDs []int64) ([]*city.City, error) {
//	sql := `
//		select c.id, c.name
//		from docstar_site_city c
//		where c.id = any($1)
//	`
//
//	var cityDAO cityDAO.CityDAO
//	if err := pgxscan.Select(ctx, r.db, &cityDAO, sql, citiesIDs); err != nil {
//		return nil, err
//	}
//
//	return cityDAO.ToDomain(), nil
//}

//// GetSpecialitiesByIDs получение информации о специальностях доктора
//func (r Repository) GetSpecialitiesByIDs(ctx context.Context, specialitiesIDs []int64) ([]*speciality.Speciality, error) {
//	sql := `
//		select s.id, s.name
//		from docstar_site_speciallity s
//		where s.id = any($1)
//	`
//
//}

const (
	manyToManyCity       = `select city_id from docstar_site_doctor_additional_cities where doctor_id = $1`
	manyToManySpeciality = `select speciallity_id from docstar_site_doctor_additional_specialties where doctor_id = $1`
) // todo
