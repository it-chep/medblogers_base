package page

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/page/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestService_GetPagesCount(t *testing.T) {
	// максимальное количество докторов на странице - 30

	t.Parallel()

	mockStorage := mocks.NewMockStorage(gomock.NewController(t))
	service := New(mockStorage)
	t.Run("Успешный подсчет количества страниц", func(t *testing.T) {
		t.Parallel()
		mockStorage.EXPECT().GetDoctorsCountByFilter(gomock.Any(), gomock.Any()).Return(int64(90), nil)
		expectedCount := int64(3)

		pageCount, err := service.GetPagesCount(context.Background(), dto.Filter{})

		assert.NoError(t, err)
		assert.Equal(t, expectedCount, pageCount)
	})

	t.Run("Успешный подсчет количества страниц", func(t *testing.T) {
		t.Parallel()
		mockStorage.EXPECT().GetDoctorsCountByFilter(gomock.Any(), gomock.Any()).Return(int64(65), nil)
		expectedCount := int64(3)

		pageCount, err := service.GetPagesCount(context.Background(), dto.Filter{})

		assert.NoError(t, err)
		assert.Equal(t, expectedCount, pageCount)
	})

	t.Run("Ошибка из базы", func(t *testing.T) {
		t.Parallel()
		mockStorage.EXPECT().GetDoctorsCountByFilter(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("not found"))
		expectedCount := int64(1)

		pageCount, err := service.GetPagesCount(context.Background(), dto.Filter{})

		assert.Error(t, err)
		assert.Equal(t, expectedCount, pageCount)
	})
}
