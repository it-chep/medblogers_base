package doctor

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

type Storage interface {
	GetMBCAggregation(ctx context.Context) ([]dto.MBC, error)
	GetDoctorsByIDs(ctx context.Context, ids []int64) ([]dto.RatingItem, error)
	GetCitiesByIDs(ctx context.Context, ids []int64) ([]*city.City, error)
	GetSpecialitiesByIDs(ctx context.Context, ids []int64) ([]*speciality.Speciality, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetRating(ctx context.Context) ([]dto.RatingItem, error) {
	mbcItems, err := s.storage.GetMBCAggregation(ctx)
	if err != nil {
		return nil, err
	}

	if len(mbcItems) == 0 {
		return []dto.RatingItem{}, nil
	}

	doctorIDs := lo.Map(mbcItems, func(item dto.MBC, _ int) int64 {
		return item.DoctorID
	})

	doctors, err := s.storage.GetDoctorsByIDs(ctx, doctorIDs)
	if err != nil {
		return nil, err
	}

	cityIDs, specIDs := extractLookupIDs(doctors)

	cityMap, err := s.loadCityMap(ctx, cityIDs)
	if err != nil {
		return nil, err
	}

	specMap, err := s.loadSpecialityMap(ctx, specIDs)
	if err != nil {
		return nil, err
	}

	return buildRatingItems(mbcItems, doctors, cityMap, specMap), nil
}

func extractLookupIDs(doctors []dto.RatingItem) ([]int64, []int64) {
	cityIDSet := make(map[int64]struct{})
	specIDSet := make(map[int64]struct{})
	for _, d := range doctors {
		cityIDSet[d.CityID] = struct{}{}
		specIDSet[d.SpecialityID] = struct{}{}
	}

	cityIDs := make([]int64, 0, len(cityIDSet))
	for id := range cityIDSet {
		cityIDs = append(cityIDs, id)
	}

	specIDs := make([]int64, 0, len(specIDSet))
	for id := range specIDSet {
		specIDs = append(specIDs, id)
	}

	return cityIDs, specIDs
}

func (s *Service) loadCityMap(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	cities, err := s.storage.GetCitiesByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]string, len(cities))
	for _, c := range cities {
		result[int64(c.ID())] = c.Name()
	}
	return result, nil
}

func (s *Service) loadSpecialityMap(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	specs, err := s.storage.GetSpecialitiesByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]string, len(specs))
	for _, s := range specs {
		result[int64(s.ID())] = s.Name()
	}
	return result, nil
}

func buildRatingItems(
	mbcItems []dto.MBC,
	doctors []dto.RatingItem,
	cityMap map[int64]string,
	specMap map[int64]string,
) []dto.RatingItem {
	doctorMap := make(map[int64]dto.RatingItem, len(doctors))
	for _, d := range doctors {
		doctorMap[d.DoctorID] = d
	}

	result := make([]dto.RatingItem, 0, len(doctors))
	for _, mbc := range mbcItems {
		d, ok := doctorMap[mbc.DoctorID]
		if !ok {
			continue
		}

		item := dto.RatingItem{
			Slug:           d.Slug,
			Name:           d.Name,
			S3Image:        d.S3Image,
			MBCCoins:       mbc.MBCCount,
			CityID:         d.CityID,
			CityName:       cityMap[d.CityID],
			SpecialityID:   d.SpecialityID,
			SpecialityName: specMap[d.SpecialityID],
		}

		result = append(result, item)
	}

	return result
}
