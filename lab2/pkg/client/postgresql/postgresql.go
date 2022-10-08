package postgresql

import (
	"context"
	"github.com/jackc/pgx"
)

const (
	CustomerTable = "customer"
	MovieTable    = "movie"
	TicketTable   = "ticket"
	HallTable     = "hall"
	RowTable      = "row"
	PositionTable = "position"
	SessionTable  = "session"
)

type StorageConfig struct {
	Username string
	Password string
	Host     string
	Port     uint16
	Database string
}

func NewPostgresClient(ctx context.Context, cfg StorageConfig) (*pgx.Conn, error) {
	dbConn, err := pgx.Connect(pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
		User:     cfg.Username,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, err
	}

	if err = dbConn.Ping(ctx); err != nil {
		return nil, err
	}

	return dbConn, err
}
