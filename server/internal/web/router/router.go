package router

import (
	"bugtracker/internal/database"
	"bugtracker/internal/web/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	db, err := database.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	bugHandler := handler.NewBugHandler(db)
	userHandler := handler.NewUserHandler(db)

	r.Route("/bugs", func(r chi.Router) {
		r.Get("/", bugHandler.GetBugs)
		r.Post("/", bugHandler.CreateNewBug)
		r.Get("/{bugId}", bugHandler.GetBugByID)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.CreateNewUser)
		r.Post("/", userHandler.CreateNewUser)
		r.Get("/{userId}", userHandler.GetUserById)
	})

	return r
}
