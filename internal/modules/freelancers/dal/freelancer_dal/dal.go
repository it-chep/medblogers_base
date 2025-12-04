package freelancer_dal

import (
	"context"
	"fmt"
	freelancerDao "medblogers_base/internal/modules/freelancers/dal/freelancer_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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

// GetFreelancersCount .
func (r *Repository) GetFreelancersCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) as count
		from freelancer f
		where f.is_active = true
	`

	var count int64
	if err := pgxscan.Get(ctx, r.db, &count, sql); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

// FreelancersCountByFilter - Считаем количество фрилансеров по фильтрам
func (r *Repository) FreelancersCountByFilter(ctx context.Context, filter freelancer.Filter) (int64, error) {
	logger.Message(ctx, "[Repo] Считаем количество фрилансеров по фильтрам")
	sql, phValues := sqlStmt(filter)

	var count int64
	err := pgxscan.Get(ctx, r.db, &count, sql, phValues...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

// sqlStmt к-ор запроса
func sqlStmt(filter freelancer.Filter) (_ string, phValues []any) {
	defaultSql := `
	select
		count(f.*) as count
	from
    	freelancer f
	where 
    	f.is_active = true
        `

	whereStmtBuilder := strings.Builder{}
	phCounter := 1 // Счетчик для плейсхолдеров

	if len(filter.Cities) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
		and (
			f.city_id = any($%d::bigint[])
				or exists (select 1
						   from freelancer_city_m2m fc
						   where fc.freelancer_id = f.id
							 and fc.city_id = any($%d::bigint[]))
			)`, phCounter, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.Cities))
		phCounter++
	}

	if len(filter.Specialities) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
		and (
			f.speciality_id = any($%d::bigint[])
				or exists (select 1
						   from freelancer_speciality_m2m fs
						   where fs.freelancer_id = f.id
							 and fs.speciality_id = any($%d::bigint[]))
			)`, phCounter, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.Specialities))
		phCounter++
	}

	if len(filter.SocialNetworks) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
			and exists (
				select 1 from freelancer_social_networks_m2m fs
				where fs.freelancer_id = f.id
				and fs.social_network_id = any($%d::bigint[]))
		`, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.SocialNetworks))
		phCounter++
	}

	if len(filter.PriceCategory) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
			and f.price_category = any($%d::bigint[])
		`, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.PriceCategory))
		phCounter++
	}

	return fmt.Sprintf(`
        %s
        %s
    `, defaultSql, whereStmtBuilder.String()), phValues
}

// GetFreelancerInfo детальная информация о фрилансере
func (r *Repository) GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error) {
	sql := `
		select id, slug, name, tg_username, portfolio_link, speciality_id, city_id, price_category, s3_image, agency_representative, start_working_date
		    from freelancer
		where slug = $1 and is_active = true
	`

	var fDao freelancerDao.FreelancerDetail
	err := pgxscan.Get(ctx, r.db, &fDao, sql, slug)
	if err != nil {
		return nil, err
	}
	return fDao.ToDomain(), nil
}
