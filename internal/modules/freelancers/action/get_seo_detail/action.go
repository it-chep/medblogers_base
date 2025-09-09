package get_seo_detail

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/get_seo_detail/dal"
	"medblogers_base/internal/modules/freelancers/action/get_seo_detail/dto"
	"medblogers_base/internal/modules/freelancers/action/get_seo_detail/service/freelancer"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	freelancerService *freelancer.Service
}

func NewAction(pool postgres.PoolWrapper) *Action {
	return &Action{
		freelancerService: freelancer.New(dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, slug string) (dto.Response, error) {
	frlcer, err := a.freelancerService.GetFreelancerInfo(ctx, slug)
	if err != nil {
		logger.Error(ctx, "Ошибка получения доктора для SEO", err)
		return dto.Response{}, err
	}

	description, err := a.freelancerService.ConfigureDoctorDescription(ctx, frlcer.GetID())
	if err != nil {
		logger.Error(ctx, "Ошибка получения описания для SEO", err)
		return dto.Response{}, err
	}

	return dto.Response{
		Description: description,
		Title:       frlcer.GetName(),
	}, nil

}
