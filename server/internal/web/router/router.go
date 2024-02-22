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

	bugHandler := handler.NewBugHandler(db)
	userHandler := handler.NewUserHandler(db)
	projectHandler := handler.NewProjectHandler(db)
	commentHandler := handler.NewCommentHandler(db)

	r.Route("/bugs", func(r chi.Router) {
		mw.ApplyAuthMiddleware(r)
		r.Get("/", bugHandler.GetBugs)
		r.Post("/", bugHandler.CreateNewBug)
		r.Get("/{bugId}", bugHandler.GetBugByID)
		r.Get("/projects/{projectId}", bugHandler.GetBugsByProjectID)
	})

	r.Route("/comments", func(r chi.Router) {
		mw.ApplyAuthMiddleware(r)
		r.Post("/", commentHandler.CreateNewComment)
		r.Put("/", commentHandler.UpdateComment)
		r.Delete("/", commentHandler.DeleteComment)
	})

	r.Route("/projects", func(r chi.Router) {
		mw.ApplyAuthMiddleware(r)
		r.Post("/", projectHandler.CreateNewProject)
	})

	r.Route("/users", func(r chi.Router) {
		r.With(mw.AuthMiddleware).Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateNewUser)
		r.Post("/login", userHandler.UserLogin)
		r.With(mw.AuthMiddleware).Get("/{userId}", userHandler.GetUserById)
	})

	return r, nil
}
