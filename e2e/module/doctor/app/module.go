package app

import (
	"medblogers_base/internal/modules/doctors"
	"medblogers_base/internal/pkg/postgres"

	"github.com/golang/mock/gomock"
)

type TestableModule struct {
	Module      *doctors.Module
	PoolWrapper *postgres.PoolWrapper
	Http        *mocks.MockHttpClient
	MockCtrl    *gomock.Controller
}
