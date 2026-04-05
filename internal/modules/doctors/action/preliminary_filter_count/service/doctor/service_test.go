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
	storage   *mocks.MockStorage
	commonDal *mocks.MockCommonDal
}

func p(t *testing.T) fields {
	ctrl := gomock.NewController(t)
	f := fields{
		storage:   mocks.NewMockStorage(ctrl),
		commonDal: mocks.NewMockCommonDal(ctrl),
	}
	return f
}

func TestService_GetDoctorsIDsByFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		filter      dto.Filter
		mockIDs     []int64
		mockError   error
		expectedIDs []int64
	}{
		{
			name: "successful count",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			mockIDs:     []int64{10, 20, 30},
			mockError:   nil,
			expectedIDs: []int64{10, 20, 30},
		},
		{
			name: "empty result",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			mockIDs:     []int64{},
			mockError:   nil,
			expectedIDs: []int64{},
		},
		{
			name: "database error",
			filter: dto.Filter{
				Specialities: []int64{1, 2, 3},
			},
			mockIDs:     nil,
			mockError:   errors.New("database connection failed"),
			expectedIDs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := p(t)
			service := New(deps.storage, deps.commonDal)

			deps.storage.EXPECT().FilterDoctors(gomock.Any(), tt.filter).
				Return(tt.mockIDs, tt.mockError)

			ids, err := service.GetDoctorsIDsByFilter(context.Background(), tt.filter)

			if tt.mockError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.mockError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedIDs, ids)
		})
	}
}
