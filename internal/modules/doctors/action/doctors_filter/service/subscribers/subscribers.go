package subscribers

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"

	"github.com/samber/lo"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SubscribersGetter

// SubscribersGetter получение информации о подписчиках
type SubscribersGetter interface {
	GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) (indto.GetDoctorsByFilterResponse, error)
	GetDoctorsByFilterWithIDs(ctx context.Context, request indto.GetDoctorsByFilterRequest, doctorsIDs []int64) (indto.GetDoctorsByFilterResponse, error)
}

type Service struct {
	getter SubscribersGetter
}

func New(getter SubscribersGetter) *Service {
	return &Service{
		getter: getter,
	}
}

func (s *Service) FilterDoctorsBySubscribersWithDoctorsIDs(ctx context.Context, filter dto.Filter, doctorsIDs []int64) (dto.SubscribersInfo, error) {
	response, err := s.getter.GetDoctorsByFilterWithIDs(
		ctx, indto.GetDoctorsByFilterRequest{
			MinSubscribers: filter.MinSubscribers,
			MaxSubscribers: filter.MaxSubscribers,
			SocialMedia: lo.Map(filter.SocialMedia, func(socialMedia string, index int) indto.SocialMedia {
				return indto.NewSocialMedia(socialMedia)
			}),
			Offset: filter.Page,
			Sort:   filter.Sort.String(),
		},
		doctorsIDs,
	)
	if err != nil {
		return dto.SubscribersInfo{}, err
	}

	result := make(map[int64]dto.DoctorSubscribersInfoDTO, len(response.Doctors))
	for _, doctor := range response.Doctors {
		result[doctor.DoctorID] = dto.DoctorSubscribersInfoDTO{
			InstSubsCount:        doctor.InstSubsCount,
			InstSubsCountText:    doctor.InstSubsCountText,
			TgSubsCount:          doctor.TgSubsCount,
			TgSubsCountText:      doctor.TgSubsCountText,
			YouTubeSubsCount:     doctor.YouTubeSubsCount,
			YouTubeSubsCountText: doctor.YouTubeSubsCountText,
		}
	}

	return dto.SubscribersInfo{
		Doctors:      result,
		DoctorsCount: response.DoctorsCount,
		SubsCount:    response.SubscribersCount,
		OrderedIDs:   response.OrderedIDs,
	}, nil
}

// MapDoctorsWithSubscribers маппим параметры докторов с сортировкой и данными из подписчиков
func (s *Service) MapDoctorsWithSubscribers(doctorsMap map[int64]dto.Doctor, subsMap map[int64]dto.DoctorSubscribersInfoDTO, orderedIDs []int64) []dto.Doctor {
	mappedDoctors := make([]dto.Doctor, 0, len(subsMap))

	for _, doctorID := range orderedIDs {
		doctorData := doctorsMap[doctorID]
		subs := subsMap[doctorID]

		if doctorData.ID == 0 { // todo заменить на mo ?
			continue
		}

		mappedDoctors = append(mappedDoctors, dto.Doctor{
			InstLink:          doctorData.InstLink,
			InstSubsCount:     subs.InstSubsCount,
			InstSubsCountText: subs.InstSubsCountText,

			TgLink:          doctorData.TgLink,
			TgSubsCount:     subs.TgSubsCount,
			TgSubsCountText: subs.TgSubsCountText,

			YouTubeLink:          doctorData.YouTubeLink,
			YouTubeSubsCount:     subs.YouTubeSubsCount,
			YouTubeSubsCountText: subs.YouTubeSubsCountText,

			ID:    doctorData.ID,
			Name:  doctorData.Name,
			Slug:  doctorData.Slug,
			Image: doctorData.Image,

			MainCityID:       doctorData.MainCityID,
			MainSpecialityID: doctorData.MainSpecialityID,

			Specialities: doctorData.Specialities,
			Cities:       doctorData.Cities,

			S3Key:      doctorData.S3Key,
			IsKFDoctor: doctorData.IsKFDoctor,
		})
	}

	return mappedDoctors
}

// FilterDoctorsBySubscribers - фильтрация по количеству подписчиков
func (s *Service) FilterDoctorsBySubscribers(ctx context.Context, filter dto.Filter) (dto.SubscribersInfo, error) {
	response, err := s.getter.GetDoctorsByFilter(
		ctx, indto.GetDoctorsByFilterRequest{
			MinSubscribers: filter.MinSubscribers,
			MaxSubscribers: filter.MaxSubscribers,
			SocialMedia: lo.Map(filter.SocialMedia, func(socialMedia string, index int) indto.SocialMedia {
				return indto.NewSocialMedia(socialMedia)
			}),
			Offset: filter.Page,
			Sort:   filter.Sort.String(),
		},
	)
	if err != nil {
		return dto.SubscribersInfo{}, err
	}

	result := make(map[int64]dto.DoctorSubscribersInfoDTO, len(response.Doctors))
	for _, doctor := range response.Doctors {
		result[doctor.DoctorID] = dto.DoctorSubscribersInfoDTO{
			InstSubsCount:        doctor.InstSubsCount,
			InstSubsCountText:    doctor.InstSubsCountText,
			TgSubsCount:          doctor.TgSubsCount,
			TgSubsCountText:      doctor.TgSubsCountText,
			YouTubeSubsCount:     doctor.YouTubeSubsCount,
			YouTubeSubsCountText: doctor.YouTubeSubsCountText,
		}
	}

	return dto.SubscribersInfo{
		Doctors:      result,
		DoctorsCount: response.DoctorsCount,
		SubsCount:    response.SubscribersCount,
		OrderedIDs:   response.OrderedIDs,
	}, nil
}
