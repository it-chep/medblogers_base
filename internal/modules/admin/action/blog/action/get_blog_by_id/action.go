package get_blog_by_id

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/get_blog_by_id/dal"
	"medblogers_base/internal/modules/admin/action/blog/action/get_blog_by_id/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

// Action .
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do получение статьи по ID
func (a *Action) Do(ctx context.Context, blogID uuid.UUID) (dto.Response, error) {
	blog, err := a.dal.GetBlogByID(ctx, blogID)
	if err != nil {
		return dto.Response{}, err
	}

	var doctorName string
	if blog.DoctorID.Valid {
		doctorName, err = a.dal.GetDoctorInfo(ctx, blog.DoctorID.Int64)
		if err != nil {
			return dto.Response{}, err
		}
	}

	categories, err := a.dal.GetBlogCategories(ctx, blogID)
	if err != nil {
		return dto.Response{}, err
	}

	return dto.Response{
		Blog:       blog,
		DoctorName: doctorName,
		Categories: categories,
	}, nil
}
