package v1

import (
	"medblogers_base/internal/app/middleware"
	"medblogers_base/internal/pkg/config"

	"github.com/go-chi/chi/v5"
	base_middleware "github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	mutableConfig config.Config

	Router *chi.Mux
}

func NewRouter(mutableConfig config.Config) *Router {
	r := &Router{
		mutableConfig: mutableConfig,
		Router:        chi.NewRouter(),
	}

	r.setupMiddlewares()

	return r
}

func (r *Router) setupMiddlewares() {
	r.Router.Use(middleware.MetadataMiddleware)
	r.Router.Use(base_middleware.Recoverer)
	r.Router.Use(middleware.CORSMiddleware)
	//s.router.Use(middleware.CSRFMiddleware)
	r.Router.Use(middleware.ConfigMiddleware(r.mutableConfig))
	r.Router.Use(middleware.LoggerMiddleware)
	r.Router.Use(middleware.RateLimitMiddleware)
	r.Router.Use(middleware.ResponseTimeMiddleware)
}
