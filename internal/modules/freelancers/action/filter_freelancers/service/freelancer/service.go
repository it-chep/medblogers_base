package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
)

type Storage interface {
	FilterFreelancers(ctx context.Context, filter freelancer.Filter) (map[int64]*freelancer.Freelancer, []int64, error)
}

type ImageEnricher interface {
	GetUserPhotos(ctx context.Context) (map[string]string, error)
}

type AdditionalStorage interface {
	GetDoctorAdditionalCities(ctx context.Context, medblogersIDs []int64) (map[int64][]*city.City, error)
	GetDoctorAdditionalSpecialities(ctx context.Context, medblogersIDs []int64) (map[int64][]*speciality.Speciality, error)
}
