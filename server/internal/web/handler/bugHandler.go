package handler

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/bug/aggregate"
	"bugtracker/internal/domain/bug/service"
	"bugtracker/internal/web/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BugHandler struct {
	service *service.BugService
}

func NewBugHandler(db *database.PostgresDB) *BugHandler {
	return &BugHandler{
		service: service.NewBugService(db),
	}
}

func (b *BugHandler) CreateNewBug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createBugReq aggregate.CreateBugRequest
	err := json.NewDecoder(r.Body).Decode(&createBugReq)
	if err != nil {
		response.BadRequest(w)
		return
	}

	b.service.CreateBug(ctx, createBugReq)
	response.Created(w)
}

func (b *BugHandler) GetBugByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	bug, err := b.service.GetBugByID(ctx, id)
	if err != nil {
		response.BadRequest(w)
		return
	}
	response.JSON(w, http.StatusOK, bug)
}

func (b *BugHandler) GetBugs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	b.service.GetBugs(ctx)
}

func (b *BugHandler) UpdateBug(w http.ResponseWriter, r *http.Request) {
	var updateBugRequest aggregate.UpdateBugRequest
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&updateBugRequest)
	if err != nil {
		response.BadRequest(w)
		return
	}
	b.service.UpdateBug(ctx, &updateBugRequest)
	response.OK(w)
}
