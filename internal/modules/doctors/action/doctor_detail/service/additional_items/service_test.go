package additional_items

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/additional_items/mocks"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
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
			[]*city.City{
				city.BuildCity(city.WithName("мск"), city.WithID(1)),
				city.BuildCity(city.WithName("спб"), city.WithID(2)),
			},
			nil,
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

func TestService_GetAdditionalSpecialities(t *testing.T) {
	t.Parallel()

	doctorID := doctor.MedblogersID(1234)
	t.Run("успешное получение специальностей", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		deps.storageMock.EXPECT().GetDoctorAdditionalSpecialities(gomock.Any(), doctorID).Return(
			[]*speciality.Speciality{
				speciality.BuildSpeciality(speciality.WithName("терапевт"), speciality.WithID(1)),
				speciality.BuildSpeciality(speciality.WithName("кардиолог"), speciality.WithID(2)),
			},
			nil,
		)
		service := New(deps.storageMock)

		expectedMainSpeciality := dto.SpecialityItem{
			Name: "терапевт",
			ID:   1,
		}
		expectedAdditionalSpecialities := []dto.SpecialityItem{
			{
				Name: "кардиолог",
				ID:   2,
			},
		}

		mainSpeciality, additionalSpecialities, err := service.GetAdditionalSpecialities(context.Background(), doctorID, 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedMainSpeciality, mainSpeciality)
		assert.Len(t, expectedAdditionalSpecialities, len(additionalSpecialities))
		assert.Equal(t, expectedAdditionalSpecialities, additionalSpecialities)
	})

	t.Run("ошибка получения специальностей", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		deps.storageMock.EXPECT().GetDoctorAdditionalSpecialities(gomock.Any(), doctorID).Return(
			nil, errors.New("ошибка базки"),
		)
		service := New(deps.storageMock)

		expectedMainSpeciality := dto.SpecialityItem{}
		expectedAdditionalSpecialities := []dto.SpecialityItem{}

		mainSpeciality, additionalSpecialities, err := service.GetAdditionalSpecialities(context.Background(), doctorID, 1)

		assert.Error(t, err)
		assert.Equal(t, expectedMainSpeciality, mainSpeciality)
		assert.Len(t, expectedAdditionalSpecialities, len(additionalSpecialities))
		assert.Equal(t, expectedAdditionalSpecialities, additionalSpecialities)
	})
}
