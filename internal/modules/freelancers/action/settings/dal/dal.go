package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/freelancers/action/settings/dto"
	"medblogers_base/internal/pkg/postgres"
	"strings"
)

// Repository специальности
type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе со специальностями
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

type PriceCategoryDAO struct {
	ID               int64 `db:"id" json:"id"`
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}

// GetPriceCategoriesInfo - получение информации о ценовых категориях
func (r *Repository) GetPriceCategoriesInfo(ctx context.Context) ([]dto.PriceCategory, error) {
	sql := `
		select price_category as id,
       		count(*) as freelancers_count
		from freelancer
		where is_active is true
		group by id
		order by id
	`

	var categories []PriceCategoryDAO
	err := pgxscan.Select(ctx, r.db, categories, sql)
	if err != nil {
		return nil, err
	}

	domain := make([]dto.PriceCategory, 0, len(categories))
	for _, category := range categories {
		domain = append(domain, dto.PriceCategory{
			ID:               category.ID,
			Name:             strings.Repeat("₽", int(category.ID)),
			FreelancersCount: category.FreelancersCount,
		})
	}

	return domain, nil
}
