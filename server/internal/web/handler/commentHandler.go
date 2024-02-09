package handler

import (
	"bugtracker/internal/domain/comment/service"
	"bugtracker/internal/web/response"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type CommentHandler struct {
	service *service.CommentService
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

func NewCommentHandler(db *sql.DB) *CommentHandler {
	return &CommentHandler{
		service: service.NewCommentService(db),
	}
}

func (c *CommentHandler) CreateNewComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req newCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w)
		return
	}

	err = c.service.NewComment(ctx, req.BugId, req.Content)
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
