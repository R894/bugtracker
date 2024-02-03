package repository

import (
	"bugtracker/internal/domain/comment/aggregate"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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
	query := sq.Insert("comments").Columns("id", "bugId", "content", "createdAt", "updatedAt").
		Values(comment.ID, comment.BugId, comment.Content, comment.CreatedAt, comment.UpdatedAt).PlaceholderFormat(sq.Dollar).RunWith(c.db)

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
		rows.Scan(&comment.ID, &comment.BugId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		comments = append(comments, comment)
	}

	return comments, nil
}
