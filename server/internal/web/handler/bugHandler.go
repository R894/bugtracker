package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/R894/bugtracker/internal/domain/bug/core/domain/aggregate"
	"github.com/R894/bugtracker/internal/domain/bug/core/service"
	"github.com/R894/bugtracker/internal/web/response"

	"github.com/go-chi/chi/v5"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type BugHandler struct {
	service  *service.BugService
	validate *validator.Validate
	trans    ut.Translator
}

func (b *BugHandler) CreateNewBug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createBugReq aggregate.CreateBugRequest
	err := json.NewDecoder(r.Body).Decode(&createBugReq)
	if err != nil {
		response.BadRequest(w)
		return
	}

	bug, err := b.service.CreateBug(ctx, createBugReq)
	if err != nil {
		log.Println(err)
		response.InternalError(w)
		return
	}
	response.JSON(w, http.StatusCreated, bug)
}

func (b *BugHandler) GetBugByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bugId")
	ctx := r.Context()
	bug, err := b.service.GetBugByID(ctx, id)
	if err != nil {
		log.Println("Error getting Bug by ID:", err)
		response.BadRequest(w)
		return
	}
	response.JSON(w, http.StatusOK, bug)
}

func (b *BugHandler) GetBugsByProjectID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "projectId")
	ctx := r.Context()
	bugs, err := b.service.GetBugsByProjectId(ctx, id)
	if err != nil {
		log.Println("Error getting Bugs by project ID:", err)
		response.BadRequest(w)
		return
	}
	response.JSON(w, http.StatusOK, bugs)
}

func (b *BugHandler) GetBugs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bugs, err := b.service.GetBugs(ctx)
	if err != nil {
		response.BadRequest(w)
		return
	}
	response.JSON(w, http.StatusOK, bugs)
}

func (b *BugHandler) UpdateBug(w http.ResponseWriter, r *http.Request) {
	var updateBugRequest aggregate.UpdateBugRequest
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&updateBugRequest)
	if err != nil {
		response.BadRequest(w)
		return
	}
	err = b.service.UpdateBug(ctx, &updateBugRequest)
	if err != nil {
		response.InternalError(w)
	}
	response.OK(w)
}

func (b *BugHandler) AssignUser(w http.ResponseWriter, r *http.Request) {
	type AssignUserRequest struct {
		BugId    string `json:"bugId"`
		Username string `json:"userName"`
	}

	var req AssignUserRequest
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w)
		return
	}

	err = b.service.AssignBugTo(ctx, req.BugId, req.Username)
	if err != nil {
		response.InternalError(w)
	}
	response.OK(w)
}
