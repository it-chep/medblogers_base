package update_draft_blog

import (
	"context"
	"html"
	"medblogers_base/internal/modules/admin/entities/blog/action/update_draft_blog/dal"
	"medblogers_base/internal/modules/admin/entities/blog/action/update_draft_blog/dto"
	"medblogers_base/internal/pkg/postgres"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	xhtml "golang.org/x/net/html"
)

// Action .
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do обновление статьи
func (a *Action) Do(ctx context.Context, blogID uuid.UUID, req dto.Request) error {
	blog, err := a.dal.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	if blog.IsActive.Bool {
		return errors.New("Статью надо сначала снять с публикации чтобы поменять")
	}

	req.SearchText = extractSearchText(req.Body)

	return a.dal.UpdateBlog(ctx, blogID, req)
}

func extractSearchText(body string) string {
	doc, err := xhtml.Parse(strings.NewReader(body))
	if err != nil {
		return ""
	}

	var parts []string
	var walk func(*xhtml.Node)

	walk = func(node *xhtml.Node) {
		if node.Type == xhtml.ElementNode {
			switch node.Data {
			case "script", "style":
				return
			}
		}

		if node.Type == xhtml.TextNode {
			text := strings.TrimSpace(html.UnescapeString(node.Data))
			if text != "" {
				parts = append(parts, text)
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			walk(child)
		}
	}

	walk(doc)

	return strings.Join(parts, " ")
}
