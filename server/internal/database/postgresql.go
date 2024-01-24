package database

import (
	"database/sql"
	"os"
)

type PostgresDB struct {
	Db *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		Db: db,
	}, nil
}

func (p *PostgresDB) Close() {
	p.Db.Close()
}
