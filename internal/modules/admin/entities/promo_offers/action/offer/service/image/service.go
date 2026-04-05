package image

type Image interface {
	GetBrandPhotoLink(s3Key string) string
}

type Service struct {
	image Image
}

func New(image Image) *Service {
	return &Service{image: image}
}

func (s *Service) GetImageURL(s3Key string) string {
	return s.image.GetBrandPhotoLink(s3Key)
}
