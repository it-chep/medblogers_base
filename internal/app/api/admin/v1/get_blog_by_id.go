package v1

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
)

func (i *Implementation) GetBlogByID(ctx context.Context, req *desc.GetBlogByIDRequest) (resp *desc.GetBlogByIDResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}", func(ctx context.Context) error {
		blog, err := i.admin.Actions.BlogModule.GetBlogByID.Do(ctx, uuid.MustParse(req.GetBlogId()))
		if err != nil {
			return err
		}

		resp = &desc.GetBlogByIDResponse{
			BlogId:            blog.ID.String(),
			Slug:              blog.Slug.String,
			Title:             blog.Name,
			Body:              blog.Body.String,
			IsActive:          blog.IsActive.Bool,
			PreviewText:       blog.PreviewText.String,
			SocietyPreview:    blog.SocietyPreviewText.String,
			AdditionalSeoText: blog.AdditionalSEOText.String,
		}

		return nil
	})
}
