package v1

import (
	"context"
	"net/url"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"medblogers_base/internal/modules/analytics/actions/save_analytics/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/analytics/v1"
)

// SaveAnalytics сохраняет аналитику сайта.
func (i *Implementation) SaveAnalytics(ctx context.Context, req *desc.SaveAnalyticsRequest) (*desc.SaveAnalyticsResponse, error) {
	if !isAllowedDomain(req.GetDomainName(), i.config.GetAllowedHosts()) {
		return nil, status.Error(codes.PermissionDenied, "domain is not allowed")
	}

	err := i.analytics.Actions.SaveAnalytics.Do(ctx, dto.SaveAnalyticsRequest{
		UtmSource:   req.GetUtmSource(),
		UtmMedium:   req.GetUtmMedium(),
		UtmCampaign: req.GetUtmCampaign(),
		UtmTerm:     req.GetUtmTerm(),
		UtmContent:  req.GetUtmContent(),
		DomainName:  req.GetDomainName(),
		CookieID:    req.GetCookieId(),
		Company:     req.Company,
		Event:       req.Event,
	})
	if err != nil {
		return nil, nil
	}

	return &desc.SaveAnalyticsResponse{}, nil
}

func isAllowedDomain(domain string, allowedHosts []string) bool {
	normalized := normalizeDomain(domain)
	if normalized == "" {
		return false
	}

	for _, allowedHost := range allowedHosts {
		if normalized == normalizeDomain(allowedHost) {
			return true
		}
	}

	return false
}

func normalizeDomain(value string) string {
	trimmed := strings.ToLower(strings.TrimSpace(value))
	if trimmed == "" {
		return ""
	}

	if parsed, err := url.Parse(trimmed); err == nil {
		if parsed.Hostname() != "" {
			return strings.ToLower(parsed.Hostname())
		}
	}

	if parsed, err := url.Parse("https://" + trimmed); err == nil && parsed.Hostname() != "" {
		return strings.ToLower(parsed.Hostname())
	}

	return trimmed
}
