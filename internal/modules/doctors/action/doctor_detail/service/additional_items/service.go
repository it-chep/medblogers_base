package additional_items

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"

	"github.com/samber/lo"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	GetDoctorAdditionalCities(ctx context.Context, doctorID doctor.MedblogersID) (map[city.CityID]*city.City, error)
	GetDoctorAdditionalSpecialities(ctx context.Context, doctorID doctor.MedblogersID) (map[speciality.SpecialityID]*speciality.Speciality, error)
}

type Service struct {
	store Storage
}

func New(storage Storage) *Service {
	return &Service{
		store: storage,
	}
}

func (s *Service) GetAdditionalCities(ctx context.Context, doctorID doctor.MedblogersID, mainCityID city.CityID) (_ []dto.CityItem, err error) {
	citiesMap, err := s.store.GetDoctorAdditionalCities(ctx, doctorID)
	if err != nil {
		return []dto.CityItem{}, err
	}

	cities := lo.Map(
		// исключаем основной город
		lo.Filter(
			lo.Values(citiesMap),
			func(item *city.City, _ int) bool {
				return item.ID() != mainCityID
			},
		),
		func(item *city.City, _ int) dto.CityItem {
			return dto.CityItem{
				ID:   int64(item.ID()),
				Name: item.Name(),
			}
		},
	)

	return cities, nil
}

func (s *Service) GetAdditionalSpecialities(ctx context.Context, doctorID doctor.MedblogersID, mainSpecialityID speciality.SpecialityID) (_ []dto.SpecialityItem, err error) {
	specialitiesMap, err := s.store.GetDoctorAdditionalSpecialities(ctx, doctorID)
	if err != nil {
		return []dto.SpecialityItem{}, err
	}

	specialities := lo.Map(
		// исключаем основную специальность
		lo.Filter(
			lo.Values(specialitiesMap),
			func(item *speciality.Speciality, _ int) bool {
				return item.ID() != mainSpecialityID
			},
		),
		func(item *speciality.Speciality, _ int) dto.SpecialityItem {
			return dto.SpecialityItem{
				ID:   int64(item.ID()),
				Name: item.Name(),
			}
		},
	)

	return specialities, nil
}
