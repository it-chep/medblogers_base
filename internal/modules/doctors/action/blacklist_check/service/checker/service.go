package checker

import "context"

// Storage ...
type Storage interface {
	CheckTelegramInBlackList(ctx context.Context, telegram string) (bool, error)
}

// Service ...
type Service struct {
	storage Storage
}

// NewService ...
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// Check ...
func (s *Service) Check(ctx context.Context, telegram string) (bool, error) {
	return s.storage.CheckTelegramInBlackList(ctx, telegram)
}
