package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	cityDAO "medblogers_base/internal/modules/freelancers/dal/city_dal/dao"
	specialityDAO "medblogers_base/internal/modules/freelancers/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с фрилансерами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetFreelancerInfo получает информацию о докторе
func (r Repository) GetFreelancerInfo(ctx context.Context, slug string) (*doctor.Doctor, error) {
	sql := `
		select 
			id, name, slug, 
			inst_url, vk_url, dzen_url, tg_url, youtube_url, prodoctorov, tg_channel_url, tiktok_url, 
			s3_image, is_active, medical_directions, main_blog_theme, 
			city_id, speciallity_id
		from freelancer
		where slug = $1
	`

	var doctorDAO dao.DoctorDAO
	err := pgxscan.Get(ctx, r.db, &doctorDAO, sql, slug)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, fmt.Errorf("freelancer with slug %s not found", slug)
	case err != nil:
		return nil, fmt.Errorf("failed to get doctor: %w", err)
	}

	return doctorDAO.ToDomain(), nil
}

// GetFreelancerAdditionalCities получение информации о городах фрилансера
func (r Repository) GetFreelancerAdditionalCities(ctx context.Context, freelancerID int64) (map[int64]*city.City, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных городов фрилансера")
	sql := `
		select c.id, c.name
        from freelancers_city c
        inner join freelancer_city_m2m fc ON c.id = fc.city_id
        where fc.freelancer_id = $1
        order by c.name
	`

	var cities []*cityDAO.CityDAO
	if err := pgxscan.Select(ctx, r.db, &cities, sql, freelancerID); err != nil {
		return nil, err
	}

	result := make(map[int64]*city.City, len(cities))
	for _, c := range cities {
		result[c.ID] = c.ToDomain()
	}

	return result, nil
}

// GetFreelancerAdditionalSpecialities получение информации о специальностях фрилансера
func (r Repository) GetFreelancerAdditionalSpecialities(ctx context.Context, freelancerID int64) (map[int64]*speciality.Speciality, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных специальностей фрилансера")
	sql := `
		select s.id, s.name
		from freelancers_speciality s
		inner join freelancer_speciality_m2m fs ON s.id = fs.speciality_id
		where fs.freelancer_id = $1
        order by s.name
	`

	var specialities []*specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, freelancerID); err != nil {
		return nil, err
	}

	result := make(map[int64]*speciality.Speciality, len(specialities))
	for _, s := range specialities {
		result[s.ID] = s.ToDomain()
	}

	return result, nil
}
