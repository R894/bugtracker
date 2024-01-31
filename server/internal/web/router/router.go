package router

import (
	"bugtracker/internal/database"
	"bugtracker/internal/web/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(db *database.PostgresDB) (*chi.Mux, error) {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	bugHandler := handler.NewBugHandler(db)
	userHandler := handler.NewUserHandler(db)

	r.Route("/bugs", func(r chi.Router) {
		r.Get("/", bugHandler.GetBugs)
		r.Post("/", bugHandler.CreateNewBug)
		r.Get("/{bugId}", bugHandler.GetBugByID)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateNewUser)
		r.Post("/login", userHandler.UserLogin)
		r.Get("/{userId}", userHandler.GetUserById)
	})

	return r, nil
}
