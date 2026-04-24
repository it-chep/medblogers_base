package save_site_form_answer

import (
	"context"
	"medblogers_base/internal/modules/analytics/actions/save_site_form_answer/dal"
	"medblogers_base/internal/modules/analytics/actions/save_site_form_answer/dto"
	"medblogers_base/internal/pkg/postgres"
	"strings"

	"github.com/google/uuid"
)

// Action сохраняет ответ формы сайта.
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
func (a *Action) Do(ctx context.Context, req dto.SaveSiteFormAnswerRequest) error {
	if strings.TrimSpace(req.FormName) == "" || len(req.Answer) == 0 || req.CookieID == "" {
		return nil
	}

	parsedCookieID, err := uuid.Parse(req.CookieID)
	if err != nil {
		return nil
	}

	return a.dal.CreateSiteFormAnswer(ctx, dto.CreateSiteFormAnswerRequest{
		FormName: strings.TrimSpace(req.FormName),
		Answer:   req.Answer,
		CookieID: parsedCookieID,
		Source:   strings.TrimSpace(req.Source),
		TG:       req.TG,
	})
}
