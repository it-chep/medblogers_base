package dal

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/pkg/postgres"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CountFilterDoctors(ctx context.Context, filter dto.Filter) (int64, error) {
	queryByFilters, phValues := r.queryByFilters(filter)

	sql := fmt.Sprintf(`
			select count(*) from docstar_site_doctor d where d.is_active = true 
		   	%s
    	`, queryByFilters)

	var doctorsCount int64
	err := pgxscan.Get(ctx, r.db, &doctorsCount, sql, phValues...)
	if err != nil {
		return 0, err
	}

	return doctorsCount, nil
}

func (r *Repository) CountByFilterAndIDs(ctx context.Context, filter dto.Filter, doctorIDs []int64) (int64, error) {
	queryByFilters, phValues := r.queryByFilters(filter)

	doctorIDsPhCounter := len(phValues) + 1
	phValues = append(phValues, pq.Int64Array(doctorIDs))

	sql := fmt.Sprintf(`
			select count(*) from docstar_site_doctor d where d.is_active = true 
		   	%s and d.id = any($%d::bigint[])
    	`, queryByFilters, doctorIDsPhCounter)

	var doctorsCount int64
	err := pgxscan.Get(ctx, r.db, &doctorsCount, sql, phValues...)
	if err != nil {
		return 0, err
	}

	return doctorsCount, nil
}

func (r *Repository) queryByFilters(filter dto.Filter) (string, []any) {
	whereStmtBuilder := strings.Builder{}
	phValues := []any{}
	phCounter := 1

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

	return whereStmtBuilder.String(), phValues
}
