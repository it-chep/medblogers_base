package fixture

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	"medblogers_base/e2e/module/doctor/app"
	"medblogers_base/internal/config"
	configMock "medblogers_base/internal/config/mocks"
	moduleobject "medblogers_base/internal/modules/doctors"
	pkgHttp "medblogers_base/internal/pkg/http"
	pkgHttpMocks "medblogers_base/internal/pkg/http/mocks"
	"medblogers_base/internal/pkg/postgres"
)

var SetupModule = func(ctx context.Context, pool postgres.PoolWrapper) *app.TestableModule {
	By("Создание модуля")

	mockCtrl := gomock.NewController(GinkgoT())
	mockHttp := pkgHttpMocks.NewMockExecutor(mockCtrl)
	configMock := configMock.NewMockAppConfig(gomock.NewController(GinkgoT()))

	// notification
	configMock.EXPECT().GetSalebotHost().Return("localhost").AnyTimes()
	configMock.EXPECT().GetCreateNotificationChatID().Return(int64(1234)).AnyTimes()

	// subscribers
	configMock.EXPECT().GetSubscribersHost().Return("localhost").AnyTimes()
	configMock.EXPECT().GetSubscribersPort().Return("8080").AnyTimes()

	// s3
	configMock.EXPECT().GetUserPhotosBucket().Return("bucket").AnyTimes()
	configMock.EXPECT().GetS3Config().Return(config.S3Config{}).AnyTimes()

	httpCons := map[string]pkgHttp.Executor{
		config.Subscribers: mockHttp,
		config.Salebot:     mockHttp,
	}

	return &app.TestableModule{
		Http:        mockHttp,
		Module:      moduleobject.New(httpCons, configMock, pool),
		MockCtrl:    mockCtrl,
		PoolWrapper: pool,
		ConfigMock:  configMock,
	}
}
