package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"
	"github.com/samber/lo"
	consts "medblogers_base/internal/dto"
	cityDAO "medblogers_base/internal/modules/freelancers/dal/city_dal/dao"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_dal/dao"
	socia
	socialDao "medblogers_base/internal/modules/freelancers/dal/society_dal/dao"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
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

// GetFreelancers получение фрилансеров без фильтров
func (r *Repository) GetFreelancers(ctx context.Context, filter *freelancer.Filter) (map[int64]*freelancer.Freelancer, []int64, error) {
	sql := `
	select 
	    id, name, slug, s3_image, price_category, is_worked_with_doctors
	from
    	freelancer f
	where 
    	f.is_active = true
	order by name
	limit $1
	offset $2
	`

	var offset int64
	if filter.Page > 1 {
		offset = (filter.Page - 1) * consts.LimitDoctorsOnPage
	}

	var frlncrs dao.Miniatures
	err := pgxscan.Select(ctx, r.db, &frlncrs, sql, consts.LimitDoctorsOnPage, offset)
	if err != nil {
		return nil, nil, err
	}

	result := make(map[int64]*freelancer.Freelancer, len(frlncrs))
	orderedIDs := make([]int64, 0, len(frlncrs))
	for _, frlncr := range frlncrs {
		result[frlncr.ID] = frlncr.ToDomain()
		orderedIDs = append(orderedIDs, frlncr.ID)
	}

	return result, orderedIDs, nil
}

// FilterFreelancers фильтрация
func (r *Repository) FilterFreelancers(ctx context.Context, filter freelancer.Filter) (map[int64]*freelancer.Freelancer, []int64, error) {
	logger.Message(ctx, "[Repo] Фильтруем фрилансеров")
	sql, phValues := sqlStmt(filter)

	sql, phValues = sqlAddLimitOffset(sql, phValues, filter)

	var frlncrs dao.Miniatures
	err := pgxscan.Select(ctx, r.db, &frlncrs, sql, phValues...)
	if err != nil {
		return nil, nil, err
	}

	result := make(map[int64]*freelancer.Freelancer, len(frlncrs))
	orderedIDs := make([]int64, 0, len(frlncrs))
	for _, frlncr := range frlncrs {
		result[frlncr.ID] = frlncr.ToDomain()
		orderedIDs = append(orderedIDs, frlncr.ID)
	}

	return result, orderedIDs, nil
}

func sqlAddLimitOffset(sql string, phValues []any, filter freelancer.Filter) (_ string, _ []any) {
	phCounter := len(phValues) + 1

	var offset int64
	if filter.Page > 1 {
		offset = (filter.Page - 1) * consts.LimitDoctorsOnPage
	}

	whereStmtBuilder := strings.Builder{}
	// limit
	whereStmtBuilder.WriteString(fmt.Sprintf(`
			limit $%d
		`, phCounter))
	phValues = append(phValues, consts.LimitDoctorsOnPage)
	phCounter++
	// offset
	whereStmtBuilder.WriteString(fmt.Sprintf(`
			limit $%d
		`, phCounter))
	phValues = append(phValues, offset)
	phCounter++

	return fmt.Sprintf(`
		%s
		%s
    `, sql, whereStmtBuilder.String()), phValues
}

// sqlStmt к-ор запроса
func sqlStmt(filter freelancer.Filter) (_ string, phValues []any) {
	defaultSql := `
	select 
	    id, name, slug, s3_image, price_category, is_worked_with_doctors
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

// GetAdditionalCities получение информации о городах доктора
func (r *Repository) GetAdditionalCities(ctx context.Context, medblogersIDs []int64) (map[int64][]*city.City, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных городов фрилансера")
	sql := `
	   select c.id, c.name, fc.freelancer_id as "freelancer_id"
	   from freelancer_city_m2m fc
	       join freelancers_city c ON fc.city_id = c.id
	   where fc.freelancer_id = any($1::bigint[])
	   order by fc.freelancer_id, c.name
	`

	var cities []*cityDAO.CityDAOWithFreelancerID
	if err := pgxscan.Select(ctx, r.db, &cities, sql, medblogersIDs); err != nil {
		return nil, err
	}

	result := make(map[int64][]*city.City, len(cities))
	for _, c := range cities {
		if _, exists := result[c.FreelancerID]; !exists {
			result[c.FreelancerID] = make([]*city.City, 0)
		}
		result[c.FreelancerID] = append(result[c.FreelancerID], c.ToDomain())
	}

	return result, nil
}

// GetAdditionalSpecialities получение информации о специальностях доктора
func (r *Repository) GetAdditionalSpecialities(ctx context.Context, medblogersIDs []int64) (map[int64][]*speciality.Speciality, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных специальностей фрилансера")
	sql := `
	   select s.id, s.name, fs.freelancer_id as "freelancer_id"
	   from freelancer_speciality_m2m fs
	       join freelancers_speciality s ON fs.speciality_id = s.id
	   where fs.freelancer_id = any($1::bigint[])
	   order by fs.freelancer_id, s.name
	`

	var specialities []*specialityDAO.SpecialityDAOWithFreelancerID
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, medblogersIDs); err != nil {
		return nil, err
	}

	result := make(map[int64][]*speciality.Speciality, len(specialities))
	for _, s := range specialities {
		if _, exists := result[s.FreelancerID]; !exists {
			result[s.FreelancerID] = make([]*speciality.Speciality, 0)
		}
		result[s.FreelancerID] = append(result[s.FreelancerID], s.ToDomain())
	}

	return result, nil
}

// GetSocialNetworks получение информации о cоц сетях
func (r *Repository) GetSocialNetworks(ctx context.Context, medblogersIDs []int64) (map[int64][]*social_network.SocialNetwork, error) {
	logger.Message(ctx, "[Dal] Получение соц.сетей фрилансера")
	sql := `
	   select s.id, s.name, fs.freelancer_id as "freelancer_id"
	   from freelancer_social_networks_m2m fs
	       join social_networks s ON fs.social_network_id = s.id
	   where fs.freelancer_id = any($1::bigint[])
	   order by fs.freelancer_id, s.name
	`

	var networks []socialDao.SocialNetworkWithFreelancerID
	if err := pgxscan.Select(ctx, r.db, &networks, sql, medblogersIDs); err != nil {
		return nil, err
	}

	result := make(map[int64][]*social_network.SocialNetwork, len(networks))
	for _, n := range networks {
		if _, exists := result[n.FreelancerID]; !exists {
			result[n.FreelancerID] = make([]*social_network.SocialNetwork, 0)
		}
		result[n.FreelancerID] = append(result[n.FreelancerID], n.ToDomain())
	}

	return result, nil
}
