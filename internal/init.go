package internal

import (
	"context"
	"github.com/it-chep/medblogers_base/internal/config"
	"github.com/it-chep/medblogers_base/internal/pkg/postgres"
	"net/http"
	"time"

	moduleadmin "github.com/it-chep/medblogers_base/internal/modules/admin"
	moduledoctors "github.com/it-chep/medblogers_base/internal/modules/doctors"
)

func (a *App) initPostgres(ctx context.Context) *App {
	postgres.NewPoolWrapper()
	// todo gracefull
	//	a.postgresConn.Close()
	//

	return a
}

func (a *App) initConfig(_ context.Context) *App {
	a.config = config.NewConfig()
	return a
}

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		admin:   moduleadmin.New(),
		doctors: moduledoctors.New(),
	}

	return a
}

func (a *App) initCache(_ context.Context) *App {
	return a
}

func (a *App) initServer(_ context.Context) *App {

	a.server = &http.Server{
		Addr:         a.config.HTTPServer.Address,
		Handler:      ,// a.controller.telegramWebhookController,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}
}
