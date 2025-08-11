package v1

import (
	"medblogers_base/internal/app/middleware"
	"medblogers_base/internal/config"
	pkgConfig "medblogers_base/internal/pkg/config"
	"medblogers_base/internal/pkg/logger"

	"github.com/go-chi/chi/v5"
	base_middleware "github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	mutableConfig pkgConfig.Config
	staticConfig  config.AppConfig
	Router        *chi.Mux
}

func NewRouter(
	staticConfig config.AppConfig,
	mutableConfig pkgConfig.Config,
) *Router {
	r := &Router{
		mutableConfig: mutableConfig,
		Router:        chi.NewRouter(),
		staticConfig:  staticConfig,
	}

	r.setupMiddlewares()

	return r
}

func (r *Router) setupMiddlewares() {
	r.Router.Use(base_middleware.Recoverer)
	//r.Router.Use(middleware.CORSMiddleware(r.staticConfig)) // todo вернуть
	//s.router.Use(middleware.CSRFMiddleware)
	r.Router.Use(middleware.ConfigMiddleware(r.mutableConfig))
	r.Router.Use(middleware.LoggerMiddleware(logger.New()))
	r.Router.Use(middleware.RateLimitMiddleware)
	r.Router.Use(middleware.ResponseTimeMiddleware)
}
