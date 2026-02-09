package get_doctor_blogs

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_doctor_blogs/dal"
	"medblogers_base/internal/modules/blogs/action/get_doctor_blogs/dto"
	"medblogers_base/internal/modules/blogs/dal/blogs"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение статей врача
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
func (a *Action) Do(ctx context.Context, doctorSlug string) (dto.Response, error) {
	// Получаем статьи
	blogs, err := a.dal.GetDoctorBlogs(ctx, doctorSlug)
	if err != nil {
		return dto.Response{}, err
	}

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
