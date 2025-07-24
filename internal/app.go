package internal

import (
	"context"
	"fmt"
	v1 "medblogers_base/internal/app/api/doctors/v1"
	"medblogers_base/internal/config"
	pkgConfig "medblogers_base/internal/pkg/config"

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
	postgres postgres.PoolWrapper

	modules modules

	// http сервер
	controllers controllers
	server      *http.Server

	// конфиги
	config        *config.Config
	mutableConfig pkgConfig.Config

	// периодические задачи
	//tasks
	//worker_pool
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initConfig(ctx).
		initPostgres(ctx).
		initMutableConfig(ctx).
		initModules(ctx).
		initControllers(ctx).
		initServer(ctx)

	return a
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("application recovered from panic")
		}
	}()

	fmt.Printf("[APP] Запуск приложения, подключение http://localhost%s \n", a.config.Server.Address)
	//a.workerPool.Run(ctx)

	if err := a.server.ListenAndServe(); err != nil {
		fmt.Printf("[APP] Не удалось запустить приложение: %e", err)
	}
}
