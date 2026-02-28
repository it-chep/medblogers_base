package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dto"
	city_dao "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	doctor_dao "medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetMBCAggregation(ctx context.Context) ([]dto.MBC, error) {
	q := `
		select m.doctor_id, COALESCE(SUM(m.mbc_count), 0) as mbc_coins
		from mbc_operation m
		where m.occurred_at >= NOW() - INTERVAL '1 year'
		group by m.doctor_id
		having COALESCE(SUM(m.mbc_count), 0) > 0
		order by mbc_coins desc
	`

	var items []doctor_dao.MBCAggDAO
	if err := pgxscan.Select(ctx, r.db, &items, q); err != nil {
		return nil, err
	}

	return lo.Map(items, func(item doctor_dao.MBCAggDAO, index int) dto.MBC {
		return dto.MBC{
			DoctorID: item.DoctorID,
			MBCCount: item.MBCCoins,
		}
	}), nil
}

func (r *Repository) GetDoctorsByIDs(ctx context.Context, ids []int64) ([]dto.RatingItem, error) {
	q := `
		select d.id, d.slug, d.name, d.s3_image, d.city_id, d.speciallity_id
		from docstar_site_doctor d
		where d.id = any($1::bigint[]) and d.is_active = true
	`

	var doctors []doctor_dao.RatingDoctorDAO
	if err := pgxscan.Select(ctx, r.db, &doctors, q, ids); err != nil {
		return nil, err
	}

	return lo.Map(doctors, func(item doctor_dao.RatingDoctorDAO, index int) dto.RatingItem {
		return dto.RatingItem{
			DoctorID:     item.ID,
			Slug:         item.Slug,
			Name:         item.Name,
			S3Image:      item.S3Image.String,
			CityID:       item.CityID.Int64,
			SpecialityID: item.SpecialityID.Int64,
		}
	}), nil
}

func (r *Repository) GetCitiesByIDs(ctx context.Context, ids []int64) ([]*city.City, error) {
	q := `
		select c.id, c.name
		from docstar_site_city c
		where c.id = any($1::bigint[])
	`

	var cities []city_dao.CityDAO
	if err := pgxscan.Select(ctx, r.db, &cities, q, ids); err != nil {
		return nil, err
	}

	return lo.Map(cities, func(item city_dao.CityDAO, _ int) *city.City {
		return item.ToDomain()
	}), nil
}

func (r *Repository) GetSpecialitiesByIDs(ctx context.Context, ids []int64) ([]*speciality.Speciality, error) {
	q := `
		select s.id, s.name
		from docstar_site_speciallity s
		where s.id = any($1::bigint[])
	`

	var specs []specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specs, q, ids); err != nil {
		return nil, err
	}

	return lo.Map(specs, func(item specialityDAO.SpecialityDAO, _ int) *speciality.Speciality {
		return item.ToDomain()
	}), nil
}
