package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
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
	SearchTickets(params domain.TicketsSearchParams) ([]domain.SelectTicketDTO, time.Duration, error)
	SearchHalls(params domain.HallsSearchParams) ([]domain.SelectHallDTO, time.Duration, error)
	InsertRandomisedMovies(movieAmount int) error
	InsertRandomisedSessions(amount int) error
	DeleteCustomer(id int) error
	DeleteMovie(id int) error
	DeleteSession(id int) error
	DeleteTicket(id int) error
	UpdateCustomer(customer domain.Customer) error
	UpdateMovie(movie domain.Movie) error
	UpdateSession(session domain.Session) error
	UpdateTicket(ticket domain.Ticket) error
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
		var movieDuration time.Duration

		err = rows.Scan(&ctr.Id, &ctr.Title, &ctr.Description, &movieDuration)
		if err != nil {
			return nil, err
		}

		ctr.Duration = movieDuration.String()

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

//

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

//

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

func (s *storage) SearchTickets(params domain.TicketsSearchParams) (tickets []domain.SelectTicketDTO, d time.Duration, err error) {
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
		var dto domain.SelectTicketDTO

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

func (s *storage) SearchHalls(params domain.HallsSearchParams) (halls []domain.SelectHallDTO, d time.Duration, err error) {
	q := `
		SELECT hall.id,
        	hall.title,
        	hall.description,
        	hall.capacity,
        	array_agg(r.number_in_hall) as rows
		FROM hall
		LEFT JOIN row r on hall.id = r.hall_id
		WHERE LOWER(hall.title) LIKE LOWER(CONCAT('%', $1::varchar, '%'))
			AND hall.capacity >= $2 AND hall.capacity <= $3
		GROUP BY hall.id
	`

	start := time.Now()

	rows, err := s.db.Query(q, params.HallTitle, params.CapacityGt, params.CapacityLt)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to select halls")
	}

	queryTime := time.Now().Sub(start)

	for rows.Next() {
		var dto domain.SelectHallDTO
		var r pgtype.Int4Array
		if err = rows.Scan(
			&dto.Id,
			&dto.Title,
			&dto.Description,
			&dto.Capacity,
			&r,
		); err != nil {
			return nil, 0, err
		}

		for _, e := range r.Elements {
			dto.Rows = append(dto.Rows, e.Int)
		}

		halls = append(halls, dto)
	}

	return halls, queryTime, err
}

//

func (s *storage) InsertRandomisedMovies(movieAmount int) error {
	q :=
		`
		INSERT INTO 
		movie (title, description, duration) 
		SELECT 
			chr(trunc(65+random()*25)::int) || chr(trunc(65+random()*25)::int),
			chr(trunc(65+random()*25)::int) || chr(trunc(65+random()*25)::int),
			trunc(random()  * 180) * '1 minute'::interval 
		FROM generate_series(1, $1)
	`

	_, err := s.db.Exec(q, movieAmount)
	if err != nil {
		return fmt.Errorf("failed to insert movies")
	}

	return err
}

func (s *storage) InsertRandomisedSessions(amount int) error {
	q := `
	INSERT INTO 
	session (movie_id, hall_id, start_at) 
	SELECT 
		trunc(random()*((SELECT MAX(id) FROM movie) - 1 + 1) + 1)::int,
		trunc(random()*((SELECT MAX(id) FROM hall) - 1 + 1) + 1)::int,
		NOW() + (random() * (interval '90 days')) + '30 days'
	FROM generate_series(1, $1)
	`
	_, err := s.db.Exec(q, amount)
	if err != nil {
		return err
	}

	return nil
}

//

func (s *storage) DeleteCustomer(id int) error {
	queryDeleteCustomer := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgresql.CustomerTable)

	_, err := s.db.Exec(queryDeleteCustomer, id)
	if err != nil {
		return fmt.Errorf("failed to delete customer id = %d due to error: %v", id, err)
	}

	return nil
}

func (s *storage) DeleteMovie(id int) error {
	queryDeleteMovie := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgresql.MovieTable)

	_, err := s.db.Exec(queryDeleteMovie, id)
	if err != nil {
		return fmt.Errorf("failed to delete movie id = %d due to error: %v", id, err)
	}

	return nil
}

func (s *storage) DeleteSession(id int) error {
	queryDeleteSession := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgresql.SessionTable)

	_, err := s.db.Exec(queryDeleteSession, id)
	if err != nil {
		return fmt.Errorf("failed to delete session id = %d due to error: %v", id, err)
	}

	return nil
}

func (s *storage) DeleteTicket(id int) error {
	queryDeleteTicket := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgresql.TicketTable)

	_, err := s.db.Exec(queryDeleteTicket, id)
	if err != nil {
		return fmt.Errorf("failed to delete ticket id = %d due to error: %v", id, err)
	}

	return nil
}

//

func (s *storage) UpdateCustomer(customer domain.Customer) error {
	queryUpdateCustomer := fmt.Sprintf(
		"UPDATE %s SET first_name = $1, last_name = $2 WHERE id = $3",
		postgresql.CustomerTable,
	)

	_, err := s.db.Exec(queryUpdateCustomer, customer.FirstName, customer.LastName, customer.Id)
	if err != nil {
		return fmt.Errorf("failed to update customer id = %d, due to error: %v", customer.Id, err)
	}

	return err
}

func (s *storage) UpdateMovie(movie domain.Movie) error {
	queryUpdateMovie := fmt.Sprintf(
		"UPDATE %s SET title = $1, description = $2, duration = $3 WHERE id = $4",
		postgresql.MovieTable,
	)

	_, err := s.db.Exec(queryUpdateMovie, movie.Title, movie.Description, movie.Duration, movie.Id)
	if err != nil {
		return fmt.Errorf("failed to update movie id = %d, due to error: %v", movie.Id, err)
	}

	return err
}

func (s *storage) UpdateSession(session domain.Session) error {
	queryUpdateSession := fmt.Sprintf(
		"UPDATE %s SET movie_id = $1, hall_id = $2, start_at = $3 WHERE id = $4",
		postgresql.SessionTable,
	)

	_, err := s.db.Exec(queryUpdateSession, session.MovieId, session.HallId, session.StartAt, session.Id)
	if err != nil {
		return fmt.Errorf("failed to update session id = %d, due to error: %v", session.Id, err)
	}

	return err
}

func (s *storage) UpdateTicket(ticket domain.Ticket) error {
	queryUpdateSession := fmt.Sprintf(
		"UPDATE %s SET customer_id = $1, session_id = $2, price = $3, row_id = $4, position_id = $5 WHERE id = $6",
		postgresql.TicketTable,
	)

	_, err := s.db.Exec(
		queryUpdateSession,
		ticket.CustomerId,
		ticket.SessionId,
		ticket.Price,
		ticket.RowId,
		ticket.PositionId,
		ticket.Id,
	)
	if err != nil {
		return fmt.Errorf("failed to update ticket id = %d, due to error: %v", ticket.Id, err)
	}

	return err
}
