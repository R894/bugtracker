package handler

import (
	"bugtracker/internal/domain/project/service"
	"bugtracker/internal/web/middleware"
	"bugtracker/internal/web/response"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type ProjectHandler struct {
	service *service.ProjectService
}

type newProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewProjectHandler(db *sql.DB) *ProjectHandler {
	return &ProjectHandler{
		service: service.NewProjectService(db),
	}
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
