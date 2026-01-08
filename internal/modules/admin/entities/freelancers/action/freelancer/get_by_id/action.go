package get_by_id

import (
	"context"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/city"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/image"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/network"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/price_list"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/service/recommendation"
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
	commonDal      CommonDal
	recommendation *recommendation.Service
	priceList      *price_list.Service
	speciality     *speciality.Service
	network        *network.Service
	city           *city.Service
	image          *image.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repo := dal.NewRepository(pool)
	return &Action{
		commonDal:      commondal.NewRepository(pool),
		recommendation: recommendation.New(repo),
		priceList:      price_list.New(repo),
		speciality:     speciality.New(repo),
		network:        network.New(repo),
		city:           city.New(repo),
		image:          image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID int64) (*dto.FreelancerDTO, error) {
	frlncr, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return nil, err
	}

	freelancerDTO := dto.New(frlncr)

	err = pipe.With(a.image.Enrich).
		With(a.priceList.Enrich).
		With(a.speciality.Enrich).
		With(a.network.Enrich).
		With(a.city.Enrich).
		With(a.recommendation.Enrich).
		Run(ctx, freelancerDTO).
		Err()
	if err != nil {
		return nil, err
	}

	return freelancerDTO, nil
}
