package get_additional_specialities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_additional_specialities/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	GetAdditionalSpecialities(ctx context.Context, freelancerID int64) ([]*speciality.Speciality, error)
}

type Action struct {
	dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID int64) ([]*speciality.Speciality, error) {
	return a.dal.GetAdditionalSpecialities(ctx, freelancerID)
}
