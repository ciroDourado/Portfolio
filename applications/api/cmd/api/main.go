package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/cirodourado/portfolio/applications/api/internal/database"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

func main() {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/projects", func(w http.ResponseWriter, r *http.Request) {
		var projects []models.Project

		result := db.Find(&projects)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	})

	http.ListenAndServe(":8080", r)
}