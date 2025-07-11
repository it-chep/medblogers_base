package salebot

import "net/http"

type Gateway struct {
	client *http.Client
}

// NewGateway - конструктор
func NewGateway(client *http.Client) *Gateway {
	return &Gateway{
		client: client,
	}
}
