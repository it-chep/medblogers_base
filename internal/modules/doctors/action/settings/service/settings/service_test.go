package settings

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/settings/dto"
	"medblogers_base/internal/modules/doctors/action/settings/service/settings/mocks"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	cityStorageMock       *mocks.MockCityStorage
	specialityStorageMock *mocks.MockSpecialityStorage
	subscribersGetterMock *mocks.MockSubscribersGetter
}

func p(t *testing.T) fields {
	return fields{
		cityStorageMock:       mocks.NewMockCityStorage(gomock.NewController(t)),
		specialityStorageMock: mocks.NewMockSpecialityStorage(gomock.NewController(t)),
		subscribersGetterMock: mocks.NewMockSubscribersGetter(gomock.NewController(t)),
	}
}

func TestService_GetSettings(t *testing.T) {
	t.Parallel()

	t.Run("Успешное получение всех настроек", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.specialityStorageMock.EXPECT().GetSpecialitiesWithDoctorsCount(gomock.Any()).Return([]*speciality.Speciality{
			speciality.BuildSpeciality(speciality.WithID(1), speciality.WithName("хирург"), speciality.WithDoctorsCount(123)),
			speciality.BuildSpeciality(speciality.WithID(2), speciality.WithName("гинеколог"), speciality.WithDoctorsCount(123)),
		}, nil)
		deps.cityStorageMock.EXPECT().GetCitiesWithDoctorsCount(gomock.Any()).Return([]*city.City{
			city.BuildCity(city.WithID(1), city.WithName("мск"), city.WithDoctorsCount(123)),
			city.BuildCity(city.WithID(2), city.WithName("спб"), city.WithDoctorsCount(123)),
		}, nil)
		deps.subscribersGetterMock.EXPECT().GetFilterInfo(gomock.Any()).Return([]indto.FilterInfoResponse{{Name: "tg", Slug: "tg"}}, nil)

		expectedResult := &dto.Settings{
			FilterInfo:      []dto.FilterItem{{Name: "tg", Slug: "tg"}},
			Cities:          []dto.CityItem{{ID: 1, Name: "мск", DoctorsCount: 123}, {ID: 2, Name: "спб", DoctorsCount: 123}},
			Specialities:    []dto.SpecialityItem{{ID: 1, Name: "хирург", DoctorsCount: 123}, {ID: 2, Name: "гинеколог", DoctorsCount: 123}},
			NewDoctorBanner: true,
		}

		service := NewSettingsService(deps.cityStorageMock, deps.specialityStorageMock, deps.subscribersGetterMock)
		settings, err := service.GetSettings(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, settings)
	})

	t.Run("Ошибка получения фильтра подписчиков", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.specialityStorageMock.EXPECT().GetSpecialitiesWithDoctorsCount(gomock.Any()).Return([]*speciality.Speciality{
			speciality.BuildSpeciality(speciality.WithID(1), speciality.WithName("хирург"), speciality.WithDoctorsCount(123)),
			speciality.BuildSpeciality(speciality.WithID(2), speciality.WithName("гинеколог"), speciality.WithDoctorsCount(123)),
		}, nil)
		deps.cityStorageMock.EXPECT().GetCitiesWithDoctorsCount(gomock.Any()).Return([]*city.City{
			city.BuildCity(city.WithID(1), city.WithName("мск"), city.WithDoctorsCount(123)),
			city.BuildCity(city.WithID(2), city.WithName("спб"), city.WithDoctorsCount(123)),
		}, nil)
		deps.subscribersGetterMock.EXPECT().GetFilterInfo(gomock.Any()).Return([]indto.FilterInfoResponse{}, errors.New("internal error"))

		expectedResult := &dto.Settings{
			FilterInfo:      []dto.FilterItem{},
			Cities:          []dto.CityItem{{ID: 1, Name: "мск", DoctorsCount: 123}, {ID: 2, Name: "спб", DoctorsCount: 123}},
			Specialities:    []dto.SpecialityItem{{ID: 1, Name: "хирург", DoctorsCount: 123}, {ID: 2, Name: "гинеколог", DoctorsCount: 123}},
			NewDoctorBanner: true,
		}

		service := NewSettingsService(deps.cityStorageMock, deps.specialityStorageMock, deps.subscribersGetterMock)
		settings, err := service.GetSettings(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, settings)
	})

	t.Run("Ошибка получения фильтра городов", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.specialityStorageMock.EXPECT().GetSpecialitiesWithDoctorsCount(gomock.Any()).Return([]*speciality.Speciality{
			speciality.BuildSpeciality(speciality.WithID(1), speciality.WithName("хирург"), speciality.WithDoctorsCount(123)),
			speciality.BuildSpeciality(speciality.WithID(2), speciality.WithName("гинеколог"), speciality.WithDoctorsCount(123)),
		}, nil)
		deps.cityStorageMock.EXPECT().GetCitiesWithDoctorsCount(gomock.Any()).Return([]*city.City{}, errors.New("internal error"))
		deps.subscribersGetterMock.EXPECT().GetFilterInfo(gomock.Any()).Return([]indto.FilterInfoResponse{{Name: "tg", Slug: "tg"}}, nil)

		expectedResult := &dto.Settings{
			Cities:          []dto.CityItem{},
			FilterInfo:      []dto.FilterItem{{Name: "tg", Slug: "tg"}},
			Specialities:    []dto.SpecialityItem{{ID: 1, Name: "хирург", DoctorsCount: 123}, {ID: 2, Name: "гинеколог", DoctorsCount: 123}},
			NewDoctorBanner: true,
		}

		service := NewSettingsService(deps.cityStorageMock, deps.specialityStorageMock, deps.subscribersGetterMock)
		settings, err := service.GetSettings(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, settings)
	})

	t.Run("Ошибка получения фильтра специальностей", func(t *testing.T) {
		t.Parallel()
		deps := p(t)

		deps.specialityStorageMock.EXPECT().GetSpecialitiesWithDoctorsCount(gomock.Any()).Return([]*speciality.Speciality{}, errors.New("internal error"))
		deps.cityStorageMock.EXPECT().GetCitiesWithDoctorsCount(gomock.Any()).Return([]*city.City{
			city.BuildCity(city.WithID(1), city.WithName("мск"), city.WithDoctorsCount(123)),
			city.BuildCity(city.WithID(2), city.WithName("спб"), city.WithDoctorsCount(123)),
		}, nil)
		deps.subscribersGetterMock.EXPECT().GetFilterInfo(gomock.Any()).Return([]indto.FilterInfoResponse{{Name: "tg", Slug: "tg"}}, nil)

		expectedResult := &dto.Settings{
			Specialities:    []dto.SpecialityItem{},
			FilterInfo:      []dto.FilterItem{{Name: "tg", Slug: "tg"}},
			Cities:          []dto.CityItem{{ID: 1, Name: "мск", DoctorsCount: 123}, {ID: 2, Name: "спб", DoctorsCount: 123}},
			NewDoctorBanner: true,
		}

		service := NewSettingsService(deps.cityStorageMock, deps.specialityStorageMock, deps.subscribersGetterMock)
		settings, err := service.GetSettings(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, settings)
	})
}
