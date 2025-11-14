package external

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Config,NotificationClient

type Config interface {
	GetCreateNotificationChatID() int64
}

type NotificationClient interface {
	NotificatorCreateFreelancer(ctx context.Context, createDTO dto.CreateRequest, adminChatID int64)
}

type Service struct {
	notificationClient NotificationClient
	config             Config
}

func NewService(notificationClient NotificationClient, config Config) *Service {
	return &Service{
		notificationClient: notificationClient,
		config:             config,
	}
}

func (s *Service) NotificatorAdmins(ctx context.Context, createDTO dto.CreateRequest) {
	s.notificationClient.NotificatorCreateFreelancer(ctx, createDTO, s.config.GetCreateNotificationChatID())
}
