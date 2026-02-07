package rules

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/publish_blog/dto"
)

// RuleFieldsAvailableToPublish проверяет что все необходимые поля заполнены
var RuleFieldsAvailableToPublish = func() func(_ context.Context, blog *dto.Blog) (bool, error) {
	return func(_ context.Context, blog *dto.Blog) (bool, error) {
		if blog.Slug.String == "" {
			return false, dto.ValidationError{
				Text:  "Необходимо указать поле СЛАГ",
				Field: "slug",
			}
		}

		if blog.PreviewText.String == "" {
			return false, dto.ValidationError{
				Text:  "Необходимо указать поле Превью",
				Field: "previewText",
			}
		}

		if blog.SocietyPreviewText.String == "" {
			return false, dto.ValidationError{
				Text:  "Необходимо указать поле Превью для соцсетей",
				Field: "societyPreview",
			}
		}
		return true, dto.ValidationError{}
	}
}

// RuleNotPublished проверяет что статья еще не опубликована
var RuleNotPublished = func() func(_ context.Context, blog *dto.Blog) (bool, error) {
	return func(_ context.Context, blog *dto.Blog) (bool, error) {
		if blog.IsActive.Bool {
			return false, dto.ValidationError{
				Text:  "Статья уже опубликована",
				Field: "isActive",
			}
		}
		return true, dto.ValidationError{}
	}
}
