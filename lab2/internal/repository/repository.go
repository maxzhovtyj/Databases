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
	GetTickets() ([]domain.Ticket, error)
	InsertMovie(movie domain.Movie) (int, error)
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

func (s *storage) GetTickets() (tickets []domain.Ticket, err error) {
	querySelectTickets := fmt.Sprintf(
		"SELECT id, customer_id, session_id, price, row_id, position_id FROM %s",
		postgresql.TicketTable,
	)

	rows, err := s.db.Query(querySelectTickets)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ticket domain.Ticket

		if err = rows.Scan(
			&ticket.Id,
			&ticket.CustomerId,
			&ticket.SessionId,
			&ticket.Price,
			&ticket.RowId,
			&ticket.PositionId,
		); err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, err
}

func (s *storage) InsertMovie(movie domain.Movie) (int, error) {
	var movieId int

	queryInsertMovie := fmt.Sprintf(
		"INSERT INTO %s (title, description, duration) values ($1, $2, $3) RETURNING id",
		postgresql.MovieTable,
	)

	row := s.db.QueryRow(queryInsertMovie, movie.Title, movie.Description, movie.Duration)

	if err := row.Scan(&movieId); err != nil {
		return 0, err
	}

	return movieId, nil
}
