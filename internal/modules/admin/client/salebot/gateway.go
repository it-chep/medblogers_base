package salebot

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/client/salebot/dto"
	pkgHttp "medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/logger"
	"net/http"
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

// NotificateError отправка уведомления об ошибке
func (g *Gateway) NotificateError(ctx context.Context, errText string, clientID int64) {
	requestData := dto.ErrorRequest{
		Message:  dto.ErrorEvent,
		Error:    errText,
		ClientID: clientID,
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
	if resp != nil && resp.StatusCode != http.StatusOK {
		logger.Message(ctx, "[Notificator] Ошибка при отправке уведомления")
	}
}

// MMNotification отправка сообщения об ММ
func (g *Gateway) MMNotification(ctx context.Context, clientID int64) error {
	requestData := dto.MMRequest{
		Message:  dto.MMEvent,
		ClientID: clientID,
	}
	// Кодируем данные в JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		logger.Error(ctx, "Ошибка при декодировании в json", err)
		return err
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
	if resp != nil && resp.StatusCode != http.StatusOK {
		return errors.New("Ошибка при отправке пуша об ММ")
	}

	return nil
}
