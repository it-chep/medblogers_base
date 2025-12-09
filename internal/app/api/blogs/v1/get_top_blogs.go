package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/blogs/domain/blog"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"
)

func (i *Implementation) GetTopBlogs(ctx context.Context, req *desc.GetTopBlogsRequest) (*desc.GetTopBlogsResponse, error) {
	blogs, err := i.blogs.Actions.GetTopBlogs.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetTopBlogsResponse{
		Blogs: lo.Map(blogs, func(item *blog.Blog, index int) *desc.GetTopBlogsResponse_BlogMiniatures {
			return &desc.GetTopBlogsResponse_BlogMiniatures{
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				PreviewText: item.GetPreviewText(),
				CreatedAt:   converter.FormatDateRussian(item.GetCreatedAt()),
				PhotoLink:   item.GetPrimaryPhotoURL(),
			}
		}),
	}, nil
}
