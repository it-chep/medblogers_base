package app

import (
	configMock "medblogers_base/internal/config/mocks"
	"medblogers_base/internal/modules/doctors"
	pkgHttpMocks "medblogers_base/internal/pkg/http/mocks"
	"medblogers_base/internal/pkg/postgres"

	"github.com/golang/mock/gomock"
)

type TestableModule struct {
	Module      *doctors.Module
	PoolWrapper postgres.PoolWrapper
	Http        *pkgHttpMocks.MockExecutor
	MockCtrl    *gomock.Controller
	ConfigMock  *configMock.MockAppConfig
}
