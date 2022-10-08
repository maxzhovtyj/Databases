package repository

import (
	"fmt"
	"github.com/jackc/pgx"
	"lab2/internal/domain"
	"lab2/pkg/client/postgresql"
)

type storage struct {
	db *pgx.Conn
}

type Repository interface {
	GetCustomers() ([]domain.Customer, error)
	GetMovies() ([]domain.Movie, error)
	GetHalls() ([]domain.Hall, error)
	GetSessions() ([]domain.Session, error)
}

func NewRepository(db *pgx.Conn) Repository {
	return &storage{db: db}
}

func (s *storage) GetCustomers() (customers []domain.Customer, err error) {
	querySelectCustomers := fmt.Sprintf(`SELECT id, first_name, last_name FROM %s`, postgresql.CustomersTable)
	rows, err := s.db.Query(querySelectCustomers)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ctr domain.Customer

		err = rows.Scan(&ctr.Id, &ctr.FirstName, &ctr.LastName)
		if err != nil {
			return nil, err
		}

		customers = append(customers, ctr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, err
}

func (s *storage) GetMovies() (movies []domain.Movie, err error) {
	querySelectMovies := fmt.Sprintf(`SELECT id, title, description, duration FROM %s`, postgresql.MovieTable)
	rows, err := s.db.Query(querySelectMovies)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ctr domain.Movie

		err = rows.Scan(&ctr.Id, &ctr.Title, &ctr.Description, &ctr.Duration)
		if err != nil {
			return nil, err
		}

		movies = append(movies, ctr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, err
}

func (s *storage) GetHalls() (halls []domain.Hall, err error) {
	querySelectHalls := fmt.Sprintf(`SELECT id, title, description, capacity FROM %s`, postgresql.HallTable)
	rows, err := s.db.Query(querySelectHalls)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ctr domain.Hall

		err = rows.Scan(&ctr.Id, &ctr.Title, &ctr.Description, &ctr.Capacity)
		if err != nil {
			return nil, err
		}

		halls = append(halls, ctr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return halls, err
}

func (s *storage) GetSessions() (sessions []domain.Session, err error) {
	querySelectSessions := fmt.Sprintf(`SELECT id, movie_id, hall_id, start_at FROM %s`, postgresql.SessionTable)
	rows, err := s.db.Query(querySelectSessions)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ctr domain.Session

		err = rows.Scan(&ctr.Id, &ctr.MovieId, &ctr.HallId, &ctr.StartAt)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, ctr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sessions, err
}
