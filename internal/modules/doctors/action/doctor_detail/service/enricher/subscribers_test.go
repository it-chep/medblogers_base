package enricher

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/enricher/mocks"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func fs(t *testing.T) fieldsSubs {
	return fieldsSubs{
		subscribersMock: mocks.NewMockSubscribersGetter(gomock.NewController(t)),
	}
}

type fieldsSubs struct {
	subscribersMock *mocks.MockSubscribersGetter
}

func TestSubscribersEnricher_Enrich(t *testing.T) {
	t.Parallel()

	doctorID := doctor.MedblogersID(123)
	t.Run("Успешное обогащение подписчиков", func(t *testing.T) {
		t.Parallel()
		deps := fs(t)
		deps.subscribersMock.EXPECT().GetDoctorSubscribers(gomock.Any(), doctorID).Return(
			indto.GetDoctorSubscribersResponse{
				TgSubsCount:         "123",
				TgSubsCountText:     "подписчика",
				TgLastUpdatedDate:   "11.05.2004",
				InstSubsCount:       "123",
				InstSubsCountText:   "подписчика",
				InstLastUpdatedDate: "11.05.2004",
			}, nil,
		)
		service := NewSubscribersEnricher(deps.subscribersMock)

		expectedRes := &dto.DoctorDTO{
			TgSubsCount:         "123",
			TgSubsCountText:     "подписчика",
			TgLastUpdatedDate:   "11.05.2004",
			InstSubsCount:       "123",
			InstSubsCountText:   "подписчика",
			InstLastUpdatedDate: "11.05.2004",
		}
		result, err := service.Enrich(context.Background(), doctorID, &dto.DoctorDTO{})

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})

	t.Run("Ошибка от сервиса подписчиков", func(t *testing.T) {
		t.Parallel()
		deps := fs(t)
		deps.subscribersMock.EXPECT().GetDoctorSubscribers(gomock.Any(), doctorID).Return(
			indto.GetDoctorSubscribersResponse{}, errors.New("ошибка от подписчиков"),
		)
		service := NewSubscribersEnricher(deps.subscribersMock)

		expectedRes := &dto.DoctorDTO{}
		result, err := service.Enrich(context.Background(), doctorID, &dto.DoctorDTO{})

		assert.Error(t, err)
		assert.Equal(t, expectedRes, result)
	})

	t.Run("Пришло число 0 (не блокирующая операция)", func(t *testing.T) {
		t.Parallel()

		deps := fs(t)
		deps.subscribersMock.EXPECT().GetDoctorSubscribers(gomock.Any(), doctorID).Return(
			indto.GetDoctorSubscribersResponse{
				TgSubsCount:         "0",
				TgSubsCountText:     "подписчиков",
				TgLastUpdatedDate:   "11.05.2004",
				InstSubsCount:       "0",
				InstSubsCountText:   "подписчиков",
				InstLastUpdatedDate: "11.05.2004",
			}, nil,
		)
		service := NewSubscribersEnricher(deps.subscribersMock)

		expectedRes := &dto.DoctorDTO{
			TgSubsCount:         "0",
			TgSubsCountText:     "подписчиков",
			TgLastUpdatedDate:   "11.05.2004",
			InstSubsCount:       "0",
			InstSubsCountText:   "подписчиков",
			InstLastUpdatedDate: "11.05.2004",
		}
		result, err := service.Enrich(context.Background(), doctorID, &dto.DoctorDTO{})

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})
}
