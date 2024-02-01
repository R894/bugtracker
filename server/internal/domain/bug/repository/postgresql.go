package repository

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/bug/aggregate"
	"context"

	sq "github.com/Masterminds/squirrel"
)

type PostgresBugRepository struct {
	db *database.PostgresDB
}

func NewPostgresBugRepository(db *database.PostgresDB) *PostgresBugRepository {
	return &PostgresBugRepository{
		db: db,
	}
}

func (r *PostgresBugRepository) SaveBug(ctx context.Context, bug *aggregate.Bug) error {
	query := sq.Insert("bugs").
		Columns("id", "title", "description", "status", "priority", "assignee", "created_at", "updated_at").
		Values(bug.ID, bug.Title, bug.Description, bug.Status, bug.Priority, bug.Assignee, bug.CreatedAt, bug.UpdatedAt).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresBugRepository) GetBugByID(ctx context.Context, bugID string) (*aggregate.Bug, error) {
	var bug aggregate.Bug
	query := sq.Select("*").From("bugs").Where(sq.Eq{"id": bugID}).PlaceholderFormat(sq.Dollar).RunWith(r.db.Db)

	err := query.QueryRow().Scan(&bug.ID, &bug.Title, &bug.Description, &bug.Status, &bug.Priority, &bug.Assignee, &bug.CreatedAt, &bug.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &bug, nil
}

func (r *PostgresBugRepository) GetBugs(ctx context.Context) ([]aggregate.Bug, error) {
	query := sq.Select("*").From("bugs").PlaceholderFormat(sq.Dollar).RunWith(r.db.Db)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// TODO: refactor this dumpsterfuck
	var bugs []aggregate.Bug
	for rows.Next() {
		var bug aggregate.Bug
		err := rows.Scan(
			&bug.ID,
			&bug.Title,
			&bug.Description,
			&bug.Status,
			&bug.Priority,
			&bug.Assignee,
			&bug.CreatedAt,
			&bug.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bugs = append(bugs, bug)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bugs, nil
}

func (r *PostgresBugRepository) UpdateBug(ctx context.Context, bug *aggregate.UpdateBugRequest) error {
	query := sq.Update("bugs").
		Set("title", bug.Title).
		Set("description", bug.Description).
		Set("status", bug.Status).
		Set("priority", bug.Priority).
		Set("assignee", bug.Assignee).
		Set("updated_at", bug.UpdatedAt).
		Where(sq.Eq{"id": bug.ID}).
		PlaceholderFormat(sq.Dollar).RunWith(r.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresBugRepository) AssignBugTo(ctx context.Context, bugId, username string) error {
	query := sq.Update("bugs").
		Set("assignee", username).
		Where(sq.Eq{"id": bugId}).
		PlaceholderFormat(sq.Dollar).RunWith(r.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresBugRepository) DeleteBug(ctx context.Context, bugID string) error {
	query := sq.Delete("bugs").Where(sq.Eq{"id": bugID}).PlaceholderFormat(sq.Dollar).RunWith(r.db.Db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
