package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/action/blog/action/get_blog_by_id/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"

	"github.com/samber/lo"

	"github.com/google/uuid"
)

func (i *Implementation) GetBlogByID(ctx context.Context, req *desc.GetBlogByIDRequest) (resp *desc.GetBlogByIDResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}", func(ctx context.Context) error {
		blogDTO, err := i.admin.Actions.BlogModule.GetBlogByID.Do(ctx, uuid.MustParse(req.GetBlogId()))
		if err != nil {
			return err
		}

		blog := blogDTO.Blog

		resp = &desc.GetBlogByIDResponse{
			BlogId:            blog.ID.String(),
			Slug:              blog.Slug.String,
			Title:             blog.Name,
			Body:              blog.Body.String,
			IsActive:          blog.IsActive.Bool,
			PreviewText:       blog.PreviewText.String,
			SocietyPreview:    blog.SocietyPreviewText.String,
			AdditionalSeoText: blog.AdditionalSEOText.String,
			OrderingNumber:    blog.OrderingNumber.Int64,

			Doctor: &desc.GetBlogByIDResponse_Doctor{
				DoctorId:   blog.DoctorID.Int64,
				DoctorName: blogDTO.DoctorName,
			},

			Categories: lo.Map(blogDTO.Categories, func(item dto.Category, index int) *desc.GetBlogByIDResponse_Category {
				return &desc.GetBlogByIDResponse_Category{
					Id:   item.ID,
					Name: item.Name,
				}
			}),
		}

		return nil
	})
}
