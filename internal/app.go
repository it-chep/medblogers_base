package internal

import (
	"context"
	"go.uber.org/zap"
	v1 "medblogers_base/internal/app/api/doctors/v1"
	"medblogers_base/internal/config"
	"medblogers_base/internal/pkg/postgres"
	"net/http"

	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/doctors"
)

type modules struct {
	admin   *admin.Module
	doctors *doctors.Module
}

type controllers struct {
	restController *v1.Service
}

type App struct {
	logger   *zap.Logger
	postgres postgres.PoolWrapper

	modules modules

	// http сервер
	controllers controllers
	server      *http.Server

	config *config.Config

	// периодические задачи
	//tasks
	//worker_pool
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initConfig(ctx).
		initLogger(ctx).
		initPostgres(ctx).
		initModules(ctx).
		initControllers(ctx).
		initServer(ctx)

	return a
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Error("application recovered from panic")
		}
	}()

	a.logger.Info("[APP] Запуск приложения")
	//a.workerPool.Run(ctx)

	if err := a.server.ListenAndServe(); err != nil {
		a.logger.Fatal("[APP] Не удалось запустить приложение: %s", zap.Error(err))
	}
}
