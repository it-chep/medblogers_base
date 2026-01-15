package get_blog_detail

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/dal"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/dto"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/service/doctor"
	"medblogers_base/internal/modules/blogs/client"
	"medblogers_base/internal/modules/blogs/dal/blogs"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение карточки статьи
type Action struct {
	dal           *dal.Repository
	doctorService *doctor.Service
	commonDal     *blogs.Repository
}

// New .
func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	repo := dal.NewRepository(pool)
	return &Action{
		dal:           repo,
		commonDal:     blogs.NewRepository(pool),
		doctorService: doctor.New(repo, clients.S3),
	}
}

// Do .
func (a *Action) Do(ctx context.Context, slug string) (dto.BlogDTO, error) {
	blogEntity, err := a.dal.GetBlogDetail(ctx, slug)
	if err != nil {
		return dto.BlogDTO{}, err
	}

	photo, err := a.dal.GetPrimaryPhoto(ctx, blogEntity.GetID())
	if err != nil {
		return dto.BlogDTO{}, err
	}

	if photo != nil {
		blogEntity.SetPrimaryPhotoURL(photo.GetID(), photo.GetFileType())
	}

	var doctorInfo dto.DoctorAuthorDTO
	if blogEntity.HasAuthor() {
		doctorInfo, err = a.doctorService.GetDoctorToBlog(ctx, blogEntity.GetDoctorID())
		if err != nil {
			return dto.BlogDTO{}, err
		}
	}

	categories, err := a.commonDal.GetBlogCategories(ctx, blogEntity.GetID())
	if err != nil {
		return dto.BlogDTO{}, err
	}

	return dto.BlogDTO{
		BlogEntity: blogEntity,
		Doctor:     doctorInfo,
		Categories: categories,
	}, nil
}
