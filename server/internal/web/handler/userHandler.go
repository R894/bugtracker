package handler

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/user/aggregate"
	"bugtracker/internal/domain/user/service"
	"bugtracker/internal/web/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(db *database.PostgresDB) *UserHandler {
	return &UserHandler{
		service: service.NewUserService(db),
	}
}

func (u *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createUserReq aggregate.UserRegisterModel
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		response.BadRequest(w)
		return
	}

	u.service.CreateUser(ctx, createUserReq)
	response.OK(w)
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	u.service.GetUserByID(ctx, id)
	response.OK(w)
}
