package search_blogs

import (
	"context"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/blogs/action/search_blogs/dal"
	"medblogers_base/internal/modules/blogs/action/search_blogs/dto"
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

func (a *Action) Do(ctx context.Context, query string) (dto.Response, error) {
	blogsResp, err := a.dal.SearchBlogs(ctx, query)
	if err != nil {
		return dto.Response{}, err
	}

	blogPhotosMap, err := a.commonDal.GetPrimaryPhotos(ctx, blogsResp.GetIDs())
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
		resp.Blogs = append(resp.Blogs, bl)
	}

	return resp, nil
}
