package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"lab2/internal/domain"
	"lab2/pkg/client/postgresql"
	"time"
)

type storage struct {
	db *pgx.Conn
}

type Repository interface {
	SelectCustomers() ([]domain.Customer, error)
	SelectMovies() ([]domain.Movie, error)
	SelectHalls() ([]domain.Hall, error)
	SelectSessions() ([]domain.Session, error)
	SelectTickets() ([]domain.Ticket, error)
	InsertMovie(movie domain.Movie) (int, error)
	InsertCustomer(customer domain.Customer) (int, error)
	InsertSession(session domain.Session) (int, error)
	InsertTicket(ticket domain.Ticket) (int, error)
	SearchSessions(params domain.SessionsSearchParams) ([]domain.SelectSessionDTO, time.Duration, error)
	SearchTickets(params domain.TicketsSearchParams) ([]domain.SelectTicketsDTO, time.Duration, error)
	SearchHalls(params domain.SessionsSearchParams) ([]domain.SelectSessionDTO, time.Duration, error)
}

func NewRepository(db *pgx.Conn) Repository {
	return &storage{db: db}
}

func (s *storage) SelectCustomers() (customers []domain.Customer, err error) {
	querySelectCustomers := fmt.Sprintf(`SELECT id, first_name, last_name FROM %s`, postgresql.CustomerTable)
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

func (s *storage) SelectMovies() (movies []domain.Movie, err error) {
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

func (s *storage) SelectHalls() (halls []domain.Hall, err error) {
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

func (s *storage) SelectSessions() (sessions []domain.Session, err error) {
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

func (s *storage) SelectTickets() (tickets []domain.Ticket, err error) {
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

func (s *storage) InsertCustomer(customer domain.Customer) (int, error) {
	var customerId int

	queryInsertCustomer := fmt.Sprintf(
		"INSERT INTO %s (first_name, last_name) values ($1, $2) RETURNING id",
		postgresql.CustomerTable,
	)

	row := s.db.QueryRow(queryInsertCustomer, customer.FirstName, customer.LastName)

	if err := row.Scan(&customerId); err != nil {
		return 0, err
	}

	return customerId, nil
}

func (s *storage) InsertSession(session domain.Session) (int, error) {
	var sessionId int

	queryInsertSession := fmt.Sprintf(
		"INSERT INTO %s (movie_id, hall_id, start_at) values ($1, $2, $3) RETURNING id",
		postgresql.SessionTable,
	)

	row := s.db.QueryRow(queryInsertSession, session.MovieId, session.HallId, session.StartAt)

	if err := row.Scan(&sessionId); err != nil {
		return 0, err
	}

	return sessionId, nil
}

func (s *storage) InsertTicket(ticket domain.Ticket) (int, error) {
	var ticketId int

	queryInsertTicket := fmt.Sprintf(
		`
		INSERT INTO %s (session_id, customer_id, price, row_id, position_id) 
		values($1, $2, $3, $4, $5) 
		RETURNING id
		`,
		postgresql.TicketTable,
	)

	row := s.db.QueryRow(
		queryInsertTicket,
		ticket.SessionId,
		ticket.CustomerId,
		ticket.Price,
		ticket.RowId,
		ticket.PositionId,
	)

	if err := row.Scan(&ticketId); err != nil {
		if errors.Is(err, pgx.PgError{}) {
			return 0, fmt.Errorf("failed to insert new ticket")
		}
		return 0, err
	}

	return ticketId, nil
}

func (s *storage) SearchSessions(params domain.SessionsSearchParams) (sessions []domain.SelectSessionDTO, d time.Duration, err error) {
	querySearchSessions := fmt.Sprintf(
		`
		SELECT session.id,
				movie.title as movie,
		   		session.start_at,
		   		hall.title  as hall
		FROM session
			 JOIN movie on movie.id = session.movie_id
			 JOIN hall on hall.id = session.hall_id
		WHERE LOWER(movie.title) LIKE LOWER(CONCAT('%%',$1::varchar,'%%'))
	  		AND session.start_at > $2
	  		AND session.start_at < $3;
		`,
	)

	start := time.Now()

	rows, err := s.db.Query(querySearchSessions, params.MovieName, params.StartAtGt, params.StartAtLt)
	if err != nil {
		return nil, 0, err
	}

	d = time.Now().Sub(start)

	for rows.Next() {
		var sn domain.SelectSessionDTO

		if err = rows.Scan(&sn.Id, &sn.Movie, &sn.StartAt, &sn.Hall); err != nil {
			return nil, 0, err
		}

		sessions = append(sessions, sn)
	}

	if rows.Err() != nil {
		return nil, 0, nil
	}

	return sessions, d, err
}

func (s *storage) SearchTickets(params domain.TicketsSearchParams) (tickets []domain.SelectTicketsDTO, d time.Duration, err error) {
	querySearchSessions := fmt.Sprintf(
		`
		SELECT 
			ticket.id,
			movie.title as movie,
			session.start_at,
			hall.title as hall,
			customer.first_name,
			customer.last_name,
			ticket.price,
			movie.duration,
			row.number_in_hall as row,
			position.number_in_row as position
		FROM ticket
		JOIN session ON ticket.session_id = session.id
		JOIN customer ON ticket.customer_id = customer.id
		JOIN movie ON session.movie_id = movie.id
		JOIN hall ON session.hall_id = hall.id
		JOIN row ON ticket.row_id = row.id
		JOIN position ON ticket.position_id = position.id
		WHERE ticket.price >= $1 
			AND ticket.price <= $2
			AND movie.duration >= $3
			AND movie.duration <= $4
		`,
	)

	start := time.Now()

	rows, err := s.db.Query(
		querySearchSessions,
		params.PriceGt,
		params.PriceLt,
		params.MovieDurationGt,
		params.MovieDurationLt,
	)
	if err != nil {
		return nil, 0, err
	}

	d = time.Now().Sub(start)

	for rows.Next() {
		var dto domain.SelectTicketsDTO

		if err = rows.Scan(
			&dto.Id,
			&dto.MovieTitle,
			&dto.SessionStartAt,
			&dto.HallTitle,
			&dto.CustomerFirstname,
			&dto.CustomerLastname,
			&dto.Price,
			&dto.MovieDuration,
			&dto.Row,
			&dto.Position,
		); err != nil {
			return nil, 0, err
		}

		tickets = append(tickets, dto)
	}

	if rows.Err() != nil {
		return nil, 0, nil
	}

	return tickets, d, err
}

func (s *storage) SearchHalls(params domain.SessionsSearchParams) ([]domain.SelectSessionDTO, time.Duration, error) {
	return nil, 0, nil
}
