package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/R894/bugtracker/internal/domain/project/service"
	"github.com/R894/bugtracker/internal/web/middleware"
	"github.com/R894/bugtracker/internal/web/response"

	"github.com/go-chi/chi/v5"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ProjectHandler struct {
	service  *service.ProjectService
	validate *validator.Validate
	trans    ut.Translator
}

type newProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p *ProjectHandler) CreateNewProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := r.Context().Value(middleware.ContextKeyUsername).(string)
	var req newProjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w)
		return
	}

	project, err := p.service.CreateProject(ctx, req.Name, req.Description, username)
	if err != nil {
		log.Println("Error creating Project: ", err)
		response.InternalError(w)
		return
	}

	response.JSON(w, http.StatusOK, project)
}

func (p *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projects, err := p.service.GetProjects(ctx)
	if err != nil {
		log.Println("Error getting Projects: ", err)
		response.InternalError(w)
		return
	}

	response.JSON(w, http.StatusOK, projects)
}

func (p *ProjectHandler) GetProjectsByOwnerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "userId")

	projects, err := p.service.GetProjectsByOwnerId(ctx, id)
	if err != nil {
		log.Println("Error getting Projects: ", err)
		response.InternalError(w)
		return
	}

	response.JSON(w, http.StatusOK, projects)
}
