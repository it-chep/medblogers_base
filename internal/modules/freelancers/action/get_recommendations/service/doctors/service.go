package doctors

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/dto"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/doctor"
	"medblogers_base/internal/modules/freelancers/domain/speciality"

	"github.com/samber/lo"
)

type DoctorsDal interface {
	GetDoctorsInfo(ctx context.Context, doctorIDs []int64) ([]*doctor.Doctor, error)
	GetDoctorAdditionalCities(ctx context.Context, medblogersIDs []int64) (map[int64][]*city.City, error)
	GetDoctorAdditionalSpecialities(ctx context.Context, medblogersIDs []int64) (map[int64][]*speciality.Speciality, error)
}

type ImageEnricher interface {
	GetDoctorsPhotos(ctx context.Context) (map[doctor.S3Key]string, error)
}

type Service struct {
	dal         DoctorsDal
	imageGetter ImageEnricher
}

func NewService(dal DoctorsDal, imageEnricher ImageEnricher) *Service {
	return &Service{
		dal:         dal,
		imageGetter: imageEnricher,
	}
}

// GetDoctorsInfo получение информации о докторах
func (s *Service) GetDoctorsInfo(ctx context.Context, doctorIDs []int64) (map[int64]dto.Doctor, []int64, error) {
	doctors, err := s.dal.GetDoctorsInfo(ctx, doctorIDs)

	if err != nil {
		return nil, nil, err
	}
	// конвертация в DTO
	dtoMap := s.convertToDTOMap(lo.SliceToMap(doctors, func(item *doctor.Doctor) (doctor.MedblogersID, *doctor.Doctor) {
		return item.GetID(), item
	}))

	orderedIDs := lo.Map(doctors, func(item *doctor.Doctor, _ int) int64 {
		return int64(item.GetID())
	})

	return dtoMap, orderedIDs, nil
}

func (s *Service) convertToDTOMap(doctorsMap map[doctor.MedblogersID]*doctor.Doctor) map[int64]dto.Doctor {
	dtoMap := make(map[int64]dto.Doctor, len(doctorsMap))
	for _, doc := range doctorsMap {
		dtoMap[int64(doc.GetID())] = dto.Doctor{
			ID:               int64(doc.GetID()),
			Slug:             doc.GetSlug(),
			Name:             doc.GetName(),
			MainCityID:       doc.GetMainCityID(),
			MainSpecialityID: doc.GetMainSpecialityID(),
			S3Key:            doc.GetS3Key().String(),
		}
	}

	return dtoMap
}
