package get_doctor_additional_specialities

import "medblogers_base/internal/pkg/postgres"

type Dal interface {
}

type Action struct {
	Dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		Dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*freelancer.Freelancer, error) {
	return a.commonDal.GetFreelancers(ctx)
}
