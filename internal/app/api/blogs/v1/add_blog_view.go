package v1

import (
	"context"

	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
)

// AddBlogView добавление просмотра статьи.
func (i *Implementation) AddBlogView(ctx context.Context, req *desc.AddBlogViewRequest) (*desc.AddBlogViewResponse, error) {
	err := i.blogs.Actions.AddBlogView.Do(ctx, req.GetBlogSlug(), req.GetCookieId())
	if err != nil {
		return nil, err
	}

	return &desc.AddBlogViewResponse{}, nil
}
