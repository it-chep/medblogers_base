package freelancer_dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"strings"
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

	if filter.ExperienceWithDoctors != nil {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
			 and f.is_worked_with_doctors = $%d
		`, phCounter))
		phValues = append(phValues, lo.FromPtr(filter.ExperienceWithDoctors))
		phCounter++
	}

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

	// возвращаем для первой страницы
	return fmt.Sprintf(`
		%s
		%s
		group by f.id
    `, defaultSql, whereStmtBuilder.String()), phValues
}
