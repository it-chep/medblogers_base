package doctor

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
)

type CommonDal interface {
	GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
}

type Dal interface {
	UpdateDoctor(ctx context.Context, doctorID int64, req dto.UpdateRequest) error
}

type Service struct {
	commonDal CommonDal
	dal       Dal
}

func NewService(commonDal CommonDal, dal Dal) *Service {
	return &Service{
		commonDal: commonDal,
		dal:       dal,
	}
}

// GetDoctor получение существующего доктора
func (s *Service) GetDoctor(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	return s.commonDal.GetDoctorByID(ctx, doctorID)
}

// UpdateDoctor обновление доктора
func (s *Service) UpdateDoctor(ctx context.Context, doctorID int64, updateReq dto.UpdateRequest) error {
	return s.dal.UpdateDoctor(ctx, doctorID, updateReq)
}
