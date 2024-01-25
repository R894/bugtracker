package repository

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/user/aggregate"
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type PostgresUserRepository struct {
	db database.PostgresDB
}

func NewPostgresUserRepository(db database.PostgresDB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) SaveUser(ctx context.Context, user aggregate.User) error {
	sql, args, err := sq.Insert("users").
		Columns("id", "username", "email", "password", "createdAt", "updatedAt").
		Values(user.ID, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
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

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, userID string) (*aggregate.User, error) {
	var user aggregate.User
	sql, _, err := sq.Select("*").From("users").Where(sq.Eq{"id": userID}).ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Db.QueryRow(sql).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user aggregate.User) error {
	sql, args, err := sq.Update("users").
		Set("username", user.Username).
		Set("password", user.Password).
		Set("email", user.Email).
		Set("updatedAt", time.Now()).
		Where(sq.Eq{"id": user.ID}).
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

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, userID string) error {
	sql, args, err := sq.Delete("users").Where(sq.Eq{"id": userID}).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
