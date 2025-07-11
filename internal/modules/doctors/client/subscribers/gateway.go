package subscribers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"medblogers_base/internal/modules/doctors/client/subscribers/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"net/http"
)

// todo indto переделать

// Gateway в сервис subscribers
type Gateway struct {
	client *http.Client
}

// NewGateway - конструктор
func NewGateway(client *http.Client) *Gateway {
	return &Gateway{
		client: client,
	}
}

// GetDoctorSubscribers - Получение количества подписчиков у доктора
func (g *Gateway) GetDoctorSubscribers(ctx context.Context, medblogersID doctor.MedblogersID) (indto.GetDoctorSubscribersResponse, error) {
	var response dto.GetDoctorSubscribersResponse

	if medblogersID == 0 {
		return indto.GetDoctorSubscribersResponse{}, errors.New("medblogersID is required")
	}

	endpointURL := fmt.Sprintf("/subscribers/%d", int64(medblogersID))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
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
	endpointURL := "/doctors/filter/"
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

	return endpointURL
}

// GetDoctorsByFilter - получение докторов по переданным фильтрам
func (g *Gateway) GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) ([]indto.GetDoctorsByFilterResponse, error) {
	endpointURL := g.configureFilterRequest(request)

	var response dto.GetDoctorsByFilterResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return []indto.GetDoctorsByFilterResponse{}, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return []indto.GetDoctorsByFilterResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []indto.GetDoctorsByFilterResponse{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []indto.GetDoctorsByFilterResponse{}, err
	}

	result := make([]indto.GetDoctorsByFilterResponse, 0, len(response.Doctors))
	for _, doc := range response.Doctors {
		result = append(result, indto.GetDoctorsByFilterResponse{
			DoctorID:          doc.DoctorID,
			TgSubsCount:       doc.TgSubsCount,
			TgSubsCountText:   doc.TgSubsCountText,
			InstSubsCount:     doc.InstSubsCount,
			InstSubsCountText: doc.InstSubsCountText,
		})
	}

	return result, nil
}

// GetSubscribersByDoctorIDs - получение количества подписчиков для миниатюр по переданным IDs
func (g *Gateway) GetSubscribersByDoctorIDs(
	ctx context.Context,
	medblogersIDs doctor.MedblogersIDs,
) (map[doctor.MedblogersID]indto.GetSubscribersByDoctorIDsResponse, error) {
	var response dto.GetSubscribersByDoctorIDsResponse
	if len(medblogersIDs) == 0 {
		return nil, errors.New("medblogersIDs is required")
	}

	endpointURL := fmt.Sprintf("/doctors/by_ids/?doctor_ids=%s", medblogersIDs.String())

	resultMap := make(map[doctor.MedblogersID]indto.GetSubscribersByDoctorIDsResponse, len(medblogersIDs))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
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
		resultMap[doctor.MedblogersID(doctorID)] = indto.GetSubscribersByDoctorIDsResponse{
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
	var response dto.GetAllSubscribersInfoResponse

	endpointURL := "/subscribers/count/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
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
	var response dto.FilterInfoResponse

	endpointURL := "/filter/info/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
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

	body, err := json.Marshal(dto.UpdateDoctorRequest{
		Telegram:  request.Telegram,
		Instagram: request.Instagram,
	})
	if err != nil {
		return 0, err
	}

	endpointURL := fmt.Sprintf("/doctors/%d", int64(medblogersID))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewReader(body))
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

	body, err := json.Marshal(dto.UpdateDoctorRequest{
		Telegram:  request.Telegram,
		Instagram: request.Instagram,
	})
	if err != nil {
		return 0, err
	}

	endpointURL := fmt.Sprintf("/doctors/%d", int64(medblogersID))
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpointURL, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return int64(resp.StatusCode), err
	}

	return int64(resp.StatusCode), nil
}
