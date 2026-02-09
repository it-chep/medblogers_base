package filter_blogs

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/filter_blogs/dal"
	"medblogers_base/internal/modules/blogs/action/filter_blogs/dto"
	"medblogers_base/internal/modules/blogs/dal/blogs"
	"medblogers_base/internal/pkg/postgres"
)

// Action фильтрация статей по параметрам
type Action struct {
	dal       *dal.Repository
	commonDal *blogs.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: blogs.NewRepository(pool),
		dal:       dal.NewRepository(pool),
	}
}

// Do .
func (a *Action) Do(ctx context.Context, req dto.FilterRequest) (dto.Response, error) {
	// фильтруем статьи
	blogs, err := a.dal.FilterBlogs(ctx, req)
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

	resp := dto.Response{}
	for _, bl := range blogs {
		photo, ok := blogPhotosMap[bl.GetID()]
		if !ok {
			continue
		}
		bl.SetPrimaryPhotoURL(photo.GetID(), photo.GetFileType())
		resp.Blogs = append(resp.Blogs, dto.Blog{
			Blog:       bl,
			Categories: categoriesMap[bl.GetID()],
		})
	}

	return resp, nil
}
