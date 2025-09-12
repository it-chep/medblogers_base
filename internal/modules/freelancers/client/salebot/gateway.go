package salebot

import (
	"bytes"
	"context"
	"encoding/json"
	indto "medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"medblogers_base/internal/modules/freelancers/client/salebot/dto"
	"medblogers_base/internal/pkg/logger"
	"net/http"

	pkgHttp "medblogers_base/internal/pkg/http"
)

// Gateway в сервис нотификации
type Gateway struct {
	host   string
	client pkgHttp.Executor
}

// NewGateway - конструктор
func NewGateway(host string, client pkgHttp.Executor) *Gateway {
	return &Gateway{
		host:   host,
		client: client,
	}
}

func (g *Gateway) NotificatorCreateFreelancer(ctx context.Context, createDTO indto.CreateRequest, adminChatID int64) {
	logger.Message(ctx, "[Notificator] Уведомление о создании фрилансера")
	// Подготавливаем данные для запроса
	requestData := dto.CreateFreelancerNotification{
		URL:      createDTO.Slug,
		Name:     createDTO.Name,
		TgURL:    createDTO.TgUsername,
		Message:  dto.CreateFreelancerEvent,
		ClientID: adminChatID,
	}

	// Кодируем данные в JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		logger.Error(ctx, "Ошибка при декодировании в json", err)
		return
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
	defer func(resp *http.Response) {
		if resp == nil {
			return
		}
		resp.Body.Close()
	}(resp)

	if resp != nil && resp.StatusCode == http.StatusOK {
		logger.Message(ctx, "[Notificator] Уведомление успешно отправлено")
	}
}
