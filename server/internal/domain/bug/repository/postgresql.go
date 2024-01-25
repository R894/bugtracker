package repository

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/bug/aggregate"
	"context"

	sq "github.com/Masterminds/squirrel"
)

type PostgresBugRepository struct {
	db database.PostgresDB
}

func NewPostgresBugRepository(db database.PostgresDB) *PostgresBugRepository {
	return &PostgresBugRepository{
		db: db,
	}
}

func (r *PostgresBugRepository) SaveBug(ctx context.Context, bug *aggregate.Bug) error {
	sql, args, err := sq.Insert("bugs").
		Columns("id", "title", "description", "status", "priority", "assignee", "created_at", "updated_at").
		Values(bug.ID, bug.Title, bug.Description, bug.Status, bug.Priority, bug.Assignee, bug.CreatedAt, bug.UpdatedAt).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresBugRepository) GetBugByID(ctx context.Context, bugID string) (*aggregate.Bug, error) {
	var bug aggregate.Bug
	sql, _, err := sq.Select("*").From("bugs").Where(sq.Eq{"id": bugID}).ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Db.QueryRow(sql).Scan(&bug)
	if err != nil {
		return nil, err
	}

	return &bug, nil
}

func (r *PostgresBugRepository) UpdateBug(ctx context.Context, bug *aggregate.Bug) error {
	sql, args, err := sq.Update("bugs").
		Set("title", bug.Title).
		Set("description", bug.Description).
		Set("status", bug.Status).
		Set("priority", bug.Priority).
		Set("assignee", bug.Assignee).
		Set("updated_at", bug.UpdatedAt).
		Where(sq.Eq{"id": bug.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresBugRepository) DeleteBug(ctx context.Context, bugID string) error {
	sql, args, err := sq.Delete("bugs").Where(sq.Eq{"id": bugID}).ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
