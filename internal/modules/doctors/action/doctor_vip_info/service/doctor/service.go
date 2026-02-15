package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/vip_card"
)

type CommonDal interface {
	GetDoctorInfo(ctx context.Context, slug string) (*doctor.Doctor, error)
}

type Dal interface {
	GetVipCardInfo(ctx context.Context, doctorID int64) (*vip_card.VipCard, error)
}

type Service struct {
	commonDal CommonDal
	dal       Dal
}

func New(commonDal CommonDal, dal Dal) *Service {
	return &Service{
		commonDal: commonDal,
		dal:       dal,
	}
}

// GetDoctorBySlug получение доктора по слагу
func (s *Service) GetDoctorBySlug(ctx context.Context, slug string) (*doctor.Doctor, error) {
	return s.commonDal.GetDoctorInfo(ctx, slug)
}

// GetDoctorVIPInfo получение информации о вип карточке
func (s *Service) GetDoctorVIPInfo(ctx context.Context, doctorID int64) (*vip_card.VipCard, error) {
	return s.dal.GetVipCardInfo(ctx, doctorID)
}
