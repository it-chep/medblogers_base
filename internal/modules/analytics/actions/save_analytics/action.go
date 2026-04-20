package save_analytics

import (
	"context"
	"medblogers_base/internal/modules/analytics/actions/save_analytics/dal"
	"medblogers_base/internal/modules/analytics/actions/save_analytics/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

// Action сохраняет событие аналитики.
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
func (a *Action) Do(ctx context.Context, req dto.SaveAnalyticsRequest) error {
	if req.CookieID == "" {
		return nil
	}

	parsedCookieID, err := uuid.Parse(req.CookieID)
	if err != nil {
		return nil
	}

	isCookieExists, err := a.dal.IsCookieUserExists(ctx, parsedCookieID)
	if err != nil {
		return err
	}
	if !isCookieExists {
		return nil
	}

	isAnalyticsExists, err := a.dal.IsAnalyticsExistsForLast7Days(ctx, dto.CreateAnalyticsRequest{
		UtmSource:   req.UtmSource,
		UtmMedium:   req.UtmMedium,
		UtmCampaign: req.UtmCampaign,
		UtmTerm:     req.UtmTerm,
		UtmContent:  req.UtmContent,
		DomainName:  req.DomainName,
		CookieID:    parsedCookieID,
		Company:     req.Company,
		Event:       req.Event,
	})
	if err != nil {
		return err
	}
	if isAnalyticsExists {
		return nil
	}

	analID, _ := uuid.NewV7()
	return a.dal.CreateAnalytics(ctx, dto.CreateAnalyticsRequest{
		ID:          analID,
		UtmSource:   req.UtmSource,
		UtmMedium:   req.UtmMedium,
		UtmCampaign: req.UtmCampaign,
		UtmTerm:     req.UtmTerm,
		UtmContent:  req.UtmContent,
		DomainName:  req.DomainName,
		CookieID:    parsedCookieID,
		Company:     req.Company,
		Event:       req.Event,
	})
}
