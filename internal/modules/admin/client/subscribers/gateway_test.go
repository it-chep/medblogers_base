package subscribers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"medblogers_base/internal/modules/doctors/client/subscribers/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/http/mocks"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fields struct {
	http *mocks.MockExecutor
}

func p(t *testing.T) fields {
	f := fields{
		http: mocks.NewMockExecutor(gomock.NewController(t)),
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

	t.Run("Неполное получение данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
            "telegram_short": "123",
            "telegram_text": "подписчика",
            "tg_last_updated_date": "11.05.2004",
            "instagram_short": "0",
            "instagram_text": "",
            "inst_last_updated_date": ""
        }`
		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody))}, nil)
		gw := NewGateway("", deps.http)

		expectedResult := indto.GetDoctorSubscribersResponse{
			TgSubsCount:         "123",
			TgSubsCountText:     "подписчика",
			TgLastUpdatedDate:   "11.05.2004",
			InstSubsCount:       "0",
			InstSubsCountText:   "",
			InstLastUpdatedDate: "",
		}
		result, err := gw.GetDoctorSubscribers(context.Background(), doctor.MedblogersID(1))

		require.NoError(t, err, "Не должно быть ошибки")
		require.NotNil(t, result, "Результат не должен быть nil")

		assert.Equal(t, expectedResult, result, "Результат не соответствует ожидаемому")
	})

	t.Run("Ошибка при получении данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusInternalServerError}, errors.New("internal server error"))
		gw := NewGateway("", deps.http)

		expectedResult := indto.GetDoctorSubscribersResponse{}
		result, err := gw.GetDoctorSubscribers(context.Background(), doctor.MedblogersID(1))

		require.Error(t, err, "Должна быть ошибка")
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
					"doctor": {
						"doctor_id": 1,
						"telegram_short": "123",
						"telegram_text": "подписчика",
						"inst_short": "123",
						"inst_text": "подписчика"
					}
				}
            ]
        }`

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
		}, nil)

		gw := NewGateway("", deps.http)
		expectedResult := map[int64]indto.GetDoctorsByFilterDoctor{
			1: indto.GetDoctorsByFilterDoctor{
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
		assert.Equal(t, expectedResult[1].DoctorID, result.Doctors[1].DoctorID, "ID доктора не совпадает")
		assert.Equal(t, expectedResult[1].TgSubsCount, result.Doctors[1].TgSubsCount, "Количество подписчиков TG не совпадает")
		assert.Equal(t, expectedResult[1].TgSubsCountText, result.Doctors[1].TgSubsCountText, "Текст подписчиков TG не совпадает")
		assert.Equal(t, expectedResult[1].InstSubsCount, result.Doctors[1].InstSubsCount, "Количество подписчиков Instagram не совпадает")
		assert.Equal(t, expectedResult[1].InstSubsCountText, result.Doctors[1].InstSubsCountText, "Текст подписчиков Instagram не совпадает")
	})

	t.Run("Ошибка при получении данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusInternalServerError}, errors.New("internal server error"))
		gw := NewGateway("", deps.http)

		expectedResult := map[int64]indto.GetDoctorsByFilterDoctor{}
		result, err := gw.GetDoctorsByFilter(context.Background(), indto.GetDoctorsByFilterRequest{
			SocialMedia: []indto.SocialMedia{
				indto.Telegram,
			},
			MaxSubscribers: 200,
			MinSubscribers: 100,
		})

		require.Error(t, err, "Должна быть ошибка")
		require.NotNil(t, result, "Результат не должен быть nil")

		assert.Equal(t, expectedResult, result, "Результат не соответствует ожидаемому")
	})
}

func TestGetSubscribersByDoctorIDs(t *testing.T) {
	t.Parallel()

	t.Run("Успешное получение подписчиков по ID докторов", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
            "data": {
                "1": {
                    "doctor_id": 1,
                    "telegram_subs_count": "123",
                    "telegram_subs_text": "подписчика",
                    "instagram_subs_count": "456",
                    "instagram_subs_text": "подписчиков"
                },
                "2": {
                    "doctor_id": 2,
                    "telegram_subs_count": "789",
                    "telegram_subs_text": "подписчиков",
                    "instagram_subs_count": "101",
                    "instagram_subs_text": "подписчик"
                }
            }
        }`

		// Проверяем корректность запроса
		deps.http.EXPECT().
			Do(gomock.Any()).
			DoAndReturn(func(req *http.Request) (*http.Response, error) {
				r, err := url.QueryUnescape(req.URL.String())
				assert.NoError(t, err)
				assert.Contains(t, r, "doctor_ids=1,2", "URL должен содержать ID докторов")

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
					Header:     map[string][]string{"Content-Type": {"application/json"}},
				}, nil
			})

		gw := NewGateway("api.example.com", deps.http)

		expectedResult := map[int64]indto.GetSubscribersByDoctorIDsResponse{
			1: {
				DoctorID:          1,
				TgSubsCount:       "123",
				TgSubsCountText:   "подписчика",
				InstSubsCount:     "456",
				InstSubsCountText: "подписчиков",
			},
			2: {
				DoctorID:          2,
				TgSubsCount:       "789",
				TgSubsCountText:   "подписчиков",
				InstSubsCount:     "101",
				InstSubsCountText: "подписчик",
			},
		}

		result, err := gw.GetSubscribersByDoctorIDs(context.Background(), []int64{1, 2})

		require.NoError(t, err, "Не должно быть ошибки")
		require.NotNil(t, result, "Результат не должен быть nil")
		require.Len(t, result, 2, "Должны вернуться данные по 2 докторам")

		for id, expected := range expectedResult {
			assert.Equal(t, expected, result[id], "Данные доктора %d не соответствуют ожидаемым", id)
		}
	})

	t.Run("Пустой список ID докторов", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		gw := NewGateway("", deps.http)

		_, err := gw.GetSubscribersByDoctorIDs(context.Background(), []int64{})

		require.Error(t, err, "Должна быть ошибка при пустом списке ID")
		assert.Contains(t, err.Error(), "medblogersIDs is required", "Текст ошибки должен указывать на проблему")
	})

	t.Run("Ошибка HTTP запроса", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		expectedErr := errors.New("connection error")
		deps.http.EXPECT().
			Do(gomock.Any()).
			Return(nil, expectedErr)

		gw := NewGateway("", deps.http)

		_, err := gw.GetSubscribersByDoctorIDs(context.Background(), []int64{1})

		require.Error(t, err, "Должна быть ошибка")
		assert.ErrorIs(t, err, expectedErr, "Должна вернуться исходная ошибка")
	})

	t.Run("Неверный формат ответа", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().
			Do(gomock.Any()).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`invalid json`)),
			}, nil)

		gw := NewGateway("", deps.http)

		_, err := gw.GetSubscribersByDoctorIDs(context.Background(), []int64{1})

		require.Error(t, err, "Должна быть ошибка парсинга")
		assert.Contains(t, err.Error(), "invalid character 'i' looking for beginning of value", "Текст ошибки должен указывать на проблему")
	})

	t.Run("Частичный ответ (не все запрошенные доктора)", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
            "data": {
                "1": {
                    "doctor_id": 1,
                    "telegram_subs_count": "123",
                    "telegram_subs_text": "подписчика",
                    "instagram_subs_count": "456",
                    "instagram_subs_text": "подписчиков"
                }
			}
        }`

		deps.http.EXPECT().
			Do(gomock.Any()).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
			}, nil)

		gw := NewGateway("", deps.http)

		result, err := gw.GetSubscribersByDoctorIDs(context.Background(), []int64{1, 2})

		require.NoError(t, err, "Не должно быть ошибки")
		require.Len(t, result, 1, "Должен вернуться только 1 доктор")
		assert.Contains(t, result, int64(1), "Должен быть доктор с ID 1")
		assert.NotContains(t, result, int64(2), "Не должно быть доктора с ID 2")
	})
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

	t.Run("Ошибка при получении данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusInternalServerError}, errors.New("internal server error"))
		gw := NewGateway("", deps.http)

		expectedResult := indto.GetAllSubscribersInfoResponse{}
		result, err := gw.GetAllSubscribersInfo(context.Background())

		require.Error(t, err, "Должна быть ошибка")
		require.NotNil(t, result, "Результат не должен быть nil")

		assert.Equal(t, expectedResult, result, "Результат не соответствует ожидаемому")
	})
}

func TestGetFilterInfo(t *testing.T) {
	t.Parallel()

	t.Run("Успешное получение информации по фильтрам", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		mockResponseBody := `{
			"messengers": [
				{"name": "Телеграм", "slug": "tg"},
				{"name": "Инстаграм", "slug": "inst"}
			]
        }`
		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
		}, nil)

		expectedResult := []indto.FilterInfoResponse{
			{Name: "Телеграм", Slug: "tg"},
			{Name: "Инстаграм", Slug: "inst"},
		}
		gw := NewGateway("", deps.http)
		res, err := gw.GetFilterInfo(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedResult, res)
	})

	t.Run("Ошибка при получении данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusInternalServerError}, errors.New("internal server error"))
		gw := NewGateway("", deps.http)

		expectedResult := []indto.FilterInfoResponse{}
		result, err := gw.GetFilterInfo(context.Background())

		require.Error(t, err, "Должна быть ошибка")
		require.NotNil(t, result, "Результат не должен быть nil")

		assert.Equal(t, expectedResult, result, "Результат не соответствует ожидаемому")
	})
}

func TestCreateDoctor(t *testing.T) {
	t.Parallel()

	t.Run("Успешное создание доктора", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		// Ожидаемое тело запроса
		expectedRequest := dto.UpdateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		}

		// Мок ответа
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"status": "success"}`)),
			Header:     make(http.Header),
		}
		mockResponse.Header.Set("Content-Type", "application/json")

		// Проверяем корректность запроса
		deps.http.EXPECT().
			Do(gomock.Any()).
			DoAndReturn(func(req *http.Request) (*http.Response, error) {
				// Проверяем метод и URL
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "http:///doctors/create/", req.URL.String())

				// Проверяем тело запроса
				var actualBody dto.UpdateDoctorRequest
				err := json.NewDecoder(req.Body).Decode(&actualBody)
				assert.NoError(t, err)
				assert.Equal(t, expectedRequest, actualBody)

				// Проверяем заголовки
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

				return mockResponse, nil
			})

		gw := NewGateway("", deps.http)
		statusCode, err := gw.CreateDoctor(context.Background(), doctor.MedblogersID(1), indto.CreateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		})

		require.NoError(t, err)
		require.Equal(t, int64(http.StatusOK), statusCode)
	})

	t.Run("Ошибка при получении данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		expectedErr := errors.New("internal server error")
		deps.http.EXPECT().
			Do(gomock.Any()).
			Return(&http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewBufferString(`{"error": "server error"}`)),
			}, expectedErr)

		gw := NewGateway("", deps.http)

		statusCode, err := gw.CreateDoctor(context.Background(), doctor.MedblogersID(1), indto.CreateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		})

		require.Error(t, err)
		require.Equal(t, int64(http.StatusInternalServerError), statusCode)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("Нулевой medblogersID", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		gw := NewGateway("", deps.http)

		statusCode, err := gw.CreateDoctor(context.Background(), doctor.MedblogersID(0), indto.CreateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "medblogersID is required")
		require.Equal(t, int64(0), statusCode)
	})

	t.Run("Неверный формат ответа", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().
			Do(gomock.Any()).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`invalid json`)),
			}, nil)

		gw := NewGateway("", deps.http)

		statusCode, err := gw.CreateDoctor(context.Background(), doctor.MedblogersID(1), indto.CreateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		})

		require.NoError(t, err) // Ошибка парсинга не возвращается, только статус
		require.Equal(t, int64(http.StatusOK), statusCode)
	})
}

func TestUpdateDoctor(t *testing.T) {
	t.Parallel()

	t.Run("Успешное обновление доктора", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{
			StatusCode: http.StatusOK,
		}, nil)
		gw := NewGateway("", deps.http)
		statusCode, err := gw.UpdateDoctor(context.Background(), doctor.MedblogersID(1), indto.UpdateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		})

		require.NoError(t, err)
		require.Equal(t, int64(http.StatusOK), statusCode)
	})

	t.Run("Ошибка при получении данных", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.http.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusInternalServerError}, errors.New("internal server error"))
		gw := NewGateway("", deps.http)

		statusCode, err := gw.UpdateDoctor(context.Background(), doctor.MedblogersID(1), indto.UpdateDoctorRequest{
			Telegram:  "telegram",
			Instagram: "instagram",
		})
		require.Error(t, err)
		require.Equal(t, int64(http.StatusInternalServerError), statusCode)
	})
}

// TestConfigureFilterRequest табличный тест конфигурации запроса по фильтрации врачей
func TestConfigureFilterRequest(t *testing.T) {
	t.Parallel()
	deps := p(t)

	gw := NewGateway("api.example.com", deps.http)

	tests := []struct {
		name     string
		request  indto.GetDoctorsByFilterRequest
		expected string
	}{
		{
			name:     "Пустой запрос",
			request:  indto.GetDoctorsByFilterRequest{},
			expected: "http://api.example.com/doctors/filter/",
		},
		{
			name: "Только соцсети",
			request: indto.GetDoctorsByFilterRequest{
				SocialMedia: []indto.SocialMedia{indto.Telegram, indto.Instagram},
			},
			expected: `http://api.example.com/doctors/filter/?social_media=["tg","inst"]`,
		},
		{
			name: "Все параметры",
			request: indto.GetDoctorsByFilterRequest{
				SocialMedia:    []indto.SocialMedia{indto.Telegram},
				MinSubscribers: 100,
				MaxSubscribers: 1000,
				Offset:         10,
				Limit:          20,
			},
			expected: `http://api.example.com/doctors/filter/?limit=20&max_subscribers=1000&min_subscribers=100&offset=10&social_media=["tg"]`,
		},
		{
			name: "Только лимит и офсет",
			request: indto.GetDoctorsByFilterRequest{
				Limit:  30,
				Offset: 5,
			},
			expected: "http://api.example.com/doctors/filter/?limit=30&offset=5",
		},
		{
			name: "Только подписчики",
			request: indto.GetDoctorsByFilterRequest{
				MinSubscribers: 50,
				MaxSubscribers: 500,
			},
			expected: "http://api.example.com/doctors/filter/?max_subscribers=500&min_subscribers=50",
		},
		{
			name: "Одна соцсеть",
			request: indto.GetDoctorsByFilterRequest{
				SocialMedia: []indto.SocialMedia{indto.Instagram},
			},
			expected: `http://api.example.com/doctors/filter/?social_media=["inst"]`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := gw.configureFilterRequest(tt.request)
			r, err := url.QueryUnescape(result)
			assert.Equal(t, tt.expected, r, "Сформированный URL не соответствует ожидаемому")
			assert.NoError(t, err, "Сформированный URL должен быть валидным")
		})
	}
}
