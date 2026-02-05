package main

import (
    "fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hibiken/asynq"
	"github.com/cirodourado/portfolio/applications/api/internal/config"
	"github.com/cirodourado/portfolio/applications/api/internal/database"
	"github.com/cirodourado/portfolio/applications/api/internal/handlers"
	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

func main() {
    cfg, err := config.LoadApiConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	db, err := database.NewConnection(cfg.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

    queueClient := asynq.NewClient(
        asynq.RedisClientOpt{
            Addr: cfg.GetRedisURL(),
        },
    )
    defer queueClient.Close()

    projectRepo := repositories.NewGormProjectRepository(db)
    projectHandler := handlers.NewProjectHandler(projectRepo)
    experienceRepo := repositories.NewGormExperienceRepository(db)
	experienceHandler := handlers.NewExperienceHandler(experienceRepo)
	profileRepo := repositories.NewGormProfileRepository(db)
	profileHandler := handlers.NewProfileHandler(profileRepo, cfg.StorageURL)
	contactHandler := handlers.NewContactHandler(queueClient)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            next.ServeHTTP(w, r)
        })
    })
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   cfg.AllowedOrigins,
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true,
    }))

    r.Get("/health", handlers.HealthCheck)
    r.Route("/v1", func(r chi.Router) {
        r.Get("/projects", projectHandler.GetAll)
        r.Get("/projects/{slug}", projectHandler.GetBySlug)
        r.Get("/experiences", experienceHandler.GetAll)
        r.Get("/profile", profileHandler.GetProfile)
        r.Post("/contact", contactHandler.HandleSubmit)
    })

    srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.AppPort),
		Handler: r,
	}

	srv.ListenAndServe()
}