package service

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/comment/aggregate"
	"bugtracker/internal/domain/comment/repository"
	"context"
)

type CommentService struct {
	repo *repository.PostgresCommentRepository
}

func NewCommentService(db *database.PostgresDB) *CommentService {
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