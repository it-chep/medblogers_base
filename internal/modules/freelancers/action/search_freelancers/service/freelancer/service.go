package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/dto"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/logger"

	"github.com/samber/lo"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SearchStorage,ImageGetter

// SearchStorage .
type SearchStorage interface {
	SearchFreelancers(ctx context.Context, query string) ([]*freelancer.Freelancer, error)
}

// ImageGetter .
type ImageGetter interface {
	GetUserPhotos(ctx context.Context) (map[string]string, error)
}

// Service .
type Service struct {
	searchStorage SearchStorage
	imageGetter   ImageGetter
}

// NewSearchService .
func NewSearchService(searchStorage SearchStorage, imageGetter ImageGetter) *Service {
	return &Service{
		searchStorage: searchStorage,
		imageGetter:   imageGetter,
	}
}

func (s *Service) Search(ctx context.Context, query string) ([]dto.FreelancerItem, error) {
	logger.Message(ctx, "[Search] Поиск фрилансеров ...")

	freelancers, err := s.searchStorage.SearchFreelancers(ctx, query)
	if err != nil {
		return nil, err
	}

	logger.Message(ctx, "[Search] Обогащение докторов фотографиями")
	usersPhotosMap, err := s.imageGetter.GetUserPhotos(ctx)
	if err != nil {
		return nil, err
	}

	freelancersDTO := lo.Map(freelancers, func(item *freelancer.Freelancer, _ int) dto.FreelancerItem {
		return dto.FreelancerItem{
			ID:                    item.GetID(),
			Name:                  item.GetName(),
			Slug:                  item.GetSlug(),
			CityName:              item.GetCityName(),
			SpecialityName:        item.GetSpecialityName(),
			S3Image:               usersPhotosMap[item.GetSlug()],
			ExperienceWithDoctors: item.HasExperienceWithDoctor(),
			PriceCategory:         item.GetPriceCategory(),
			HasCommand:            item.GetHasCommand(),
		}
	})

	return freelancersDTO, nil
}
