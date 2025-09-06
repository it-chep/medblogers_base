package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count/dto"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"strings"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// FilterFreelancers - Считаем количество фрилансеров по фильтрам
func (r *Repository) FilterFreelancers(ctx context.Context, filter dto.Filter) (int64, error) {
	logger.Message(ctx, "[Repo] Считаем количество фрилансеров по фильтрам")
	sql, phValues := sqlStmt(filter)

	var count int64
	err := pgxscan.Get(ctx, r.db, &count, sql, phValues...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// sqlStmt к-ор запроса
func sqlStmt(filter dto.Filter) (_ string, phValues []any) {
	defaultSql := `
	select
		count(f.*) as count
	from
    	freelancer f
	where 
    	f.is_active = true and f.is_worked_with_doctors = $1
        `

	phValues = append(phValues, filter.ExperienceWithDoctors)

	whereStmtBuilder := strings.Builder{}
	phCounter := 2 // Счетчик для плейсхолдеров

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
			)
		`, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.SocialNetworks))
		phCounter++
	}

	if len(filter.PriceCategory) != 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`
			and f.price_category_id = any($%d::bigint[])
		`, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.PriceCategory))
		phCounter++
	}

	// возвращаем для первой страницы
	return fmt.Sprintf(`
		%s
		%s
		group by d.id
    `, defaultSql, whereStmtBuilder.String()), phValues
}
