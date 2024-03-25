package handler

import (
	"database/sql"

	"github.com/R894/bugtracker/internal/domain/bug/core/service"
	commentService "github.com/R894/bugtracker/internal/domain/comment/core/service"
	projectService "github.com/R894/bugtracker/internal/domain/project/service"
	userService "github.com/R894/bugtracker/internal/domain/user/service"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type AppHandler struct {
	Bug     *BugHandler
	User    *UserHandler
	Comment *CommentHandler
	Project *ProjectHandler
}

func New(db *sql.DB) *AppHandler {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	bugHandler := &BugHandler{
		service:  service.NewBugService(db),
		validate: validate,
		trans:    trans,
	}
	userHandler := &UserHandler{
		service:  userService.NewUserService(db),
		validate: validate,
		trans:    trans,
	}
	commentHandler := &CommentHandler{
		service:  commentService.NewCommentService(db),
		validate: validate,
		trans:    trans,
	}
	projectHandler := &ProjectHandler{
		service:  projectService.NewProjectService(db),
		validate: validate,
		trans:    trans,
	}

	return &AppHandler{
		Bug:     bugHandler,
		User:    userHandler,
		Comment: commentHandler,
		Project: projectHandler,
	}
}
