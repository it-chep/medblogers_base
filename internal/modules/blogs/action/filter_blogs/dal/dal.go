package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"
	"medblogers_base/internal/modules/blogs/action/filter_blogs/dto"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/pkg/postgres"
	"strings"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе со статьями
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// FilterBlogs фильтрация статей по параметрам
func (r *Repository) FilterBlogs(ctx context.Context, filter dto.FilterRequest) (blog.Blogs, error) {
	sql, phValues := sqlStmt(filter)

	var blogs dao.BlogMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &blogs, sql, phValues...)
	if err != nil {
		return nil, err
	}

	return blogs.ToDomain(), nil
}

// sqlStmt к-ор запроса
func sqlStmt(filter dto.FilterRequest) (_ string, phValues []any) {
	defaultSql := `
		select distinct b.id, b.name, b.slug, b.preview_text, b.created_at, b.ordering_number 
		from blog b 
			join m2m_blog_category mbc on b.id = mbc.blog_id
		where b.is_active is true 
	`

	whereStmtBuilder := strings.Builder{}
	phCounter := 1

	if len(filter.CategoriesIDs) > 0 {
		whereStmtBuilder.WriteString(fmt.Sprintf(`and mbc.category_id = any($%d::bigint[])`, phCounter))
		phValues = append(phValues, pq.Int64Array(filter.CategoriesIDs))
		phCounter++
	}

	// возвращаем для первой страницы
	return fmt.Sprintf(`
		%s
		%s
		order by b.ordering_number
    `, defaultSql, whereStmtBuilder.String()), phValues
}
