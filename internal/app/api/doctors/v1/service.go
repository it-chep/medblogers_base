package v1

import (
	"medblogers_base/internal/app/middleware"
	"medblogers_base/internal/modules/doctors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	doctors *doctors.Module
	router  *chi.Mux
}

func NewService(doctors *doctors.Module) *Service {
	s := &Service{
		doctors: doctors,
		router:  chi.NewRouter(),
	}
	s.setupMiddlewares()
	s.setupRoutes()

	return s
}

func (s *Service) setupRoutes() {
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/settings", s.Settings) // GET /api/v1/settings
		r.Get("/cities_list", s.AllCities)
		r.Get("/specialities_list", s.AllSpecialities)

		r.Route("/doctors", func(r chi.Router) {
			r.Get("/search", s.Search) // /api/v1/doctors/search
			r.Get("/filter", s.Filter) // filter

			r.Get("/{doctor_id}", s.DoctorDetail)  // Обрабатывает /api/v1/doctors/23
			r.Get("/{doctor_id}/", s.DoctorDetail) // Обрабатывает /api/v1/doctors/23/

			//r.Get("/") // дефолт список (мб заменить на фильтр)
		})

		//r.Post("/create_new_doctor" )

	})
}

func (s *Service) setupMiddlewares() {
	s.router.Use(middleware.CORSMiddleware)
	//s.router.Use(middleware.CSRFMiddleware)
	s.router.Use(middleware.LoggerMiddleware)
	s.router.Use(middleware.RateLimitMiddleware)
	s.router.Use(middleware.ResponseTimeMiddleware)
}

func (s *Service) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
