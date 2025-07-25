package v1

import (
	"medblogers_base/internal/app/middleware"
	"medblogers_base/internal/modules/doctors"
	"medblogers_base/internal/pkg/config"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	mutableConfig config.Config

	doctors *doctors.Module
	router  *chi.Mux
}

func NewService(doctors *doctors.Module, mutableConfig config.Config) *Service {
	s := &Service{
		doctors:       doctors,
		mutableConfig: mutableConfig,
		router:        chi.NewRouter(),
	}

	s.setupMiddlewares()
	s.setupRoutes()

	return s
}

func (s *Service) setupRoutes() {
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/settings", s.Settings)          // GET /api/v1/settings
		r.Get("/counters_info", s.CountersInfo) // GET /api/v1/counters_info

		r.Get("/cities_list", s.AllCities)             // GET /api/v1/cities_list ДЛЯ РЕГИ
		r.Get("/specialities_list", s.AllSpecialities) // GET /api/v1/specialities_list ДЛЯ РЕГИ

		r.Route("/doctors", func(r chi.Router) {
			r.Get("/search", s.Search)        // GET /api/v1/doctors/search
			r.Get("/filter", s.Filter)        // GET /api/v1/doctors/filter
			r.Post("/create", s.CreateDoctor) // POST /api/v1/doctors/create

			r.Get("/{doctor_id}", s.DoctorDetail)  // Обрабатывает /api/v1/doctors/23
			r.Get("/{doctor_id}/", s.DoctorDetail) // Обрабатывает /api/v1/doctors/23/
		})

	})
}

func (s *Service) setupMiddlewares() {
	s.router.Use(middleware.CORSMiddleware)
	//s.router.Use(middleware.CSRFMiddleware)
	s.router.Use(middleware.ConfigMiddleware(s.mutableConfig))
	s.router.Use(middleware.LoggerMiddleware)
	s.router.Use(middleware.RateLimitMiddleware)
	s.router.Use(middleware.ResponseTimeMiddleware)
}

func (s *Service) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
