package additional_items

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/additional_items/mocks"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	storageMock *mocks.MockStorage
}

func p(t *testing.T) fields {
	return fields{
		storageMock: mocks.NewMockStorage(gomock.NewController(t)),
	}
}

func TestService_GetAdditionalCities(t *testing.T) {
	t.Parallel()

	doctorID := doctor.MedblogersID(1234)
	t.Run("успешное получение городов", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		deps.storageMock.EXPECT().GetDoctorAdditionalCities(gomock.Any(), doctorID).Return(
			map[city.CityID]*city.City{
				city.CityID(1): city.BuildCity(city.WithName("мск"), city.WithID(1)),
				city.CityID(2): city.BuildCity(city.WithName("спб"), city.WithID(2)),
			}, nil,
		)
		service := New(deps.storageMock)

		expectedMainCity := dto.CityItem{
			Name: "мск",
			ID:   1,
		}
		expectedAdditionalCities := []dto.CityItem{
			{
				Name: "спб",
				ID:   2,
			},
		}

		mainCity, additionalCities, err := service.GetAdditionalCities(context.Background(), doctorID, 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedMainCity, mainCity)
		assert.Len(t, expectedAdditionalCities, len(additionalCities))
		assert.Equal(t, expectedAdditionalCities, additionalCities)
	})

	t.Run("ошибка получения городов", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		deps.storageMock.EXPECT().GetDoctorAdditionalCities(gomock.Any(), doctorID).Return(
			nil, errors.New("ошибка базки"),
		)
		service := New(deps.storageMock)

		expectedMainCity := dto.CityItem{}
		expectedAdditionalCities := []dto.CityItem{}

		mainCity, additionalCities, err := service.GetAdditionalCities(context.Background(), doctorID, 1)

		assert.Error(t, err)
		assert.Equal(t, expectedMainCity, mainCity)
		assert.Len(t, expectedAdditionalCities, len(additionalCities))
		assert.Equal(t, expectedAdditionalCities, additionalCities)
	})
}
