package get_blogs

import (
	"context"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/blogs/action/get_blogs/dal"
	"medblogers_base/internal/modules/blogs/action/get_blogs/dto"
	"medblogers_base/internal/modules/blogs/dal/blogs"
	"medblogers_base/internal/pkg/postgres"
)

type Config interface {
	GetS3Config() config.S3Config
}

// Action получение всех статей
type Action struct {
	dal       *dal.Repository
	commonDal *blogs.Repository
	config    Config
}

// New .
func New(pool postgres.PoolWrapper, cfg config.AppConfig) *Action {
	return &Action{
		commonDal: blogs.NewRepository(pool),
		dal:       dal.NewRepository(pool),
		config:    cfg,
	}
}

// Do .
func (a *Action) Do(ctx context.Context) (dto.Response, error) {
	// Получаем статьи
	blogs, err := a.dal.GetBlogs(ctx)
	if err != nil {
		return dto.Response{}, err
	}

	// Получаем первые фотографии статей
	blogPhotosMap, err := a.commonDal.GetPrimaryPhotos(ctx, blogs.GetIDs())
	if err != nil {
		return dto.Response{}, err
	}

	categoriesMap, err := a.commonDal.GetBlogsCategories(ctx, blogs.GetIDs())
	if err != nil {
		return dto.Response{}, err
	}

	viewsMap, err := a.commonDal.GetBlogsViewsCount(ctx, blogs.GetIDs())
	if err != nil {
		return dto.Response{}, err
	}

	bucket := a.config.GetS3Config().Bucket.Blogs
	resp := dto.Response{}
	for _, bl := range blogs {
		photo, ok := blogPhotosMap[bl.GetID()]
		if !ok {
			continue
		}

		bl.SetViewsCount(viewsMap[bl.GetID()])
		bl.SetPrimaryPhotoURL(bucket, photo.GetID(), photo.GetFileType())
		resp.Blogs = append(resp.Blogs, dto.Blog{
			Blog:       bl,
			Categories: categoriesMap[bl.GetID()],
		})
	}

	return resp, nil
}
