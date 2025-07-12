package salebot

import "net/http"

// HTTPClient ...
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Gateway в сервис нотификации
type Gateway struct {
	client HTTPClient
}

// NewGateway - конструктор
func NewGateway(client HTTPClient) *Gateway {
	return &Gateway{
		client: client,
	}
}
