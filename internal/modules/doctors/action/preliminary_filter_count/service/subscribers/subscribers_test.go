package subscribers

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/service/subscribers/mocks"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	subscribersGetter *mocks.MockSubscribersGetter
}

func p(t *testing.T) fields {
	f := fields{
		subscribersGetter: mocks.NewMockSubscribersGetter(gomock.NewController(t)),
	}
	return f
}

func TestService_FilterDoctorsBySubscribers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		filter         dto.Filter
		mockResponse   indto.GetDoctorsByFilterResponse
		mockError      error
		expectedResult int64
		expectedError  error
	}{
		{
			name: "successful filtering",
			filter: dto.Filter{
				MinSubscribers: 100,
				MaxSubscribers: 1000,
				SocialMedia:    []string{"telegram", "vk"},
			},
			mockResponse: indto.GetDoctorsByFilterResponse{
				DoctorsCount: 3,
			},
			mockError:      nil,
			expectedResult: 3,
			expectedError:  nil,
		},
		{
			name: "empty result",
			filter: dto.Filter{
				MinSubscribers: 1000,
				MaxSubscribers: 2000,
				SocialMedia:    []string{"telegram"},
			},
			mockResponse:   indto.GetDoctorsByFilterResponse{},
			mockError:      nil,
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name: "error from getter",
			filter: dto.Filter{
				MinSubscribers: 100,
				MaxSubscribers: 1000,
				SocialMedia:    []string{"vk"},
			},
			mockResponse:   indto.GetDoctorsByFilterResponse{},
			mockError:      errors.New("database error"),
			expectedResult: 0,
			expectedError:  errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := p(t)
			service := New(deps.subscribersGetter)

			expectedRequest := indto.GetDoctorsByFilterRequest{
				MinSubscribers: tt.filter.MinSubscribers,
				MaxSubscribers: tt.filter.MaxSubscribers,
				SocialMedia: lo.Map(tt.filter.SocialMedia, func(socialMedia string, index int) indto.SocialMedia {
					return indto.NewSocialMedia(socialMedia)
				}),
			}

			deps.subscribersGetter.EXPECT().GetDoctorsByFilter(gomock.Any(), expectedRequest).Return(tt.mockResponse, tt.mockError)

			result, err := service.FilterDoctorsBySubscribers(context.Background(), tt.filter)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
