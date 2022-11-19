package repository

import (
	"fmt"
	"gorm.io/gorm"
	"lab3/internal/domain"
	"time"
)

type storage struct {
	db *gorm.DB
}

type Repository interface {
	SelectCustomers() ([]domain.Customer, error)
	SelectMovies() ([]domain.Movie, error)
	SelectHalls() ([]domain.Hall, error)
	SelectSessions() ([]domain.Session, error)
	SelectTickets() ([]domain.Ticket, error)
	InsertMovie(movie domain.Movie) (uint, error)
	InsertCustomer(customer domain.Customer) (uint, error)
	InsertSession(session domain.Session) (uint, error)
	InsertTicket(ticket domain.Ticket) (uint, error)
	SearchSessions(params domain.SessionsSearchParams) ([]domain.Session, time.Duration, error)
	SearchMovies(params domain.MovieSearchParams) ([]domain.Movie, time.Duration, error)
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
	InsertPositions() error
	DeletePosition() error
}

func NewRepository(db *gorm.DB) Repository {
	return &storage{db: db}
}

func (s *storage) SelectCustomers() (customers []domain.Customer, err error) {
	resFind := s.db.Find(&customers)
	if resFind.Error != nil {
		return nil, fmt.Errorf("failed to find customers, %v", err)
	}

	return customers, err
}

func (s *storage) SelectMovies() (movies []domain.Movie, err error) {
	resFind := s.db.Find(&movies)
	if resFind.Error != nil {
		return nil, fmt.Errorf("failed to find movies, %v", err)
	}

	return movies, err
}

func (s *storage) SelectHalls() (halls []domain.Hall, err error) {
	resFind := s.db.Find(&halls)
	if resFind.Error != nil {
		return nil, fmt.Errorf("failed to find halls, %v", err)
	}

	return halls, err
}

func (s *storage) SelectSessions() (sessions []domain.Session, err error) {
	resFind := s.db.Find(&sessions)
	if resFind.Error != nil {
		return nil, fmt.Errorf("failed to find sessions, %v", err)
	}

	return sessions, err
}

func (s *storage) SelectTickets() (tickets []domain.Ticket, err error) {
	resFind := s.db.Find(&tickets)
	if resFind.Error != nil {
		return nil, fmt.Errorf("failed to find tickets, %v", err)
	}

	return tickets, err
}

func (s *storage) InsertMovie(movie domain.Movie) (uint, error) {
	createRes := s.db.Create(&movie)
	if createRes.Error != nil {
		return 0, fmt.Errorf("failed to create new movie, %v", createRes.Error)
	}

	return movie.Model.ID, nil
}

func (s *storage) InsertCustomer(customer domain.Customer) (uint, error) {
	createRes := s.db.Create(&customer)
	if createRes.Error != nil {
		return 0, fmt.Errorf("failed to create new customer, %v", createRes.Error)
	}

	return customer.Model.ID, nil
}

func (s *storage) InsertSession(session domain.Session) (uint, error) {
	createRes := s.db.Create(&session)
	if createRes.Error != nil {
		return 0, fmt.Errorf("failed to create new customer, %v", createRes.Error)
	}

	return session.Model.ID, nil
}

func (s *storage) InsertTicket(ticket domain.Ticket) (uint, error) {
	createRes := s.db.Create(&ticket)
	if createRes.Error != nil {
		return 0, fmt.Errorf("failed to create new ticket, %v", createRes.Error)
	}

	return ticket.Model.ID, nil
}

func (s *storage) SearchSessions(params domain.SessionsSearchParams) ([]domain.Session, time.Duration, error) {
	var sessions []domain.Session

	start := time.Now()

	find := s.db.
		Where("start_at >= ? AND start_at <= ?", params.StartAtGt, params.StartAtLt).
		Find(&sessions)

	if find.Error != nil {
		return nil, 0, fmt.Errorf("failed to find sessions, %v", find.Error)
	}

	queryTime := time.Now().Sub(start)

	return sessions, queryTime, nil
}

func (s *storage) SearchMovies(params domain.MovieSearchParams) ([]domain.Movie, time.Duration, error) {
	var movies []domain.Movie

	start := time.Now()

	find := s.db.
		Where("created_at >= ?::timestamptz AND created_at <= ?::timestamptz", params.CreatedAtGt, params.CreatedAtLt).
		Find(&movies)

	if find.Error != nil {
		return nil, 0, fmt.Errorf("failed to find movies, %v", find.Error)
	}

	queryTime := time.Now().Sub(start)

	return movies, queryTime, nil
}

func (s *storage) InsertRandomisedMovies(movieAmount int) error {
	var mv []domain.Movie
	rawQuery := s.db.Raw(
		`
		SELECT 
			chr(trunc(65+random()*25)::int) || chr(trunc(65+random()*25)::int) as title,
			chr(trunc(65+random()*25)::int) || chr(trunc(65+random()*25)::int) as description,
			trunc(random()  * 180) * '1 minute'::interval as duration
		FROM generate_series(1, ?)
		`,
		movieAmount,
	).Scan(&mv)
	if rawQuery.Error != nil {
		return fmt.Errorf("failed to insert randomised data, %v", rawQuery.Error)
	}

	createMovies := s.db.Create(&mv)
	if createMovies.Error != nil {
		return fmt.Errorf("failed to create movies, %v", createMovies.Error)
	}

	return nil
}

func (s *storage) InsertRandomisedSessions(amount int) error {
	var sessions []domain.Session
	rawQuery := s.db.Raw(
		`
		SELECT 
			trunc(random()*((SELECT MAX(id) FROM movies) - (SELECT MIN(id) FROM movies) + 1) + (SELECT MIN(id) FROM movies))::int as movie_id,
			trunc(random()*((SELECT MAX(id) FROM halls) - (SELECT MIN(id) FROM halls) + 1) + (SELECT MIN(id) FROM halls))::int as hall_id, 
			NOW() + (random() * (interval '90 days')) + '30 days' as start_at
		FROM generate_series(1, ?)
		`,
		amount,
	).Scan(&sessions)
	if rawQuery.Error != nil {
		return fmt.Errorf("failed to get random data, %v", rawQuery.Error)
	}

	createSessions := s.db.Create(&sessions)
	if createSessions.Error != nil {
		return fmt.Errorf("failed to create random sessions, %v", createSessions.Error)
	}

	return nil
}

func (s *storage) DeleteCustomer(id int) error {
	tx := s.db.Delete(&domain.Customer{}, id)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete customer, %v", tx.Error)
	}

	return nil
}

func (s *storage) DeleteMovie(id int) error {
	tx := s.db.Delete(&domain.Movie{}, id)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete movie, %v", tx.Error)
	}

	return nil
}

func (s *storage) DeleteSession(id int) error {
	tx := s.db.Delete(&domain.Session{}, id)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete session, %v", tx.Error)
	}

	return nil
}

func (s *storage) DeleteTicket(id int) error {
	tx := s.db.Delete(&domain.Ticket{}, id)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete ticket, %v", tx.Error)
	}

	return nil
}

func (s *storage) UpdateCustomer(customer domain.Customer) error {
	updCustomer := s.db.Updates(&customer)
	if updCustomer.Error != nil {
		return fmt.Errorf("failed to update customer, %v", updCustomer.Error)
	}

	return nil
}

func (s *storage) UpdateMovie(movie domain.Movie) error {
	updMovie := s.db.Updates(&movie)
	if updMovie.Error != nil {
		return fmt.Errorf("failed to update movie, %v", updMovie.Error)
	}

	return nil
}

func (s *storage) UpdateSession(session domain.Session) error {
	updSession := s.db.Updates(&session)
	if updSession.Error != nil {
		return fmt.Errorf("failed to update session, %v", updSession.Error)
	}

	return nil
}

func (s *storage) UpdateTicket(ticket domain.Ticket) error {
	updTicket := s.db.Updates(&ticket)
	if updTicket.Error != nil {
		return fmt.Errorf("failed to update ticket, %v", updTicket.Error)
	}

	return nil
}

func (s *storage) InsertPositions() error {
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 1,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 2,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 3,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 4,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 5,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 6,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 7,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 8,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 9,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 10,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 11,
	})
	s.db.Create(&domain.Position{
		RowID:       1,
		NumberInRow: 12,
	})

	return nil
}

func (s *storage) DeletePosition() error {
	s.db.Delete(&domain.Position{}, 41)

	return nil
}
