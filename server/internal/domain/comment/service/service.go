package service

import (
	"bugtracker/internal/domain/comment/aggregate"
	"bugtracker/internal/domain/comment/repository"
	"context"
	"database/sql"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{
		repo: repository.NewPostgresCommentRepository(db),
	}
}

func (service *CommentService) NewComment(ctx context.Context, bugId, content string) error {
	comment := aggregate.NewComment(bugId, content)
	err := service.repo.SaveComment(ctx, comment)
	if err != nil {
		return err
	}

	return nil
}

func (service *CommentService) UpdateComment(ctx context.Context, commentId, content string) error {
	err := service.repo.UpdateComment(ctx, commentId, content)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) DeleteComment(ctx context.Context, commentId string) error {
	err := service.repo.DeleteComment(ctx, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) GetByBugId(ctx context.Context, bugId string) ([]aggregate.Comment, error) {
	bugs, err := service.repo.GetCommentsByBugId(ctx, bugId)
	if err != nil {
		return nil, err
	}
	return bugs, nil
}
