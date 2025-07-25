package salebot

import (
	"bytes"
	"context"
	"encoding/json"
	indto "medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/client/salebot/dto"
	"medblogers_base/internal/pkg/logger"
	"net/http"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . HTTPClient

// HTTPClient ...
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Gateway в сервис нотификации
type Gateway struct {
	host   string
	client HTTPClient
}

// NewGateway - конструктор
func NewGateway(host string, client HTTPClient) *Gateway {
	return &Gateway{
		host:   host,
		client: client,
	}
}

func (g *Gateway) NotificatorCreateDoctor(ctx context.Context, createDTO indto.CreateDoctorRequest, adminChatID int64) {
	logger.Message(ctx, "[Notificator] Уведомление о создании доктора")
	// Подготавливаем данные для запроса
	requestData := dto.CreateDoctorNotification{
		URL:      createDTO.Slug,
		Name:     createDTO.FullName,
		InstURL:  createDTO.InstagramUsername,
		TgURL:    createDTO.TelegramUsername,
		Message:  dto.CreateDoctorEvent,
		ClientID: adminChatID,
	}

	// Кодируем данные в JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		logger.Error(ctx, "Ошибка при декодировании в json", err)
	}

	// Создаем HTTP запрос
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, g.host, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error(ctx, "Ошибка при формировании запроса", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := g.client.Do(req)
	if err != nil {
		logger.Error(ctx, "Ошибка при отправке запроса", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		logger.Message(ctx, "[Notificator] Уведомление успешно отправлено")
	}
}
