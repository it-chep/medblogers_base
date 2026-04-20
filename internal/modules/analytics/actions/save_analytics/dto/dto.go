package dto

import "github.com/google/uuid"

type SaveAnalyticsRequest struct {
	UtmSource   string
	UtmMedium   string
	UtmCampaign string
	UtmTerm     string
	UtmContent  string
	DomainName  string
	CookieID    string
	Company     *string
	Event       *string
}

type CreateAnalyticsRequest struct {
	ID          uuid.UUID
	UtmSource   string
	UtmMedium   string
	UtmCampaign string
	UtmTerm     string
	UtmContent  string
	DomainName  string
	CookieID    uuid.UUID
	Company     *string
	Event       *string
}
