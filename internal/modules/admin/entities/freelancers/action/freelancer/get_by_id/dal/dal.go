package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
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

func (r *Repository) GetCity(ctx context.Context, cityID int64) (*city.City, error) {
	sql := `
		select c.id, c.name 
		from freelancers_city c
		where c.id = $1
	`

	var cities dao.CityDAO
	err := pgxscan.Get(ctx, r.db, &cities, sql, cityID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &city.City{}, nil
		}
		return nil, err
	}
	return cities.ToDomain(), nil
}

func (r *Repository) GetSpeciality(ctx context.Context, specialityID int64) (*speciality.Speciality, error) {
	sql := `
		select s.id as "id", s.name as "name" 
		from freelancers_speciality s
		where s.id = $1
	`

	var specialities dao.SpecialityDAO
	err := pgxscan.Get(ctx, r.db, &specialities, sql, specialityID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &speciality.Speciality{}, nil
		}
		return nil, err
	}
	return specialities.ToDomain(), nil
}

func (r *Repository) GetCooperationType(ctx context.Context, cooperationTypeID int64) (*freelancer.CooperationType, error) {
	sql := `
		select c.id as "id", c.name as "name"
		from freelancers_cooperation_type c
		where c.id = $1
	`

	var cooperationType dao.CooperationTypeDAO
	err := pgxscan.Get(ctx, r.db, &cooperationType, sql, cooperationTypeID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &freelancer.CooperationType{}, nil
		}
		return nil, err
	}

	return cooperationType.ToDomain(), nil
}
