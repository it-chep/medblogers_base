package dto

import "github.com/google/uuid"

type SaveSiteFormAnswerRequest struct {
	FormName string
	Answer   []byte
	CookieID string
	Source   string
	TG       *string
}

type CreateSiteFormAnswerRequest struct {
	ID       int64
	FormName string
	Answer   []byte
	CookieID uuid.UUID
	Source   string
	TG       *string
}
