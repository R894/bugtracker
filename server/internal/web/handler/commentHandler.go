package handler

import (
	"bugtracker/internal/domain/comment/service"
	"bugtracker/internal/web/middleware"
	"bugtracker/internal/web/response"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type CommentHandler struct {
	service  *service.CommentService
	validate *validator.Validate
	trans    ut.Translator
}

type newCommentRequest struct {
	BugId   string `json:"bugId"`
	Content string `json:"content"`
}

type updateCommentRequest struct {
	CommentId string `json:"commentId"`
	Content   string `json:"content"`
}

type deleteCommentRequest struct {
	CommentId string `json:"commentId"`
}

func (c *CommentHandler) CreateNewComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username, err := middleware.GetContextUsername(ctx)
	if err != nil {
		response.Unauthorized(w)
		return
	}

	var req newCommentRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w)
		return
	}

	err = c.service.NewComment(ctx, req.BugId, req.Content, username)
	if err != nil {
		log.Println("Error creating comment: ", err)
		response.InternalError(w)
		return
	}

	response.OK(w)
}

func (c *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req updateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w)
		return
	}

	err = c.service.UpdateComment(ctx, req.CommentId, req.Content)
	if err != nil {
		log.Println("Error updating comment: ", err)
		response.InternalError(w)
		return
	}

	response.OK(w)
}

func (c *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req deleteCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w)
		return
	}

	err = c.service.DeleteComment(ctx, req.CommentId)
	if err != nil {
		log.Println("Error deleting comment: ", err)
		response.InternalError(w)
		return
	}

	response.OK(w)
}

func (c *CommentHandler) GetCommentsByBugId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bugId := chi.URLParam(r, "bugId")

	comments, err := c.service.GetByBugId(ctx, bugId)
	if err != nil {
		log.Println("Error getting comments by bug ID: ", err)
		response.InternalError(w)
		return
	}

	response.JSON(w, http.StatusOK, comments)
}
