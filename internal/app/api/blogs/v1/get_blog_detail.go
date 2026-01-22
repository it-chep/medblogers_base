package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"
)

func (i *Implementation) GetBlogDetail(ctx context.Context, req *desc.GetBlogDetailRequest) (*desc.GetBlogDetailResponse, error) {
	blogDTO, err := i.blogs.Actions.GetBlogDetail.Do(ctx, req.GetBlogSlug())
	if err != nil {
		return nil, err
	}

	blog := blogDTO.BlogEntity
	doc := blogDTO.Doctor
	categories := blogDTO.Categories

	return &desc.GetBlogDetailResponse{
		Title:             blog.GetTitle(),
		Slug:              blog.GetSlug(),
		Body:              blog.GetBody(),
		PreviewText:       blog.GetPreviewText(),
		SocietyPreview:    blog.GetSocietyPreviewText(),
		AdditionalSeoText: blog.GetAdditionalSEOText(),
		CreatedAt:         converter.FormatDateRussian(blog.GetCreatedAt()),
		PhotoLink:         blog.GetPrimaryPhotoURL(),

		Doctor: &desc.GetBlogDetailResponse_Doctor{
			Name:           doc.Name,
			Slug:           doc.Slug,
			Image:          doc.PhotoLink,
			SpecialityName: doc.SpecialityName,
		},

		Categories: lo.Map(categories, func(item *category.Category, _ int) *desc.Category {
			return &desc.Category{
				Id:        item.ID(),
				Name:      item.Name(),
				FontColor: item.FontColor(),
				BgColor:   item.BgColor(),
			}
		}),
	}, nil
}
