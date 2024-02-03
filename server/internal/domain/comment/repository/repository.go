package repository

import (
	"bugtracker/internal/domain/comment/aggregate"
	"context"
)

type CommentRepository interface {
	SaveComment(ctx context.Context, comment aggregate.Comment) error
	UpdateComment(ctx context.Context, commentId, content string) error
	DeleteComment(ctx context.Context, commentId string) error
	GetCommentsByBugId(ctx context.Context, bugId string) ([]aggregate.Comment, error)
}
