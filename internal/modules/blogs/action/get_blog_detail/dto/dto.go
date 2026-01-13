package dto

import "medblogers_base/internal/modules/blogs/domain/blog"

type BlogDTO struct {
	BlogEntity *blog.Blog
	Doctor     DoctorAuthorDTO
}

type DoctorAuthorDTO struct {
	PhotoLink      string
	Name           string
	Slug           string
	SpecialityName string
}
