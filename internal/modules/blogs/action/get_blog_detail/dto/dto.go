package dto

import (
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/modules/blogs/domain/category"
)

type BlogDTO struct {
	BlogEntity *blog.Blog
	Categories category.Categories
	Doctor     DoctorAuthorDTO
}

type DoctorAuthorDTO struct {
	PhotoLink      string
	Name           string
	Slug           string
	SpecialityName string
}
