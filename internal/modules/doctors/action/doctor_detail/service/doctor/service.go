package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

// ТЕСТОВ НЕТ ТК ТУТ ПРОСТО КРУД ЗАПРОС И НЕТ ЛОГИКИ

// Storage .
type Storage interface {
	GetDoctorInfo(ctx context.Context, slug string) (*doctor.Doctor, error)
}

// Service сервис получения данных о докторе
type Service struct {
	storage Storage
}

// New к-ор
func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// GetDoctorInfo получение информации о докторе
func (s *Service) GetDoctorInfo(ctx context.Context, slug string) (*doctor.Doctor, error) {
	return s.storage.GetDoctorInfo(ctx, slug)
}
