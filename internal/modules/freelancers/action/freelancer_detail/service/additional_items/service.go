package additional_items

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail/dto"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/price_list"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	GetAdditionalCities(ctx context.Context, freelancerID int64) (map[int64]*city.City, error)
	GetAdditionalSpecialities(ctx context.Context, freelancerID int64) (map[int64]*speciality.Speciality, error)
	GetPriceList(ctx context.Context, freelancerID int64) (price_list.PriceList, error)
	GetSocialNetworks(ctx context.Context, freelancerID int64) (social_network.Networks, error)
}

type Service struct {
	store Storage
}

func New(storage Storage) *Service {
	return &Service{
		store: storage,
	}
}

func (s *Service) GetAdditionalCities(ctx context.Context, freelancerID, mainCityID int64) (_ dto.CityItem, _ []dto.CityItem, err error) {
	citiesMap, err := s.store.GetAdditionalCities(ctx, freelancerID)
	if err != nil {
		return dto.CityItem{}, []dto.CityItem{}, err
	}

	mainCity := dto.CityItem{}
	cities := make([]dto.CityItem, 0, len(citiesMap))
	for _, c := range citiesMap {
		// Определяем основной город
		if c.ID() == mainCityID {
			mainCity = dto.CityItem{
				ID:   c.ID(),
				Name: c.Name(),
			}
			continue
		}

		// Сохраняем дополнительные
		cities = append(cities, dto.CityItem{
			ID:   c.ID(),
			Name: c.Name(),
		})
	}

	return mainCity, cities, nil
}

func (s *Service) GetAdditionalSpecialities(ctx context.Context, freelancerID, mainSpecialityID int64) (_ dto.SpecialityItem, _ []dto.SpecialityItem, err error) {
	specialitiesMap, err := s.store.GetAdditionalSpecialities(ctx, freelancerID)
	if err != nil {
		return dto.SpecialityItem{}, []dto.SpecialityItem{}, err
	}

	mainSpeciality := dto.SpecialityItem{}
	specialities := make([]dto.SpecialityItem, 0, len(specialitiesMap))
	for _, sp := range specialitiesMap {
		// Запоминаем основную специальность
		if sp.ID() == mainSpecialityID {
			mainSpeciality = dto.SpecialityItem{
				ID:   sp.ID(),
				Name: sp.Name(),
			}
			continue
		}
		// Сохраняем дополнительные специальности
		specialities = append(specialities, dto.SpecialityItem{
			ID:   sp.ID(),
			Name: sp.Name(),
		})
	}

	return mainSpeciality, specialities, nil
}

func (s *Service) GetPriceList(ctx context.Context, freelancerID int64) ([]dto.PriceListItem, error) {
	priceList, err := s.store.GetPriceList(ctx, freelancerID)
	if err != nil {
		return []dto.PriceListItem{}, err
	}

	dtoPriceList := make([]dto.PriceListItem, 0, len(priceList))
	for _, item := range priceList {
		dtoPriceList = append(dtoPriceList, dto.PriceListItem{
			Name:  item.GetName(),
			Price: item.GetPrice(),
		})
	}

	return dtoPriceList, nil
}

func (s *Service) GetSocialNetworks(ctx context.Context, freelancerID int64) ([]dto.SocialNetworkItem, error) {
	networks, err := s.store.GetSocialNetworks(ctx, freelancerID)
	if err != nil {
		return []dto.SocialNetworkItem{}, err
	}

	dtoSocialNetworkItems := make([]dto.SocialNetworkItem, 0, len(networks))
	for _, n := range networks {
		dtoSocialNetworkItems = append(dtoSocialNetworkItems, dto.SocialNetworkItem{
			ID:   n.ID(),
			Name: n.Name(),
		})
	}

	return dtoSocialNetworkItems, nil
}
