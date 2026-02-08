package server

import (
	"fmt"
	"net/http"
	"time"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// health "initial_project/internal/handlers/healthCheck"
	health "initial_project/internal/handlers"

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
	s.Router.Get("/health", health.HealthCheck)

	// Swagger route with dynamic URL
	port := os.Getenv("PORT")
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)

	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL), // <- dynamically sets swagger.json URL
	))
}

func (s *Server) Start(port string) error {
	fmt.Println("Server running on port", port)
	return http.ListenAndServe(":"+port, s.Router)
}
