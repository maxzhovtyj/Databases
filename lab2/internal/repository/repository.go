package repository

import "github.com/jackc/pgx"

type storage struct {
	db *pgx.Conn
}

type Repository interface {
}

func NewRepository(db *pgx.Conn) Repository {
	return &storage{db: db}
}
