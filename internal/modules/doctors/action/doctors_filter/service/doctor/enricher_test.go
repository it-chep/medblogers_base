package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnrichSubscribers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		doctorsMap     map[int64]dto.Doctor
		subscribersMap map[int64]indto.GetSubscribersByDoctorIDsResponse
		expected       map[int64]dto.Doctor
	}{
		{
			name:           "Пустые мапы",
			doctorsMap:     map[int64]dto.Doctor{},
			subscribersMap: map[int64]indto.GetSubscribersByDoctorIDsResponse{},
			expected:       map[int64]dto.Doctor{},
		},
		{
			name: "Доктор с подписчиками",
			doctorsMap: map[int64]dto.Doctor{
				1: {ID: 1, Name: "Dr. Smith"},
			},
			subscribersMap: map[int64]indto.GetSubscribersByDoctorIDsResponse{
				1: {
					TgSubsCount:       "100",
					TgSubsCountText:   "подписчиков",
					InstSubsCount:     "200",
					InstSubsCountText: "подписчиков",
				},
			},
			expected: map[int64]dto.Doctor{
				1: {
					ID:                1,
					Name:              "Dr. Smith",
					TgSubsCount:       "100",
					TgSubsCountText:   "подписчиков",
					InstSubsCount:     "200",
					InstSubsCountText: "подписчиков",
				},
			},
		},
		{
			name: "Доктор без подписчиков",
			doctorsMap: map[int64]dto.Doctor{
				1: {ID: 1, Name: "Dr. Smith"},
			},
			subscribersMap: map[int64]indto.GetSubscribersByDoctorIDsResponse{},
			expected: map[int64]dto.Doctor{
				1: {ID: 1, Name: "Dr. Smith"},
			},
		},
		{
			name: "Несколько докторов",
			doctorsMap: map[int64]dto.Doctor{
				1: {ID: 1, Name: "Dr. Smith"},
				2: {ID: 2, Name: "Dr. Johnson"},
			},
			subscribersMap: map[int64]indto.GetSubscribersByDoctorIDsResponse{
				1: {
					TgSubsCount:     "150",
					TgSubsCountText: "подписчиков",
				},
			},
			expected: map[int64]dto.Doctor{
				1: {
					ID:              1,
					Name:            "Dr. Smith",
					TgSubsCount:     "150",
					TgSubsCountText: "подписчиков",
				},
				2: {ID: 2, Name: "Dr. Johnson"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			doctorsCopy := make(map[int64]dto.Doctor, len(tt.doctorsMap))
			for k, v := range tt.doctorsMap {
				doctorsCopy[k] = v
			}

			enrichSubscribers(context.Background(), doctorsCopy, tt.subscribersMap)

			assert.Equal(t, tt.expected, doctorsCopy)
		})
	}
}

func TestEnrichAdditionalCities(t *testing.T) {
	tests := []struct {
		name                string
		doctorsMap          map[int64]dto.Doctor
		additionalCitiesMap map[int64][]*city.City
		expected            map[int64]dto.Doctor
	}{
		{
			name:                "Пустые мапы",
			doctorsMap:          map[int64]dto.Doctor{},
			additionalCitiesMap: map[int64][]*city.City{},
			expected:            map[int64]dto.Doctor{},
		},
		{
			name: "Доктор без доп городов",
			doctorsMap: map[int64]dto.Doctor{
				1: {ID: 1, MainCityID: 10},
			},
			additionalCitiesMap: map[int64][]*city.City{
				1: {city.BuildCity(city.WithID(10), city.WithName("мск"))},
			},
			expected: map[int64]dto.Doctor{
				1: {ID: 1, MainCityID: 10, City: "мск"},
			},
		},
		{
			name: "Доктор с основным и дополнительными городами",
			doctorsMap: map[int64]dto.Doctor{
				1: {ID: 1, MainCityID: 10},
			},
			additionalCitiesMap: map[int64][]*city.City{
				1: {
					city.BuildCity(city.WithID(10), city.WithName("мск")),
					city.BuildCity(city.WithID(20), city.WithName("спб")),
					city.BuildCity(city.WithID(30), city.WithName("екб")),
				},
			},
			expected: map[int64]dto.Doctor{
				1: {ID: 1, MainCityID: 10, City: "мск, спб, екб"},
			},
		},
		{
			name: "Разные доктора с разными городами",
			doctorsMap: map[int64]dto.Doctor{
				1: {ID: 1, MainCityID: 10},
				2: {ID: 2, MainCityID: 20},
			},
			additionalCitiesMap: map[int64][]*city.City{
				1: {
					city.BuildCity(city.WithID(10), city.WithName("мск")),
					city.BuildCity(city.WithID(30), city.WithName("екб"))},
				2: {
					city.BuildCity(city.WithID(20), city.WithName("спб")),
				},
			},
			expected: map[int64]dto.Doctor{
				1: {ID: 1, MainCityID: 10, City: "мск, екб"},
				2: {ID: 2, MainCityID: 20, City: "спб"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Создаем копию исходного map
			doctorsCopy := make(map[int64]dto.Doctor, len(tt.doctorsMap))
			for k, v := range tt.doctorsMap {
				doctorsCopy[k] = v
			}

			enrichAdditionalCities(ctx, doctorsCopy, tt.additionalCitiesMap)

			assert.Equal(t, tt.expected, doctorsCopy)
		})
	}
}
