package enricher

import (
	"context"
	"errors"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/enricher/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func p(t *testing.T) fields {
	return fields{
		imageGetterMock: mocks.NewMockImageGetter(gomock.NewController(t)),
	}
}

type fields struct {
	imageGetterMock *mocks.MockImageGetter
}

func TestImageEnricher_Enrich(t *testing.T) {
	t.Parallel()

	s3Key := "medblogers/images/zag.png"
	t.Run("успешное обогащение фотографиями", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		deps.imageGetterMock.EXPECT().GeneratePresignedURL(gomock.Any(), s3Key).Return("some_image", nil)
		service := NewImageEnricher(deps.imageGetterMock)
		image, err := service.Enrich(context.Background(), s3Key)

		assert.NoError(t, err)
		assert.Equal(t, "some_image", image)
	})

	t.Run("Произошла ошибка и фотография не пришла", func(t *testing.T) {
		t.Parallel()

		deps := p(t)
		deps.imageGetterMock.EXPECT().GeneratePresignedURL(gomock.Any(), s3Key).Return("", errors.New("error"))
		service := NewImageEnricher(deps.imageGetterMock)
		image, err := service.Enrich(context.Background(), s3Key)

		assert.Error(t, err)
		assert.Equal(t, "", image)
	})
}
