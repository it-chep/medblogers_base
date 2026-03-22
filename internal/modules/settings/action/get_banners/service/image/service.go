package image

import "fmt"

type Service struct {
	bucket string
}

func New(bucket string) *Service {
	return &Service{
		bucket: bucket,
	}
}

func (s *Service) GetPhotoLink(imageID, fileType string) string {
	if imageID == "" || fileType == "" || s.bucket == "" {
		return ""
	}

	return fmt.Sprintf("https://storage.yandexcloud.net/%s/images/%s.%s", s.bucket, imageID, fileType)
}
