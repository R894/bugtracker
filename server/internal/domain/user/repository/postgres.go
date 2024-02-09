package repository

import (
	"bugtracker/internal/domain/user/aggregate"
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func userMap(user aggregate.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
		"updatedAt":  time.Now(),
		"projects":   pq.Array(user.Projects),
	}
}

func scanUserRow(rows *sql.Rows) (aggregate.User, error) {
	var user aggregate.User
	err := rows.Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		pq.Array(&user.Projects))

	return user, err
}

func (r *PostgresUserRepository) SaveUser(ctx context.Context, user aggregate.User) error {
	query := sq.Insert("users").
		Columns("id", "username", "email", "password", "created_at", "updated_at", "projects").
		Values(user.ID, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, pq.Array(user.Projects)).
		PlaceholderFormat(sq.Dollar).RunWith(r.db)

	_, err := query.ExecContext(ctx)
	if err != nil {
		log.Println("Error creating user: ", err)
		return err
	}

	return nil
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, userID string) (*aggregate.User, error) {
	var user aggregate.User
	query := sq.Select("*").From("users").Where(sq.Eq{"id": userID}).PlaceholderFormat(sq.Dollar).RunWith(r.db)

	err := query.QueryRowContext(ctx).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, pq.Array(&user.Projects))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetUserByName(ctx context.Context, username string) (*aggregate.User, error) {
	var user aggregate.User
	query := sq.Select("*").From("users").Where(sq.Eq{"username": username}).PlaceholderFormat(sq.Dollar).RunWith(r.db)

	err := query.QueryRowContext(ctx).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, pq.Array(&user.Projects))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, userEmail string) (*aggregate.User, error) {
	var user aggregate.User
	query := sq.Select("*").From("users").Where(sq.Eq{"email": userEmail}).PlaceholderFormat(sq.Dollar).RunWith(r.db)

	err := query.QueryRowContext(ctx).Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		pq.Array(&user.Projects))
	if err != nil {
		log.Println("Error getting user by email", err)
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context) ([]aggregate.User, error) {
	query := sq.Select("*").From("users").PlaceholderFormat(sq.Dollar).RunWith(r.db)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []aggregate.User
	for rows.Next() {
		user, err := scanUserRow(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user aggregate.User) error {
	query := sq.Update("users").SetMap(userMap(user)).Where(sq.Eq{"id": user.ID}).PlaceholderFormat(sq.Dollar).RunWith(r.db)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, userID string) error {
	query := sq.Delete("users").Where(sq.Eq{"id": userID}).PlaceholderFormat(sq.Dollar).RunWith(r.db)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}
