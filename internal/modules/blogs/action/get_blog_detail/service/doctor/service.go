package doctor

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/dto"
	"medblogers_base/internal/modules/blogs/domain/doctor_author"
)

type Dal interface {
	GetDoctorInfo(ctx context.Context, doctorID int64) (*doctor_author.Doctor, error)
}

type ImageGetter interface {
	GetPhotoLink(s3Key string) string
}

type Service struct {
	dal   Dal
	image ImageGetter
}

func New(dal Dal, image ImageGetter) *Service {
	return &Service{
		dal:   dal,
		image: image,
	}
}

// GetDoctorToBlog получение информации о докторе для отображения в статье
func (s *Service) GetDoctorToBlog(ctx context.Context, doctorID int64) (dto.DoctorAuthorDTO, error) {
	doctor, err := s.dal.GetDoctorInfo(ctx, doctorID)
	if err != nil {
		return dto.DoctorAuthorDTO{}, err
	}

	photoLink := s.image.GetPhotoLink(doctor.GetS3Key())

	return dto.DoctorAuthorDTO{
		PhotoLink:      photoLink,
		Name:           doctor.GetName(),
		Slug:           doctor.GetSlug(),
		SpecialityName: doctor.GetSpecialityName(),
	}, err
}
