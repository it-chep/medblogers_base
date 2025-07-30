package app

import (
	"medblogers_base/e2e/module/doctor/app/mocks"
	"medblogers_base/internal/modules/doctors"
	"medblogers_base/internal/pkg/postgres"
	"net/http"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . HttpClient

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type TestableModule struct {
	Module      *doctors.Module
	PoolWrapper *postgres.PoolWrapper
	Http        *mocks.MockHttpClient
	MockCtrl    *gomock.Controller
}
