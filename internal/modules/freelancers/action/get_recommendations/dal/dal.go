package dal

import (
	"context"
	cityDAO "medblogers_base/internal/modules/freelancers/dal/doctor_city_dal/dao"
	daoDoctor "medblogers_base/internal/modules/freelancers/dal/doctor_dal/dao"
	specialityDAO "medblogers_base/internal/modules/freelancers/dal/doctor_speciality_dal/dao"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_recommendation/dao"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/doctor"
	domain "medblogers_base/internal/modules/freelancers/domain/freelancer_recommendation"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
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

func (r *Repository) GetRecommendations(ctx context.Context, freelancerID int64) (domain.FreelancerRecommendations, error) {
	sql := `select * from freelancer_recommendation where freelancer_id = $1`

	var recommendations dao.FreelancerRecommendationsDAO
	err := pgxscan.Select(ctx, r.db, &recommendations, sql, freelancerID)
	if err != nil {
		return nil, err
	}

	return recommendations.ToDomain(), nil
}

func (r *Repository) GetDoctorsInfo(ctx context.Context, doctorIDs []int64) ([]*doctor.Doctor, error) {
	sql := `
		select id, name, slug, city_id, speciallity_id, s3_image 
		from docstar_site_doctor
		where id = any($1)
	`

	var doctors daoDoctor.DoctorMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &doctors, sql, doctorIDs)
	if err != nil {
		return nil, err
	}

	return doctors.ToDomain(), nil
}

// GetDoctorAdditionalCities получение информации о городах доктора
func (r *Repository) GetDoctorAdditionalCities(ctx context.Context, medblogersIDs []int64) (map[int64][]*city.City, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных городов доктора")
	sql := `
	   select c.id, c.name, dc.doctor_id as "doctor_id"
	   from docstar_site_doctor_additional_cities dc
	       join docstar_site_city c ON dc.city_id = c.id
	   where dc.doctor_id = any($1::bigint[])
	   order by dc.doctor_id, c.name
	`

	var cities []*cityDAO.CityDAOWithDoctorID
	if err := pgxscan.Select(ctx, r.db, &cities, sql, medblogersIDs); err != nil {
		return nil, err
	}

	result := make(map[int64][]*city.City, len(cities))
	for _, c := range cities {
		if _, exists := result[c.DoctorID]; !exists {
			result[c.DoctorID] = make([]*city.City, 0)
		}
		result[c.DoctorID] = append(result[c.DoctorID], c.ToDomain())
	}

	return result, nil
}

// GetDoctorAdditionalSpecialities получение информации о специальностях доктора
func (r *Repository) GetDoctorAdditionalSpecialities(ctx context.Context, medblogersIDs []int64) (map[int64][]*speciality.Speciality, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных специальностей доктора")
	sql := `
	   select s.id, s.name, ds.doctor_id as "doctor_id"
	   from docstar_site_doctor_additional_specialties ds
	       join docstar_site_speciallity s ON ds.speciallity_id = s.id
	   where ds.doctor_id = any($1::bigint[])
	   order by ds.doctor_id, s.name
	`

	var specialities []*specialityDAO.SpecialityDAOWithDoctorID
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, medblogersIDs); err != nil {
		return nil, err
	}

	result := make(map[int64][]*speciality.Speciality, len(specialities))
	for _, s := range specialities {
		if _, exists := result[s.DoctorID]; !exists {
			result[s.DoctorID] = make([]*speciality.Speciality, 0)
		}
		result[s.DoctorID] = append(result[s.DoctorID], s.ToDomain())
	}

	return result, nil
}
