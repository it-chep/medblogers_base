package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	GetDoctorInfo(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetDoctorInfo(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	return s.storage.GetDoctorInfo(ctx, doctorID)
}
