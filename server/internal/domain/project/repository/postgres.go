package repository

import (
	"bugtracker/internal/domain/project/aggregate"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type PostgresProjectRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresProjectRepository {
	return &PostgresProjectRepository{
		db: db,
	}
}

func (r *PostgresProjectRepository) Save(ctx context.Context, name, description, ownerId string) (*aggregate.Project, error) {
	project := aggregate.NewProject(name, description, ownerId)
	query := sq.Insert("projects").
		Columns("id", "name", "description", "owner_id").
		Values(project.ID, project.Name, project.Description, project.OwnerId).PlaceholderFormat(sq.Dollar).RunWith(r.db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *PostgresProjectRepository) GetProjectsByOwnerId(ctx context.Context, userId string) ([]aggregate.Project, error) {
	query := sq.Select("*").From("projects").Where(sq.Eq{"owner_id": userId}).PlaceholderFormat(sq.Dollar).RunWith(r.db)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []aggregate.Project
	for rows.Next() {
		var project aggregate.Project
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.OwnerId,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}