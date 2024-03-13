package main

import (
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	s := NewServer()
	s.MountHandlers()

	log.Fatal(http.ListenAndServe(":3000", s.Router))
}

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

// MountHandlers mounts middlewares and handles on router
func (s *Server) MountHandlers() {
	// Middlewares
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// Routes
	s.Router.Mount("/pprof", middleware.Profiler())
	s.Router.Get("/health", HealthCheck)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
