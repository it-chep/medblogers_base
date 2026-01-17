package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/blogs/action/get_doctor_blogs/dto"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"
)

func (i *Implementation) GetDoctorBlogs(ctx context.Context, req *desc.GetDoctorBlogsRequest) (*desc.GetDoctorBlogsResponse, error) {
	resp, err := i.blogs.Actions.GetDoctorBlogs.Do(ctx, req.GetDoctorSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetDoctorBlogsResponse{
		Blogs: lo.Map(resp.Blogs, func(item dto.Blog, index int) *desc.GetDoctorBlogsResponse_BlogMiniatures {
			return &desc.GetDoctorBlogsResponse_BlogMiniatures{
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				PreviewText: item.GetPreviewText(),
				CreatedAt:   converter.FormatDateRussian(item.GetCreatedAt()),
				PhotoLink:   item.GetPrimaryPhotoURL(),

				Categories: lo.Map(item.Categories, func(item *category.Category, _ int) *desc.Category {
					return &desc.Category{
						Id:        item.ID(),
						Name:      item.Name(),
						FontColor: item.FontColor(),
						BgColor:   item.BgColor(),
					}
				}),
			}
		}),
	}, nil
}
