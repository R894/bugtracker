package infrastructure

import (
	"bugtracker/internal/bug/domain"
	"bugtracker/internal/database"
	"context"

	sq "github.com/Masterminds/squirrel"
)

type BugRepository struct {
	db database.PostgresDB
}

func NewBugRepository(db database.PostgresDB) *BugRepository {
	return &BugRepository{
		db: db,
	}
}

func (r *BugRepository) SaveBug(ctx context.Context, bug *domain.Bug) error {
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

func (r *BugRepository) GetBugByID(ctx context.Context, bugID string) (*domain.Bug, error) {
	var bug domain.Bug
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

func (r *BugRepository) UpdateBug(ctx context.Context, bug *domain.Bug) error {
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

func (r *BugRepository) DeleteBug(ctx context.Context, bugID string) error {
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
