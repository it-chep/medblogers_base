package subscribers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"medblogers_base/internal/pkg/logger"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"medblogers_base/internal/modules/doctors/client/subscribers/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

// todo indto переделать ?

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . HTTPClient

const (
	defaultScheme = "http"
	secureScheme  = "https"
)

// HTTPClient ...
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Gateway в сервис subscribers
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

// GetDoctorSubscribers - Получение количества подписчиков у доктора
func (g *Gateway) GetDoctorSubscribers(ctx context.Context, medblogersID doctor.MedblogersID) (indto.GetDoctorSubscribersResponse, error) {
	logger.Message(ctx, fmt.Sprintf("[GW subs] Получение подписчиков доктора %d", medblogersID))

	var response dto.GetDoctorSubscribersResponse

	if medblogersID == 0 {
		return indto.GetDoctorSubscribersResponse{}, errors.New("medblogersID is required")
	}
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   fmt.Sprintf("/subscribers/%d", int64(medblogersID)),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
	if err != nil {
		return indto.GetDoctorSubscribersResponse{}, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return indto.GetDoctorSubscribersResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return indto.GetDoctorSubscribersResponse{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return indto.GetDoctorSubscribersResponse{}, err
	}

	return indto.GetDoctorSubscribersResponse{
		TgSubsCount:         response.TgSubsCount,
		TgSubsCountText:     response.TgSubsCountText,
		TgLastUpdatedDate:   response.TgLastUpdatedDate,
		InstSubsCount:       response.InstSubsCount,
		InstSubsCountText:   response.InstSubsCountText,
		InstLastUpdatedDate: response.InstLastUpdatedDate,
	}, nil
}

func (g *Gateway) configureFilterRequest(request indto.GetDoctorsByFilterRequest) string {
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   "/doctors/filter/",
	}

	if request.MinSubscribers != 0 {

	}

	if request.MaxSubscribers != 0 {

	}

	if request.Limit != 0 {

	}

	if request.Offset != 0 {

	}

	if len(request.SocialMedia) != 0 {

	}

	return endpointURL.String()
}

// GetDoctorsByFilter - получение докторов по переданным фильтрам
func (g *Gateway) GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) (map[int64]indto.GetDoctorsByFilterResponse, error) {
	logger.Message(ctx, "[GW subs] Получение подписчиков по фильтрам")

	endpointURL := g.configureFilterRequest(request)

	var response dto.GetDoctorsByFilterResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return map[int64]indto.GetDoctorsByFilterResponse{}, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return map[int64]indto.GetDoctorsByFilterResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[int64]indto.GetDoctorsByFilterResponse{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return map[int64]indto.GetDoctorsByFilterResponse{}, err
	}

	result := make(map[int64]indto.GetDoctorsByFilterResponse, len(response.Doctors))
	for _, doc := range response.Doctors {
		result[doc.DoctorID] = indto.GetDoctorsByFilterResponse{
			DoctorID:          doc.DoctorID,
			TgSubsCount:       doc.TgSubsCount,
			TgSubsCountText:   doc.TgSubsCountText,
			InstSubsCount:     doc.InstSubsCount,
			InstSubsCountText: doc.InstSubsCountText,
		}
	}

	return result, nil
}

// GetSubscribersByDoctorIDs - получение количества подписчиков для миниатюр по переданным IDs
func (g *Gateway) GetSubscribersByDoctorIDs(ctx context.Context, medblogersIDs []int64) (map[int64]indto.GetSubscribersByDoctorIDsResponse, error) {
	logger.Message(ctx, "[GW subs] Получение подписчиков по ID докторов")
	var response dto.GetSubscribersByDoctorIDsResponse
	if len(medblogersIDs) == 0 {
		return nil, errors.New("medblogersIDs is required")
	}
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   "/doctors/by_ids/",
	}

	// Подготовка query параметров
	q := endpointURL.Query()
	idsStr := make([]string, 0, len(medblogersIDs))
	for _, id := range medblogersIDs {
		idsStr = append(idsStr, strconv.FormatInt(id, 10))
	}
	q.Set("doctor_ids", strings.Join(idsStr, ","))
	endpointURL.RawQuery = q.Encode()

	resultMap := make(map[int64]indto.GetSubscribersByDoctorIDsResponse, len(medblogersIDs))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
	if err != nil {
		return resultMap, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return resultMap, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resultMap, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return resultMap, err
	}

	for doctorID, doctorData := range response.Data {
		resultMap[doctorID] = indto.GetSubscribersByDoctorIDsResponse{
			DoctorID:          doctorID,
			TgSubsCount:       doctorData.TgSubsCount,
			TgSubsCountText:   doctorData.TgSubsCountText,
			InstSubsCount:     doctorData.InstSubsCount,
			InstSubsCountText: doctorData.InstSubsCountText,
		}
	}

	return resultMap, nil
}

// GetAllSubscribersInfo - получение информации об общем количестве подписчиков
func (g *Gateway) GetAllSubscribersInfo(ctx context.Context) (indto.GetAllSubscribersInfoResponse, error) {
	logger.Message(ctx, "[GW subs] Получение данных об общем количестве подписчиков")

	var response dto.GetAllSubscribersInfoResponse
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   "/subscribers/count/",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
	if err != nil {
		return indto.GetAllSubscribersInfoResponse{}, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return indto.GetAllSubscribersInfoResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return indto.GetAllSubscribersInfoResponse{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return indto.GetAllSubscribersInfoResponse{}, err
	}

	return indto.GetAllSubscribersInfoResponse{
		SubscribersCount:     response.SubscribersCount,
		SubscribersCountText: response.SubscribersCountText,
		LastUpdated:          response.LastUpdated,
	}, nil
}

// GetFilterInfo - получение информации о доступных фильтрах
func (g *Gateway) GetFilterInfo(ctx context.Context) ([]indto.FilterInfoResponse, error) {
	logger.Message(ctx, "[GW subs] Получение информации о доступных фильтрах")

	var response dto.FilterInfoResponse
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   "/filter/info/",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
	if err != nil {
		return []indto.FilterInfoResponse{}, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return []indto.FilterInfoResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []indto.FilterInfoResponse{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []indto.FilterInfoResponse{}, err
	}

	result := make([]indto.FilterInfoResponse, 0, len(response.Messengers))
	for _, messenger := range response.Messengers {
		result = append(result, indto.FilterInfoResponse{
			Name: messenger.Name,
			Slug: messenger.Slug,
		})
	}
	return result, nil
}

// CreateDoctor - создание врача в сервисе подписчиков
func (g *Gateway) CreateDoctor(ctx context.Context, medblogersID doctor.MedblogersID, request indto.CreateDoctorRequest) (int64, error) {
	if medblogersID == 0 {
		return 0, errors.New("medblogersID is required")
	}
	logger.Message(ctx, fmt.Sprintf("[GW subs] Создание доктора в сервисе подписчиков %d, параметры: %v", medblogersID, request))

	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   fmt.Sprintf("/doctors/%d", int(medblogersID)),
	}

	body, err := json.Marshal(dto.UpdateDoctorRequest{
		Telegram:  request.Telegram,
		Instagram: request.Instagram,
	})
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL.String(), bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return int64(resp.StatusCode), err
	}

	return int64(resp.StatusCode), nil
}

// UpdateDoctor - обновление врача в сервисе подписчиков
func (g *Gateway) UpdateDoctor(ctx context.Context, medblogersID doctor.MedblogersID, request indto.UpdateDoctorRequest) (int64, error) {
	if medblogersID == 0 {
		return 0, errors.New("medblogersID is required")
	}
	logger.Message(ctx, fmt.Sprintf("[GW subs] Обновление доктора в сервисе подписчиков %d, параметры: %v", medblogersID, request))

	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   fmt.Sprintf("/doctors/%d", int(medblogersID)),
	}

	body, err := json.Marshal(dto.UpdateDoctorRequest{
		Telegram:  request.Telegram,
		Instagram: request.Instagram,
	})
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpointURL.String(), bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return int64(resp.StatusCode), err
	}

	return int64(resp.StatusCode), nil
}
