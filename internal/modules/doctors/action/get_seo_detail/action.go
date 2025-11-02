package get_seo_detail

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/get_seo_detail/dal"
	"medblogers_base/internal/modules/doctors/action/get_seo_detail/dto"
	"medblogers_base/internal/modules/doctors/action/get_seo_detail/service/doctors"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	doctorsService *doctors.Service
}

func NewAction(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		doctorsService: doctors.New(dal.NewRepository(pool), clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, slug string) (dto.Response, error) {
	doctor, err := a.doctorsService.GetDoctorInfo(ctx, slug)
	if err != nil {
		logger.Error(ctx, "Ошибка получения доктора для SEO", err)
		return dto.Response{}, err
	}

	description, err := a.doctorsService.ConfigureDoctorDescription(ctx, doctor.GetID())
	if err != nil {
		logger.Error(ctx, "Ошибка получения описания для SEO", err)
		return dto.Response{}, err
	}

	return dto.Response{
		Description: description,
		Title:       doctor.GetName(),
		ImageURL:    a.doctorsService.GetDoctorImage(doctor.GetS3Key()),
	}, nil
}
