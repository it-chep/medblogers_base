package app

import (
	"github.com/golang/mock/gomock"
	"medblogers_base/internal/modules/doctors"
	"medblogers_base/internal/pkg/postgres"
)

type TestableModule struct {
	Module      *doctors.Module
	PoolWrapper *postgres.PoolWrapper
	Http        *mocks.MockHttpClient
	MockCtrl    *gomock.Controller
}
