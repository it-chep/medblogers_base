package checker

import (
	"context"
	"regexp"
	"strings"
)

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
	validTelegram := cleanTelegramStringRegex(telegram)

	return s.storage.CheckTelegramInBlackList(ctx, validTelegram)
}

func cleanTelegramStringRegex(input string) string {
	if len(input) == 0 {
		return input
	}

	result := input

	re := regexp.MustCompile(`^(?:https?://)?(?:t\.me/)?@?/?`)
	result = re.ReplaceAllString(result, "")

	result = strings.Trim(result, "/ ")

	return strings.ToLower(result)
}
