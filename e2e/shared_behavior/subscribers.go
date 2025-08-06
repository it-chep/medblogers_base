package shared_behavior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	"io"
	"medblogers_base/internal/modules/doctors/client/subscribers/dto"
	"medblogers_base/internal/pkg/http/mocks"
	"net/http"
	"strings"
)

// ExpectGetAllSubscribersInfo - мок для HTTP-ручки получения информации о подписчиках
var ExpectGetAllSubscribersInfo = func(httpClient *mocks.MockExecutor, count, countText, lastUpdated string) {
	expectGetAllSubscribersInfo(httpClient, count, countText, lastUpdated)
}

func expectGetAllSubscribersInfo(httpClient *mocks.MockExecutor, count, countText, lastUpdated string) {
	By("Подготовка мока для получения информации о подписчиках")

	expectedResponse := dto.GetAllSubscribersInfoResponse{
		SubscribersCount:     count,
		SubscribersCountText: countText,
		LastUpdated:          lastUpdated,
	}

	responseBody, _ := json.Marshal(expectedResponse)

	httpClient.EXPECT().
		Do(gomock.Any()).
		DoAndReturn(func(req *http.Request) (*http.Response, error) {
			if !strings.HasSuffix(req.URL.Path, "/subscribers/count/") {
				return nil, fmt.Errorf("unexpected path: %s", req.URL.Path)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(responseBody)),
			}, nil
		}).
		AnyTimes()
}

var ExpectGetFilterInfoSuccess = func(httpClient *mocks.MockExecutor, filtersResponse dto.FilterInfoResponse) {
	By("Подготовка мока для получения информации о доступных фильтрах по подписчикам")

	responseBody, _ := json.Marshal(filtersResponse)

	httpClient.EXPECT().
		Do(gomock.Any()).
		DoAndReturn(func(req *http.Request) (*http.Response, error) {
			if !strings.HasSuffix(req.URL.Path, "/filter/info/") {
				return nil, fmt.Errorf("unexpected path: %s", req.URL.Path)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(responseBody)),
			}, nil
		}).
		AnyTimes()
}

var ExpectGetFilterInfoError = func(httpClient *mocks.MockExecutor, err error) {
	By("Подготовка ошибочного ответа по получении информации о доступных фильтрах по подписчикам")

	httpClient.EXPECT().
		Do(gomock.Any()).
		DoAndReturn(func(req *http.Request) (*http.Response, error) {
			if !strings.HasSuffix(req.URL.Path, "/filter/info/") {
				return nil, fmt.Errorf("unexpected path: %s", req.URL.Path)
			}

			return &http.Response{
				StatusCode: http.StatusInternalServerError,
			}, err
		}).
		AnyTimes()
}
