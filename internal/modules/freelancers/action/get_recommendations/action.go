package get_recommendations

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/dal"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/dto"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/service/doctors"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/service/freelancer"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	freelancerService *freelancer.Service
	doctorService     *doctors.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		doctorService:     doctors.NewService(dal.NewRepository(pool), clients.S3),
		freelancerService: freelancer.New(dal.NewRepository(pool)),
	}
}

// 1. Получаем фрилансера
// 2. Получаем его рекомендации
// 3. Получаем инфо о докторах из рекомендаций
// 4. Обогащаем инфой о городах, специальностях и фотках
// todo вообще в идеале не мешать докторов и фрилансеров и тут был бы крут API Gateway

func (a *Action) Do(ctx context.Context, freelancerID int64) (dto.Response, error) {
	recommendations, err := a.freelancerService.GetRecommendations(ctx, freelancerID)
	if err != nil {
		return dto.Response{}, err
	}

	docs, orderedIDs, err := a.doctorService.GetDoctorsInfo(ctx, recommendations.DoctorIDs())
	if err != nil {
		return dto.Response{}, err
	}

	a.doctorService.EnrichFacade(ctx, docs, orderedIDs)

	return dto.NewResponse(docs, orderedIDs), nil
}
