package repository

import "context"

type CommentRepository interface {
	CreateComment(ctx context.Context, bugId, content string) error
	UpdateComment(ctx context.Context, commentId, content string) error
	DeleteComment(ctx context.Context, commentId string) error
}
