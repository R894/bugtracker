package repository

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/comment/aggregate"
	"context"

	sq "github.com/Masterminds/squirrel"
)

type PostgresCommentRepository struct {
	db *database.PostgresDB
}

func NewPostgresCommentRepository(db *database.PostgresDB) *PostgresCommentRepository {
	return &PostgresCommentRepository{
		db: db,
	}
}

func (c *PostgresCommentRepository) SaveComment(ctx context.Context, comment aggregate.Comment) error {
	query := sq.Insert("comments").Columns("id", "bugId", "content", "createdAt", "updatedAt").
		Values(comment.ID, comment.BugId, comment.Content, comment.CreatedAt, comment.UpdatedAt).PlaceholderFormat(sq.Dollar).RunWith(c.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCommentRepository) UpdateComment(ctx context.Context, commentId, content string) error {
	query := sq.Update("comments").Set("content", content).Where(sq.Eq{"id": commentId}).PlaceholderFormat(sq.Dollar).RunWith(c.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCommentRepository) DeleteComment(ctx context.Context, commentId string) error {
	query := sq.Delete("comments").Where(sq.Eq{"id": commentId}).PlaceholderFormat(sq.Dollar).RunWith(c.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
