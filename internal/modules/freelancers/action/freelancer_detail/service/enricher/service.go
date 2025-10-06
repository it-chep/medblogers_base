package enricher

import "context"

//go:generate mockgen -destination=mocks/image_mocks.go -package=mocks . ImageGetter

type ImageGetter interface {
	GeneratePresignedURL(ctx context.Context, s3Key string) (string, error)
}

type ImageEnricher struct {
	imageGetter ImageGetter
}

func NewImageEnricher(imageGetter ImageGetter) *ImageEnricher {
	return &ImageEnricher{
		imageGetter: imageGetter,
	}
}

func (e *ImageEnricher) Enrich(ctx context.Context, image string) (string, error) {
	freelancerImageURL, err := e.imageGetter.GeneratePresignedURL(ctx, image)
	if err != nil {
		return "", err
	}
	return freelancerImageURL, nil
}
