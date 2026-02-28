package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dto"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
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

// todo мб стоит потом разбить на несколько запросов

func (r *Repository) GetRating(ctx context.Context) ([]dto.RatingItem, error) {
	sql := `
		select d.slug, d.name, d.s3_image,
		       c.id as city_id, c.name as city_name,
		       s.id as speciality_id, s.name as speciality_name,
		       COALESCE(SUM(m.mbc_count), 0) as mbc_coins
		from docstar_site_doctor d
			join mbc_operation m on m.doctor_id = d.id and m.occurred_at >= NOW() - INTERVAL '1 year'
			left join docstar_site_city c on c.id = d.city_id
			left join docstar_site_speciallity s on s.id = d.speciallity_id
		where d.is_active = true
		group by d.id, d.slug, d.name, d.s3_image, c.id, c.name, s.id, s.name
		having COALESCE(SUM(m.mbc_count), 0) > 0
		order by mbc_coins desc
	`

	var items []dao.RatingItemDAO
	err := pgxscan.Select(ctx, r.db, &items, sql)
	if err != nil {
		return nil, err
	}

	result := make([]dto.RatingItem, 0, len(items))
	for _, item := range items {
		result = append(result, dto.RatingItem{
			Slug:           item.Slug,
			Name:           item.Name,
			S3Image:        item.S3Image.String,
			CityID:         item.CityID.Int64,
			CityName:       item.CityName.String,
			SpecialityID:   item.SpecialityID.Int64,
			SpecialityName: item.SpecialityName.String,
			MBCCoins:       item.MBCCoins,
		})
	}

	return result, nil
}
