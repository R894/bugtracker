package infrastructure

import (
	"bugtracker/internal/database"
	"bugtracker/internal/user/domain"
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db database.PostgresDB
}

func NewUserRepository(db database.PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SaveUser(ctx context.Context, user domain.User) error {
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

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	var user domain.User
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

func (r *UserRepository) UpdateUser(ctx context.Context, user domain.User) error {
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

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
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
