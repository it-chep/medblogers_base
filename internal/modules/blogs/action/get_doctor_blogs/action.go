package get_doctor_blogs

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_doctor_blogs/dal"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение статей врача
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do .
func (a *Action) Do(ctx context.Context, doctorSlug string) (blog.Blogs, error) {
	// Получаем статьи
	blogs, err := a.dal.GetDoctorBlogs(ctx, doctorSlug)
	if err != nil {
		return nil, err
	}

	// Получаем первые фотографии статей
	blogPhotosMap, err := a.dal.GetPrimaryPhotos(ctx, blogs.GetIDs())
	if err != nil {
		return nil, err
	}

	// Устанавливаем фотографию если она есть
	for _, bl := range blogs {
		photo, ok := blogPhotosMap[bl.GetID()]
		if !ok {
			continue
		}

		bl.SetPrimaryPhotoURL(photo.GetID(), photo.GetFileType())
	}

	return blogs, nil
}
