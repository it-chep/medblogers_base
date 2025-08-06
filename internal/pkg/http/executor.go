package http

import "net/http"

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Executor

// Executor ...
type Executor interface {
	Do(req *http.Request) (*http.Response, error)
}
