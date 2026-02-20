package dto

import "medblogers_base/internal/modules/blogs/domain/category"

type GetCategoriesResponse struct {
	Categories       category.Categories
	CategoryCountMap map[int64]int64
	AllBlogsCount    int64
}
