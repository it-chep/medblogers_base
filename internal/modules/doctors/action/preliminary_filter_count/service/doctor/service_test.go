package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/service/doctor/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	storage *mocks.MockStorage
}

func p(t *testing.T) fields {
	f := fields{
		storage: mocks.NewMockStorage(gomock.NewController(t)),
	}
	return f
}

func TestService_GetDoctorsByFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		filter        dto.Filter
		mockCount     int64
		mockError     error
		expectedCount int64
		expectedError error
	}{
		{
			name: "successful count",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			mockCount:     42,
			mockError:     nil,
			expectedCount: 42,
			expectedError: nil,
		},
		{
			name: "empty result",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			mockCount:     0,
			mockError:     nil,
			expectedCount: 0,
			expectedError: nil,
		},
		{
			name: "database error",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			mockCount:     0,
			mockError:     errors.New("database connection failed"),
			expectedCount: 0,
			expectedError: errors.New("database connection failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := p(t)
			service := New(deps.storage)

			deps.storage.EXPECT().CountFilterDoctors(gomock.Any(), tt.filter).
				Return(tt.mockCount, tt.mockError)

			count, err := service.GetDoctorsByFilter(context.Background(), tt.filter)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedCount, count)
		})
	}
}

func TestService_GetDoctorsByFilterAndIDs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		filter        dto.Filter
		ids           []int64
		mockCount     int64
		mockError     error
		expectedCount int64
		expectedError error
	}{
		{
			name: "successful count with ids",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			ids:           []int64{1, 2, 3, 4},
			mockCount:     3,
			mockError:     nil,
			expectedCount: 3,
			expectedError: nil,
		},
		{
			name: "no matches with given ids",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			ids:           []int64{5, 6, 7},
			mockCount:     0,
			mockError:     nil,
			expectedCount: 0,
			expectedError: nil,
		},
		{
			name: "empty ids slice",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			ids:           []int64{},
			mockCount:     0,
			mockError:     nil,
			expectedCount: 0,
			expectedError: nil,
		},
		{
			name: "database error",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			ids:           []int64{1, 2, 3},
			mockCount:     0,
			mockError:     errors.New("query execution failed"),
			expectedCount: 0,
			expectedError: errors.New("query execution failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			deps := p(t)
			service := New(deps.storage)

			deps.storage.EXPECT().CountByFilterAndIDs(gomock.Any(), tt.filter, tt.ids).
				Return(tt.mockCount, tt.mockError)

			count, err := service.GetDoctorsByFilterAndIDs(context.Background(), tt.filter, tt.ids)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedCount, count)
		})
	}
}
