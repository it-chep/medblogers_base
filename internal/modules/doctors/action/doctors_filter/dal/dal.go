package dal

import (
	"context"
	"fmt"
	consts "medblogers_base/internal/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	cityDAO "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"strings"

	"github.com/lib/pq"

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

func (r *Repository) GetDoctors(ctx context.Context, currentPage int64) (map[doctor.MedblogersID]*doctor.Doctor, []int64, error) {
	logger.Message(ctx, "[Repo] Селект докторов из базы без фильтров")
	sql := `
		select 
			id, name, slug, inst_url, city_id, speciallity_id, tg_channel_url, s3_image, is_kf_doctor, youtube_url
		from docstar_site_doctor d
		where d.is_active = true
		order by d.name asc 
		limit $1 
		offset $2
	`

	var doctors []dao.DoctorMiniatureDAO
	offset := (currentPage - 1) * consts.LimitDoctorsOnPage

	err := pgxscan.Select(ctx, r.db, &doctors, sql, consts.LimitDoctorsOnPage, offset)
	if err != nil {
		return nil, nil, err
	}

	result := make(map[doctor.MedblogersID]*doctor.Doctor, len(doctors))
	orderedIDs := make([]int64, 0, len(doctors))
	for _, doctorDAO := range doctors {
		result[doctor.MedblogersID(doctorDAO.ID)] = doctorDAO.ToDomain()
		orderedIDs = append(orderedIDs, doctorDAO.ID)
	}

	return result, orderedIDs, nil
}

// FilterDoctors - **** Считаем без лимита, так как фильтрация идет по индексам и мы можем запылесосить всю базу **** //
func (r *Repository) FilterDoctors(ctx context.Context, filter dto.Filter) (map[doctor.MedblogersID]*doctor.Doctor, []int64, error) {
	logger.Message(ctx, "[Repo] Селект докторов из базы по фильтрам")
	// **** Считаем без лимита, так как фильтрация идет по индексам и мы можем запылесосить всю базу **** //
	sql, phValues := sqlStmt(filter)

	var doctors []dao.DoctorMiniatureDAO
	err := pgxscan.Select(ctx, r.db, &doctors, sql, phValues...)
	if err != nil {
		return nil, nil, err
	}

	result := make(map[doctor.MedblogersID]*doctor.Doctor, len(doctors))
	orderedIDs := make([]int64, 0, len(doctors))
	for _, doctorDAO := range doctors {
		result[doctor.MedblogersID(doctorDAO.ID)] = doctorDAO.ToDomain()
		orderedIDs = append(orderedIDs, doctorDAO.ID)
	}

	return result, orderedIDs, nil
}

func (r *Repository) GetDoctorsByIDs(ctx context.Context, currentPage int64, ids []int64) (map[doctor.MedblogersID]*doctor.Doctor, error) {
	logger.Message(ctx, "[Repo] Селект докторов из базы по IDs")
	sql := `
		select 
			id, name, slug, inst_url, city_id, speciallity_id, tg_channel_url, s3_image, is_kf_doctor, youtube_url
		from docstar_site_doctor d
		where d.is_active = true and d.id = any($1::bigint[])
	`
	var doctors []dao.DoctorMiniatureDAO

	err := pgxscan.Select(ctx, r.db, &doctors, sql, ids)
	if err != nil {
		return nil, err
	}

	result := make(map[doctor.MedblogersID]*doctor.Doctor, len(doctors))
	for _, doctorDAO := range doctors {
		result[doctor.MedblogersID(doctorDAO.ID)] = doctorDAO.ToDomain()
	}

	newRes := make(map[doctor.MedblogersID]*doctor.Doctor, consts.LimitDoctorsOnPage)
	counter := int64(0)
	for _, id := range ids {
		if counter == consts.LimitDoctorsOnPage {
			break
		}

		doc, ok := result[doctor.MedblogersID(id)]
		if !ok {
			continue // если доктор неактивен или его нет в базе докторов, чего не должно быть по сути
		}

		newRes[doctor.MedblogersID(id)] = doc
		counter++
	}

	return newRes, nil
}

// sqlStmt к-ор запроса
func sqlStmt(filter dto.Filter) (_ string, phValues []any) {
	defaultSql := `
	select
		d.id,
		d.name,
		d.slug,
		d.inst_url,
		d.city_id,
		d.speciallity_id,
		d.tg_channel_url,
		d.s3_image, 
		d.is_kf_doctor,
		d.youtube_url
	from
    	docstar_site_doctor d
	where 
    	d.is_active = true
        `

	whereStmtBuilder := strings.Builder{}
	phCounter := 1 // Счетчик для плейсхолдеров

	if len(filter.Cities) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
		and (
			d.city_id = any($%d::bigint[])
				or exists (select 1
						   from docstar_site_doctor_additional_cities dc
						   where dc.doctor_id = d.id
							 and dc.city_id = any($%d::bigint[]))
			)`, phCounter, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.Cities))
		phCounter++
	}

	if len(filter.Specialities) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
		and (
			d.speciallity_id = any($%d::bigint[])
				or exists (select 1
						   from docstar_site_doctor_additional_specialties ds
						   where ds.doctor_id = d.id
							 and ds.speciallity_id = any($%d::bigint[]))
			)`, phCounter, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.Specialities))
		phCounter++
	}

	// возвращаем для первой страницы
	return fmt.Sprintf(`
		%s
		%s
		group by d.id, d.name
        order by d.name asc
		offset 0
    `, defaultSql, whereStmtBuilder.String()), phValues
}
