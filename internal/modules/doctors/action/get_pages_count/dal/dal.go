package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/dto"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
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

// FilterDoctors - **** Считаем без лимита, так как фильтрация идет по индексам и мы можем запылесосить всю базу **** //
func (r *Repository) FilterDoctors(ctx context.Context, filter dto.Filter) ([]int64, error) {
	logger.Message(ctx, "[Repo] Селект докторов из базы по фильтрам")
	// **** Считаем без лимита, так как фильтрация идет по индексам и мы можем запылесосить всю базу **** //
	sql, phValues := sqlStmt(filter)

	var doctors []dao.DoctorMiniatureDAO
	err := pgxscan.Select(ctx, r.db, &doctors, sql, phValues...)
	if err != nil {
		return nil, err
	}

	orderedIDs := make([]int64, 0, len(doctors))
	for _, doctorDAO := range doctors {
		orderedIDs = append(orderedIDs, doctorDAO.ID)
	}

	return orderedIDs, nil
}

// sqlStmt к-ор запроса
func sqlStmt(filter dto.Filter) (_ string, phValues []any) {
	defaultSql := `
	select
		d.id
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
		group by d.id
    `, defaultSql, whereStmtBuilder.String()), phValues
}
