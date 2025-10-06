package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetPagesCount(t *testing.T) {
	// максимальное количество докторов на странице - 30

	t.Parallel()

	service := New()
	t.Run("Успешный подсчет количества страниц", func(t *testing.T) {
		t.Parallel()
		expectedCount := int64(3)

		pageCount := service.GetPagesCount(90)

		assert.Equal(t, expectedCount, pageCount)
	})

	t.Run("Успешный подсчет количества страниц", func(t *testing.T) {
		t.Parallel()
		expectedCount := int64(3)

		pageCount := service.GetPagesCount(int64(65))

		assert.Equal(t, expectedCount, pageCount)
	})

	t.Run("Ошибка из базы", func(t *testing.T) {
		t.Parallel()
		expectedCount := int64(1)

		pageCount := service.GetPagesCount(0)

		assert.Equal(t, expectedCount, pageCount)
	})
}
