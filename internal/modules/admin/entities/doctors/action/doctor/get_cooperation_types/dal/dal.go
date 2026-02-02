package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
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

// GetDoctorCooperationTypes получение всех типов сотрудничества
func (r *Repository) GetDoctorCooperationTypes(ctx context.Context) ([]*doctor.CooperationType, error) {
	sql := `select id, name from doctors_cooperation_type`

	var cooperationTypes dao.CooperationTypes
	err := pgxscan.Select(ctx, r.db, &cooperationTypes, sql)
	if err != nil {
		return nil, err
	}

	return cooperationTypes.ToDomain(), nil
}
