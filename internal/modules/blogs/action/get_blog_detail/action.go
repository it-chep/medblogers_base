package get_blog_detail

import (
	"context"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/dal"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/dto"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/service/doctor"
	"medblogers_base/internal/modules/blogs/client"
	"medblogers_base/internal/modules/blogs/dal/blogs"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type Config interface {
	GetS3Config() config.S3Config
}

// Action получение карточки статьи
type Action struct {
	dal           *dal.Repository
	doctorService *doctor.Service
	commonDal     *blogs.Repository
	config        Config
}

// New .
func New(pool postgres.PoolWrapper, clients *client.Aggregator, cfg config.AppConfig) *Action {
	repo := dal.NewRepository(pool)
	return &Action{
		dal:           repo,
		commonDal:     blogs.NewRepository(pool),
		doctorService: doctor.New(repo, clients.S3),
		config:        cfg,
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

	bucket := a.config.GetS3Config().Bucket.Blogs
	if photo != nil {
		blogEntity.SetPrimaryPhotoURL(bucket, photo.GetID(), photo.GetFileType())
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

	viewsMap, err := a.commonDal.GetBlogsViewsCount(ctx, []uuid.UUID{blogEntity.GetID()})
	if err != nil {
		return dto.BlogDTO{}, err
	}

	blogEntity.SetViewsCount(viewsMap[blogEntity.GetID()])

	return dto.BlogDTO{
		BlogEntity: blogEntity,
		Doctor:     doctorInfo,
		Categories: categories,
	}, nil
}
