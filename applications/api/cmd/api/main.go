package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/cirodourado/portfolio/applications/api/internal/database"
	"github.com/cirodourado/portfolio/applications/api/internal/handlers"
	"github.com/cirodourado/portfolio/applications/api/internal/repositories/project"
)

func main() {
    cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

    projectRepo := repositories.NewGormProjectRepository(db)
    projectHandler := handlers.NewProjectHandler(projectRepo)
	experienceHandler := handlers.NewExperienceHandler(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            next.ServeHTTP(w, r)
        })
    })

    r.Get("/health", handlers.HealthCheck)
    r.Route("/api/v1", func(r chi.Router) {
        r.Get("/projects", projectHandler.GetAll)
        r.Get("/experiences", experienceHandler.GetAll)
    })

    srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
	    log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}