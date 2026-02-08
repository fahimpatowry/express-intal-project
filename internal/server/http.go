package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)                    // logs requests
	r.Use(middleware.Recoverer)                 // recovers from panics
	r.Use(middleware.Timeout(10 * time.Second)) // timeout

	return &Server{
		Router: r,
	}
}

// Add routes here
func (s *Server) Routes() {
	s.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}

func (s *Server) Start(port string) error {
	fmt.Println("Server running on port", port)
	return http.ListenAndServe(":"+port, s.Router)
}
