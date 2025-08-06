package external

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Config,NotificationClient,SubscribersClient

type Config interface {
	GetCreateNotificationChatID() int64
}

type NotificationClient interface {
	NotificatorCreateDoctor(ctx context.Context, createDTO dto.CreateDoctorRequest, adminChatID int64)
}

type SubscribersClient interface {
	CreateDoctor(ctx context.Context, medblogersID doctor.MedblogersID, request indto.CreateDoctorRequest) (int64, error)
}

type Service struct {
	subscribersClient  SubscribersClient
	notificationClient NotificationClient
	config             Config
}

func NewService(subscribersClient SubscribersClient, notificationClient NotificationClient, config Config) *Service {
	return &Service{
		subscribersClient:  subscribersClient,
		notificationClient: notificationClient,
		config:             config,
	}
}

func (s *Service) NotificatorAdmins(ctx context.Context, createDTO dto.CreateDoctorRequest) {
	s.notificationClient.NotificatorCreateDoctor(ctx, createDTO, s.config.GetCreateNotificationChatID())
}

func (s *Service) SendToSubscribers(ctx context.Context, createDTO dto.CreateDoctorRequest) {
	_, err := s.subscribersClient.CreateDoctor(ctx, createDTO.ID, indto.CreateDoctorRequest{
		Telegram:  createDTO.TelegramChannel,
		Instagram: createDTO.InstagramUsername,
	})
	if err != nil {
		logger.Error(ctx, "[Create] Ошибка при создании врача в subscribers", err)
		return
	}
}
