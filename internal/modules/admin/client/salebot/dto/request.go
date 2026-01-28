package dto

const (
	ErrorEvent = "jopa_medblogers_error"
	MMEvent    = "push_mm"
)

type ErrorRequest struct {
	Message  string `json:"message"`
	Error    string `json:"error_text"`
	ClientID int64  `json:"client_id"`
}

type MMRequest struct {
	Message  string `json:"message"`
	ClientID int64  `json:"client_id"`
}
