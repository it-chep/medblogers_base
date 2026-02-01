package doctor_vip_info

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctor_vip_info/dal"
	"medblogers_base/internal/modules/doctors/action/doctor_vip_info/service/doctor"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal"
	"medblogers_base/internal/modules/doctors/domain/vip_card"
	"medblogers_base/internal/pkg/postgres"
)

// Action ..
type Action struct {
	doctorService *doctor.Service
}

// New ..
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		doctorService: doctor.New(doctor_dal.NewRepository(pool), dal.NewRepository(pool)),
	}
}

// Do получение информации о вип карточке врача
func (a *Action) Do(ctx context.Context, doctorSlug string) (*vip_card.VipCard, error) {
	doc, err := a.doctorService.GetDoctorBySlug(ctx, doctorSlug)
	if err != nil {
		return nil, err
	}

	vipInfo, err := a.doctorService.GetDoctorVIPInfo(ctx, int64(doc.GetID()))
	if err != nil {
		return nil, err
	}

	return vipInfo, nil
}
