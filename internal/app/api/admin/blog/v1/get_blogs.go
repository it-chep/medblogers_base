package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blogs/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetBlogs(ctx context.Context, req *desc.GetBlogsRequest) (resp *desc.GetBlogsResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog", func(ctx context.Context) error {
		resp = &desc.GetBlogsResponse{}

		blogs, err := i.admin.Actions.BlogModule.GetBlogs.Do(ctx)
		if err != nil {
			return err
		}

		resp.Blogs = lo.Map(blogs, func(item dto.Blog, _ int) *desc.GetBlogsResponse_Blog {
			return &desc.GetBlogsResponse_Blog{
				BlogId:         item.BlogID.String(),
				Title:          item.Name,
				IsActive:       item.IsActive.Bool,
				OrderingNumber: item.OrderingNumber.Int64,
			}
		})

		return nil
	})
}
