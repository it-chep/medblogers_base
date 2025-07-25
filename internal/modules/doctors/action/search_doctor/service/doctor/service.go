package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/search_doctor/dto"
	"medblogers_base/internal/pkg/logger"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/doctors/domain/doctor"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SearchStorage,ImageGetter

// SearchStorage .
type SearchStorage interface {
	SearchDoctors(ctx context.Context, query string) ([]*doctor.Doctor, error)
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

func (s *Service) Search(ctx context.Context, query string) ([]dto.DoctorItem, error) {
	logger.Message(ctx, "[Search] Поиск докторов ...")

	doctors, err := s.searchStorage.SearchDoctors(ctx, query)
	if err != nil {
		return nil, err
	}

	logger.Message(ctx, "[Search] Обогащение докторов фотографиями")
	usersPhotosMap, err := s.imageGetter.GetUserPhotos(ctx)
	if err != nil {
		return nil, err
	}

	doctorsDTO := lo.Map(doctors, func(item *doctor.Doctor, _ int) dto.DoctorItem {
		return dto.DoctorItem{
			ID:             int64(item.GetID()),
			Name:           item.GetName(),
			Slug:           item.GetSlug(),
			CityName:       item.GetMainCityName(),
			SpecialityName: item.GetMainSpecialityName(),
			S3Image:        usersPhotosMap[item.GetSlug()],
		}
	})

	return doctorsDTO, nil
}
