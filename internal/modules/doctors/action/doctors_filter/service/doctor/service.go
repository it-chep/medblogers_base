package doctor

import "context"

type Storage interface {
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetDoctors(ctx context.Context) error {
	return nil
}
