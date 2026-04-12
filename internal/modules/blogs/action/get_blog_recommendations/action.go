package get_blog_recommendations

import (
	"context"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/blogs/action/get_blog_recommendations/dal"
	"medblogers_base/internal/modules/blogs/action/get_blog_recommendations/dto"
	"medblogers_base/internal/modules/blogs/dal/blogs"
	"medblogers_base/internal/pkg/postgres"
)

type Config interface {
	GetS3Config() config.S3Config
}

type Action struct {
	dal       *dal.Repository
	commonDal *blogs.Repository
	config    Config
}

func New(pool postgres.PoolWrapper, cfg config.AppConfig) *Action {
	return &Action{
		dal:       dal.NewRepository(pool),
		commonDal: blogs.NewRepository(pool),
		config:    cfg,
	}
}

func (a *Action) Do(ctx context.Context, slug string) (dto.Response, error) {
	blogsResp, err := a.dal.GetBlogRecommendations(ctx, slug)
	if err != nil {
		return dto.Response{}, err
	}

	if len(blogsResp) == 0 {
		blogsResp, err = a.dal.GetTopBlogsFallback(ctx)
		if err != nil {
			return dto.Response{}, err
		}
	}

	blogPhotosMap, err := a.commonDal.GetPrimaryPhotos(ctx, blogsResp.GetIDs())
	if err != nil {
		return dto.Response{}, err
	}

	categoriesMap, err := a.commonDal.GetBlogsCategories(ctx, blogsResp.GetIDs())
	if err != nil {
		return dto.Response{}, err
	}

	viewsMap, err := a.commonDal.GetBlogsViewsCount(ctx, blogsResp.GetIDs())
	if err != nil {
		return dto.Response{}, err
	}

	bucket := a.config.GetS3Config().Bucket.Blogs

	resp := dto.Response{}
	for _, bl := range blogsResp {
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
