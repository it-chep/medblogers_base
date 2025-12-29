package external

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/logger"
	"regexp"
	"strings"
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
	tgLink := validateTGChannelURL(createDTO.TelegramChannel)
	instagramLink := validateInstagramURL(createDTO.InstagramUsername)
	youTube := validateYouTube(createDTO.YoutubeUsername)

	_, err := s.subscribersClient.CreateDoctor(ctx, createDTO.ID, indto.CreateDoctorRequest{
		Telegram:  tgLink,
		Instagram: instagramLink,
		YouTube:   youTube,
	})

	if err != nil {
		logger.Error(ctx, "[Create] Ошибка при создании врача в subscribers", err)
		return
	}
}

func validateTGChannelURL(url string) string {
	if url == "" {
		return ""
	}

	// Если это ссылка на супергруппу/канал с +
	if strings.Contains(url, "https://t.me/+") {
		return url
	}

	patterns := []string{
		`https?://t\.me/([a-zA-Z0-9_]+)`, // https://t.me/username
		`t\.me/([a-zA-Z0-9_]+)`,          // t.me/username
		`@?([a-zA-Z0-9_]+)`,              // @username или username
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(strings.TrimSpace(url))
		if len(matches) > 1 {
			return matches[1]
		}
	}

	return strings.TrimSpace(url)
}

func validateInstagramURL(url string) string {
	if url == "" {
		return ""
	}

	patterns := []string{
		`https?://(?:www\.)?instagram\.com/([a-zA-Z0-9_.]+)/?`, // https://instagram.com/username/
		`(?:www\.)?instagram\.com/([a-zA-Z0-9_.]+)/?`,          //  instagram.com/username
		`@?([a-zA-Z0-9_.]+)`,                                   // @username или username
		`([a-zA-Z0-9_.]+)`,                                     // username как последний fallback
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(strings.TrimSpace(url))
		if len(matches) > 1 {
			return matches[1]
		}
	}

	return strings.TrimSpace(url)
}

func validateYouTube(url string) string {
	if url == "" {
		return ""
	}

	patterns := []string{
		`https?://(?:www\.)?youtube\.com/([a-zA-Z0-9_.]+)/?`, // https://youtube.com/username/
		`(?:www\.)?youtube\.com/([a-zA-Z0-9_.]+)/?`,          //  youtube.com/username
		`@?([a-zA-Z0-9_.]+)`,                                 // @username или username
		`([a-zA-Z0-9_.]+)`,                                   // username как последний fallback
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(strings.TrimSpace(url))
		if len(matches) > 1 {
			return matches[1]
		}
	}

	return strings.TrimSpace(url)
}
