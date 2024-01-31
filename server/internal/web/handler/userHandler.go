package handler

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/user/aggregate"
	"bugtracker/internal/domain/user/service"
	"bugtracker/internal/web/middleware"
	"bugtracker/internal/web/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var loginUserReq aggregate.UserLoginModel
	err := json.NewDecoder(r.Body).Decode(&loginUserReq)
	if err != nil {
		response.BadRequest(w)
		return
	}

	user, err := u.service.GetUserByEmail(ctx, loginUserReq.Email)
	if err != nil {
		response.NotFound(w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserReq.Password))
	if err != nil {
		response.Unauthorized(w)
		return
	}

	token, err := middleware.CreateToken(user.Username)
	if err != nil {
		response.InternalError(w)
	}
	response.JSON(w, http.StatusOK, token)
}

func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := u.service.GetUsers(ctx)

	if err != nil {
		response.InternalError(w)
		return
	}
	response.JSON(w, http.StatusOK, users)
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	u.service.GetUserByID(ctx, id)
	response.OK(w)
}
