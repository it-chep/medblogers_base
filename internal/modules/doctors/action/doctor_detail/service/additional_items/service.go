package additional_items

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
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

func (s *Service) GetAdditionalCities(ctx context.Context, doctorID doctor.MedblogersID, mainCityID city.CityID) (_ dto.CityItem, _ []dto.CityItem, err error) {
	citiesMap, err := s.store.GetDoctorAdditionalCities(ctx, doctorID)
	if err != nil {
		return dto.CityItem{}, []dto.CityItem{}, err
	}

	mainCity := dto.CityItem{}
	cities := make([]dto.CityItem, 0, len(citiesMap))
	for _, c := range citiesMap {
		// Определяем основной город
		if c.ID() == mainCityID {
			mainCity = dto.CityItem{
				ID:   int64(c.ID()),
				Name: c.Name(),
			}
			continue
		}

		// Сохраняем дополнительные
		cities = append(cities, dto.CityItem{
			ID:   int64(c.ID()),
			Name: c.Name(),
		})
	}

	return mainCity, cities, nil
}

func (s *Service) GetAdditionalSpecialities(ctx context.Context, doctorID doctor.MedblogersID, mainSpecialityID speciality.SpecialityID) (_ dto.SpecialityItem, _ []dto.SpecialityItem, err error) {
	specialitiesMap, err := s.store.GetDoctorAdditionalSpecialities(ctx, doctorID)
	if err != nil {
		return dto.SpecialityItem{}, []dto.SpecialityItem{}, err
	}

	mainSpeciality := dto.SpecialityItem{}
	specialities := make([]dto.SpecialityItem, 0, len(specialitiesMap))
	for _, sp := range specialitiesMap {
		// Запоминаем основную специальность
		if sp.ID() == mainSpecialityID {
			mainSpeciality = dto.SpecialityItem{
				ID:   int64(sp.ID()),
				Name: sp.Name(),
			}
			continue
		}
		// Сохраняем дополнительные специальности
		specialities = append(specialities, dto.SpecialityItem{
			ID:   int64(sp.ID()),
			Name: sp.Name(),
		})
	}

	return mainSpeciality, specialities, nil
}
