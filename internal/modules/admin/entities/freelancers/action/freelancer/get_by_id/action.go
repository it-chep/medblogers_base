package get_by_id

import (
	"context"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/city"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/image"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/speciality"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/pipe"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type Action struct {
	commonDal  CommonDal
	speciality *speciality.Service
	city       *city.Service
	image      *image.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repo := dal.NewRepository(pool)
	return &Action{
		commonDal:  commondal.NewRepository(pool),
		speciality: speciality.New(repo),
		city:       city.New(repo),
		image:      image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID int64) (*dto.FreelancerDTO, error) {
	frlncr, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return nil, err
	}

	freelancerDTO := dto.New(frlncr)

	err = pipe.With(a.image.Enrich).
		With(a.speciality.Enrich).
		With(a.city.Enrich).
		Run(ctx, freelancerDTO).
		Err()
	if err != nil {
		return nil, err
	}

	return freelancerDTO, nil
}
