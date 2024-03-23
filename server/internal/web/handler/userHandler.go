package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/R894/bugtracker/internal/domain/user/aggregate"
	"github.com/R894/bugtracker/internal/domain/user/service"
	"github.com/R894/bugtracker/internal/web/middleware"
	"github.com/R894/bugtracker/internal/web/response"

	"github.com/go-chi/chi/v5"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
	trans    ut.Translator
}
type userLoginResponse struct {
	Token string          `json:"token"`
	User  *aggregate.User `json:"user"`
}

func (u *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createUserReq aggregate.UserRegisterModel
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		response.BadRequest(w)
		return
	}

	err = u.validate.Struct(createUserReq)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		response.ValidationErrs(w, errors, u.trans)
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
	response.JSON(w, http.StatusOK, userLoginResponse{Token: token, User: user})
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
	id := chi.URLParam(r, "userId")
	ctx := r.Context()

	user, err := u.service.GetUserByID(ctx, id)
	if err != nil {
		log.Println(err)
		response.NotFound(w)
		return
	}

	response.JSON(w, http.StatusOK, user)
}
