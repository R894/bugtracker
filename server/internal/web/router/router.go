package router

import (
	"bugtracker/internal/web/handler"
	"database/sql"

	mw "bugtracker/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter(db *sql.DB) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	appHandler := handler.New(db)

	r.Route("/bugs", func(r chi.Router) {
		mw.ApplyAuthMiddleware(r)
		r.Get("/", appHandler.Bug.GetBugs)
		r.Post("/", appHandler.Bug.CreateNewBug)
		r.Get("/{bugId}", appHandler.Bug.GetBugByID)
		r.Get("/projects/{projectId}", appHandler.Bug.GetBugsByProjectID)
	})

	r.Route("/comments", func(r chi.Router) {
		mw.ApplyAuthMiddleware(r)
		r.Post("/", appHandler.Comment.CreateNewComment)
		r.Put("/", appHandler.Comment.UpdateComment)
		r.Delete("/", appHandler.Comment.DeleteComment)
	})

	r.Route("/projects", func(r chi.Router) {
		mw.ApplyAuthMiddleware(r)
		r.Post("/", appHandler.Project.CreateNewProject)
		r.Get("/{userId}", appHandler.Project.GetProjectsByOwnerId)
	})

	r.Route("/users", func(r chi.Router) {
		r.With(mw.AuthMiddleware).Get("/", appHandler.User.GetUsers)
		r.Post("/", appHandler.User.CreateNewUser)
		r.Post("/login", appHandler.User.UserLogin)
		r.With(mw.AuthMiddleware).Get("/{userId}", appHandler.User.GetUserById)
	})

	return r, nil
}
