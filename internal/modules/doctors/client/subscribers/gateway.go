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
	"unicode/utf8"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/doctors/client/subscribers/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	pkgHttp "medblogers_base/internal/pkg/http"
)

const (
	defaultScheme = "http"
	secureScheme  = "https"
)

// Gateway в сервис subscribers
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
		Path:   fmt.Sprintf("/subscribers/%d/", int64(medblogersID)),
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
		TgSubsCount:      response.TgSubsCount,
		InstSubsCount:    response.InstSubsCount,
		YouTubeSubsCount: response.YoutubeSubsCount,
		VkSubsCount:      response.VkSubsCount,

		TgSubsCountText:      response.TgSubsCountText,
		InstSubsCountText:    response.InstSubsCountText,
		YouTubeSubsCountText: response.YoutubeSubsCountText,
		VkSubsCountText:      response.VkSubsCountText,

		TgLastUpdatedDate:      response.TgLastUpdatedDate,
		InstLastUpdatedDate:    response.InstLastUpdatedDate,
		YouTubeLastUpdatedDate: response.YoutubeLastUpdatedDate,
		VkLastUpdatedDate:      response.VkLastUpdatedDate,
	}, nil
}

func (g *Gateway) configureFilterRequest(request indto.GetDoctorsByFilterRequest) string {
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   "/doctors/filter/",
	}

	query := endpointURL.Query()

	// Добавляем параметры только если они не нулевые/пустые
	if !lo.Contains(request.SocialMedia, indto.All) && len(request.SocialMedia) > 0 && request.SocialMedia != nil {
		// Конвертируем enum SocialMedia в строковые значения
		var socials []string
		for _, sm := range request.SocialMedia {
			if sm.String() == "" {
				continue
			}
			socials = append(socials, sm.String())
		}
		socialMediaJSON, _ := json.Marshal(socials)
		query.Add("social_media", string(socialMediaJSON))
	}

	if request.MinSubscribers > 0 {
		query.Add("min_subscribers", strconv.FormatInt(request.MinSubscribers, 10))
	}

	if request.MaxSubscribers > 0 {
		query.Add("max_subscribers", strconv.FormatInt(request.MaxSubscribers, 10))
	}

	if request.Offset > 0 {
		query.Add("offset", strconv.FormatInt(request.Offset, 10))
	}

	if request.Limit > 0 {
		query.Add("limit", strconv.FormatInt(request.Limit, 10))
	}

	if utf8.RuneCountInString(request.Sort) != 0 {
		query.Add("sort", request.Sort)
	}

	endpointURL.RawQuery = query.Encode()

	return endpointURL.String()
}

// GetDoctorsByFilter - получение докторов по переданным фильтрам
func (g *Gateway) GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) (indto.GetDoctorsByFilterResponse, error) {
	logger.Message(ctx, "[GW subs] Получение подписчиков по фильтрам")

	endpointURL := g.configureFilterRequest(request)

	var response dto.GetDoctorsByFilterResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	result := make(map[int64]indto.GetDoctorsByFilterDoctor, len(response.Doctors))
	orderedIDs := make([]int64, 0, len(response.Doctors)) // тк мапа не дает гарантий упорядоченности, то делаем отсортированный список
	for _, doc := range response.Doctors {
		result[doc.Doctor.DoctorID] = indto.GetDoctorsByFilterDoctor{
			DoctorID:         doc.Doctor.DoctorID,
			TgSubsCount:      doc.Doctor.TgSubsCount,
			InstSubsCount:    doc.Doctor.InstSubsCount,
			YouTubeSubsCount: doc.Doctor.YouTubeSubsCount,
			VkSubsCount:      doc.Doctor.VkSubsCount,

			TgSubsCountText:      doc.Doctor.TgSubsCountText,
			InstSubsCountText:    doc.Doctor.InstSubsCountText,
			YouTubeSubsCountText: doc.Doctor.YouTubeSubsCountText,
			VkSubsCountText:      doc.Doctor.VkSubsCountText,
		}
		orderedIDs = append(orderedIDs, doc.Doctor.DoctorID)
	}

	return indto.GetDoctorsByFilterResponse{
		Doctors:          result,
		DoctorsCount:     response.DoctorsCount,
		SubscribersCount: response.SubscribersCount,
		OrderedIDs:       orderedIDs,
	}, nil
}

// GetDoctorsByFilterWithIDs - фильтрация врачей + doctors_IDs
func (g *Gateway) GetDoctorsByFilterWithIDs(ctx context.Context, request indto.GetDoctorsByFilterRequest, doctorsIDs []int64) (indto.GetDoctorsByFilterResponse, error) {
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   fmt.Sprintf("/doctors/filter/"),
	}

	var socials []string
	for _, sm := range request.SocialMedia {
		if sm.String() == "" {
			continue
		}
		socials = append(socials, sm.String())
	}

	reqBody, err := json.Marshal(dto.DoctorsFilterWithIDsRequest{
		SocialMedia:    socials,
		MaxSubscribers: request.MaxSubscribers,
		MinSubscribers: request.MinSubscribers,
		Limit:          request.Limit,
		CurrentPage:    request.Offset,
		Sort:           request.Sort,
		DoctorIDs:      doctorsIDs,
	})
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	var response dto.GetDoctorsByFilterResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL.String(), bytes.NewReader(reqBody))
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return indto.GetDoctorsByFilterResponse{}, err
	}

	result := make(map[int64]indto.GetDoctorsByFilterDoctor, len(response.Doctors))
	orderedIDs := make([]int64, 0, len(response.Doctors)) // тк мапа не дает гарантий упорядоченности, то делаем отсортированный список
	for _, doc := range response.Doctors {
		result[doc.Doctor.DoctorID] = indto.GetDoctorsByFilterDoctor{
			DoctorID:         doc.Doctor.DoctorID,
			TgSubsCount:      doc.Doctor.TgSubsCount,
			InstSubsCount:    doc.Doctor.InstSubsCount,
			YouTubeSubsCount: doc.Doctor.YouTubeSubsCount,
			VkSubsCount:      doc.Doctor.VkSubsCount,

			TgSubsCountText:      doc.Doctor.TgSubsCountText,
			InstSubsCountText:    doc.Doctor.InstSubsCountText,
			YouTubeSubsCountText: doc.Doctor.YouTubeSubsCountText,
			VkSubsCountText:      doc.Doctor.VkSubsCountText,
		}
		orderedIDs = append(orderedIDs, doc.Doctor.DoctorID)
	}

	return indto.GetDoctorsByFilterResponse{
		Doctors:          result,
		DoctorsCount:     response.DoctorsCount,
		SubscribersCount: response.SubscribersCount,
		OrderedIDs:       orderedIDs,
	}, nil
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
			DoctorID:         doctorID,
			TgSubsCount:      doctorData.TgSubsCount,
			InstSubsCount:    doctorData.InstSubsCount,
			YouTubeSubsCount: doctorData.YoutubeSubsCount,
			VkSubsCount:      doctorData.VkSubsCount,

			TgSubsCountText:      doctorData.TgSubsCountText,
			InstSubsCountText:    doctorData.InstSubsCountText,
			YouTubeSubsCountText: doctorData.YoutubeSubsCountText,
			VkSubsCountText:      doctorData.VkSubsCountText,
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
		Path:   "/doctors/create/",
	}

	body, err := json.Marshal(dto.CreateDoctorRequest{
		DoctorID:  int64(medblogersID),
		Telegram:  request.Telegram,
		Instagram: request.Instagram,
		Youtube:   request.YouTube,
		Vk:        request.Vk,
	})
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL.String(), bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return 0, err
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
		Path:   fmt.Sprintf("/doctors/%d/", int(medblogersID)),
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
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return int64(resp.StatusCode), err
	}

	return int64(resp.StatusCode), nil
}

// CheckTelegramInBlackList проверяет телеграм на накрутки
func (g *Gateway) CheckTelegramInBlackList(ctx context.Context, telegram string) (bool, error) {
	var response dto.CheckTelegramInBlackListResponse
	endpointURL := &url.URL{
		Scheme: defaultScheme,
		Host:   g.host,
		Path:   "/doctors/check_telegram_in_blacklist/",
	}

	body, err := json.Marshal(dto.CheckTelegramInBlackListRequest{
		Telegram: telegram,
	})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL.String(), bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	return response.IsInBlackList, nil
}
