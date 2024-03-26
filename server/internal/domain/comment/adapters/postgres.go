package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/R894/bugtracker/internal/domain/comment/core/domain/aggregate"
)

type PostgresCommentRepository struct {
	db *sql.DB
}

func NewPostgresCommentRepository(db *sql.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{
		db: db,
	}
}

func (c *PostgresCommentRepository) SaveComment(ctx context.Context, comment aggregate.Comment) error {
	query := sq.Insert("comments").Columns("id", "bug_id", "user_id", "content", "created_at", "updated_at").
		Values(comment.ID, comment.BugId, comment.UserId, comment.Content, comment.CreatedAt, comment.UpdatedAt).PlaceholderFormat(sq.Dollar).RunWith(c.db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCommentRepository) UpdateComment(ctx context.Context, commentId, content string) error {
	query := sq.Update("comments").Set("content", content).Where(sq.Eq{"id": commentId}).PlaceholderFormat(sq.Dollar).RunWith(c.db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCommentRepository) DeleteComment(ctx context.Context, commentId string) error {
	query := sq.Delete("comments").Where(sq.Eq{"id": commentId}).PlaceholderFormat(sq.Dollar).RunWith(c.db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCommentRepository) GetCommentsByBugId(ctx context.Context, bugId string) ([]aggregate.Comment, error) {
	query := sq.Select("*").From("comments").Where(sq.Eq{"bug_id": bugId}).PlaceholderFormat(sq.Dollar).RunWith(c.db)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []aggregate.Comment
	for rows.Next() {
		var comment aggregate.Comment
		if err := rows.Scan(&comment.ID, &comment.BugId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.UserId); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
