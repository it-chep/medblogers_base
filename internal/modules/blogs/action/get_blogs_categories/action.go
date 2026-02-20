package get_blogs_categories

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blogs_categories/dal"
	"medblogers_base/internal/modules/blogs/action/get_blogs_categories/dto"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение списка категорий статей
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do .
func (a *Action) Do(ctx context.Context) (dto.GetCategoriesResponse, error) {
	var resp dto.GetCategoriesResponse
	categories, err := a.dal.GetAllCategories(ctx)
	if err != nil {
		return resp, err
	}
	resp.Categories = categories

	count, err := a.dal.AllBlogsCount(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получения общего количества статей", err)
		return resp, err
	}
	resp.AllBlogsCount = count

	blogsCountMap, err := a.dal.CategoryBlogsCount(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получения количества статей по категориям", err)
		return resp, nil
	}
	resp.CategoryCountMap = blogsCountMap

	return resp, nil
}
