package repository

import "context"

type CommentRepository interface {
	CreateComment(ctx context.Context, bugId, content string)
	UpdateComment(ctx context.Context, commentId, content string)
	DeleteComment(ctx context.Context, commentId string)
}
