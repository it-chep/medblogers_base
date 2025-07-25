package subscribers

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/client/subscribers/mocks"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"net/http"
	"testing"
)

type fields struct {
	http *mocks.MockHTTPClient
}

func p(t *testing.T) fields {
	f := fields{
		http: mocks.NewMockHTTPClient(gomock.NewController(t)),
	}
	return f
}
func TestGetDoctorSubscribers(t *testing.T) {
	t.Parallel()

	t.Run("успешное получение подписчиков доктора", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
            "telegram_short": "123",
            "telegram_text": "подписчика",
            "tg_last_updated_date": "11.05.2004",
            "instagram_short": "123",
            "instagram_text": "подписчика",
            "inst_last_updated_date": "11.05.2004"
        }`
		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody))}, nil)
		gw := NewGateway("", deps.http)

		expectedResult := indto.GetDoctorSubscribersResponse{
			TgSubsCount:         "123",
			TgSubsCountText:     "подписчика",
			TgLastUpdatedDate:   "11.05.2004",
			InstSubsCount:       "123",
			InstSubsCountText:   "подписчика",
			InstLastUpdatedDate: "11.05.2004",
		}
		result, err := gw.GetDoctorSubscribers(context.Background(), doctor.MedblogersID(1))

		require.NoError(t, err, "Не должно быть ошибки")
		require.NotNil(t, result, "Результат не должен быть nil")

		assert.Equal(t, expectedResult, result, "Результат не соответствует ожидаемому")
	})
}

func TestGetDoctorSubscribersByFilter(t *testing.T) {
	t.Parallel()

	t.Run("Успешное получение 1 доктора по фильтру тг", func(t *testing.T) {
		t.Parallel()
		deps := p(t)
		mockResponseBody := `{
            "doctors": [
                {
                    "doctor_id": 1,
                    "telegram_short": "123",
                    "telegram_text": "подписчика",
                    "inst_short": "123",
                    "inst_text": "подписчика"
                }
            ]
        }`

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
		}, nil)

		gw := NewGateway("", deps.http)
		expectedResult := map[int64]indto.GetDoctorsByFilterResponse{
			1: indto.GetDoctorsByFilterResponse{
				DoctorID:          1,
				TgSubsCount:       "123",
				TgSubsCountText:   "подписчика",
				InstSubsCount:     "123",
				InstSubsCountText: "подписчика",
			},
		}

		result, err := gw.GetDoctorsByFilter(context.Background(), indto.GetDoctorsByFilterRequest{
			SocialMedia: []indto.SocialMedia{
				indto.Telegram,
			},
			MaxSubscribers: 200,
			MinSubscribers: 100,
		})

		require.NoError(t, err, "Не должно быть ошибки")
		require.NotNil(t, result, "Результат не должен быть nil")
		require.Len(t, result, 1, "Должен вернуться один доктор")
		assert.Equal(t, expectedResult[1].DoctorID, result[1].DoctorID, "ID доктора не совпадает")
		assert.Equal(t, expectedResult[1].TgSubsCount, result[1].TgSubsCount, "Количество подписчиков TG не совпадает")
		assert.Equal(t, expectedResult[1].TgSubsCountText, result[1].TgSubsCountText, "Текст подписчиков TG не совпадает")
		assert.Equal(t, expectedResult[1].InstSubsCount, result[1].InstSubsCount, "Количество подписчиков Instagram не совпадает")
		assert.Equal(t, expectedResult[1].InstSubsCountText, result[1].InstSubsCountText, "Текст подписчиков Instagram не совпадает")
	})
}

func TestGetSubscribersByDoctorIDs(t *testing.T) {
	t.Parallel()

	t.Run("Успешное получение подписчиков по  ID докторов", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
            "doctors": [
                {
                    "doctor_id": 1,
                    "telegram_short": "123",
                    "telegram_text": "подписчика",
                    "inst_short": "123",
                    "inst_text": "подписчика"
                }
            ]
        }`

		// todo request нормальный сделать
		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
		}, nil)

		gw := NewGateway("", deps.http)
		expectedResult := map[int64]indto.GetSubscribersByDoctorIDsResponse{
			1: indto.GetSubscribersByDoctorIDsResponse{
				DoctorID:          1,
				TgSubsCount:       "123",
				TgSubsCountText:   "подписчика",
				InstSubsCount:     "123",
				InstSubsCountText: "подписчика",
			},
		}

		result, err := gw.GetSubscribersByDoctorIDs(context.Background(), []int64{1, 2})

		require.NoError(t, err, "Не должно быть ошибки")
		require.NotNil(t, result, "Результат не должен быть nil")
		require.Len(t, result, 1, "Должен вернуться один доктор")
		assert.Equal(t, expectedResult[1].InstSubsCount, result[1].InstSubsCount, "Количество подписчиков Instagram не совпадает")
		assert.Equal(t, expectedResult[1].InstSubsCountText, result[1].InstSubsCountText, "Текст подписчиков Instagram не совпадает")
	})

	// todo остальные кейсы
}

func TestGetAllSubscribersInfo(t *testing.T) {
	t.Parallel()

	t.Run("Получение общей информации по сервису подписчиков", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
			"subscribers_count": "123",
			"subscribers_count_text": "Подписчика",
			"last_updated": "11.05.2004"
        }`
		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
		}, nil)

		expectedResult := indto.GetAllSubscribersInfoResponse{
			SubscribersCount:     "123",
			SubscribersCountText: "Подписчика",
			LastUpdated:          "11.05.2004",
		}

		gw := NewGateway("", deps.http)
		result, err := gw.GetAllSubscribersInfo(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})
}

func TestGetFilterInfo(t *testing.T) {
	t.Parallel()

}

func TestCreateDoctor(t *testing.T) {
	t.Parallel()
}

func TestUpdateDoctor(t *testing.T) {
	t.Parallel()
}
