package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/blogs/domain/blog"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"
)

func (i *Implementation) GetDoctorBlogs(ctx context.Context, req *desc.GetDoctorBlogsRequest) (*desc.GetDoctorBlogsResponse, error) {
	blogs, err := i.blogs.Actions.GetDoctorBlogs.Do(ctx, req.GetDoctorSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetDoctorBlogsResponse{
		Blogs: lo.Map(blogs, func(item *blog.Blog, index int) *desc.GetDoctorBlogsResponse_BlogMiniatures {
			return &desc.GetDoctorBlogsResponse_BlogMiniatures{
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				PreviewText: item.GetPreviewText(),
				CreatedAt:   converter.FormatDateRussian(item.GetCreatedAt()),
				PhotoLink:   item.GetPrimaryPhotoURL(),
			}
		}),
	}, nil
}
