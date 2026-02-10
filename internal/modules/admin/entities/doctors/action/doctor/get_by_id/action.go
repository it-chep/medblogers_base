package get_by_id

import (
	"context"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/service/cities"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/service/cooperation_type"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/service/image"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/service/specialities"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/service/subscribers"
	common_dal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/pipe"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
}

// Action активация доктора
type Action struct {
	commonDal               CommonDal
	subscribersEnricher     *subscribers.Service
	cityEnricher            *cities.Service
	specialityEnricher      *specialities.Service
	imageEnricher           *image.Service
	cooperationTypeEnricher *cooperation_type.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal:               common_dal.NewRepository(pool),
		subscribersEnricher:     subscribers.New(clients.Subscribers),
		cityEnricher:            cities.New(dal.NewRepository(pool)),
		specialityEnricher:      specialities.New(dal.NewRepository(pool)),
		imageEnricher:           image.New(clients.S3),
		cooperationTypeEnricher: cooperation_type.New(dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64) (_ *dto.DoctorDTO, err error) {
	doctorDomain, err := a.commonDal.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return
	}

	docDTO := dto.New(doctorDomain)

	docDTO, err = pipe.
		With(a.imageEnricher.Enrich).
		With(a.specialityEnricher.Enrich).
		With(a.cityEnricher.Enrich).
		With(a.subscribersEnricher.Enrich).
		With(a.cooperationTypeEnricher.Enrich).
		Run(ctx, docDTO).
		Get()
	if err != nil {
		return nil, err
	}

	return docDTO, nil
}
