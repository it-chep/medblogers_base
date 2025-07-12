package internal

import (
	"context"
	"github.com/it-chep/medblogers_base/internal/config"
	"log/slog"
	"net/http"

	"github.com/it-chep/medblogers_base/internal/modules/admin"
	"github.com/it-chep/medblogers_base/internal/modules/doctors"
)

type modules struct {
	admin   *admin.Module
	doctors *doctors.Module
}

type App struct {
	// logger  *slog.Logger

	// postgres

	modules modules

	// http сервер
	// controllers
	server *http.Server

	config *config.Config

	// периодические задачи
	//tasks
	//worker_pool
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initConfig(ctx).
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
			a.logger.Error("application recovered from panic", slog.Any("error", r))
		}
	}()

	a.logger.Info("Запуск приложения")
	a.workerPool.Run(ctx)

	if err := a.server.ListenAndServe(); err != nil {
		a.logger.Fatalf(ctx, "Не удалось запустить приложение: %s", err)
	}
}
