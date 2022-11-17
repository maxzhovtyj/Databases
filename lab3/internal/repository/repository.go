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
	SearchTickets(params domain.TicketsSearchParams) ([]domain.Ticket, time.Duration, error)
	SearchHalls(params domain.HallsSearchParams) ([]domain.Hall, time.Duration, error)
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
	return nil, 0, nil
}

func (s *storage) SearchTickets(params domain.TicketsSearchParams) ([]domain.Ticket, time.Duration, error) {
	return nil, 0, nil
}

func (s *storage) SearchHalls(params domain.HallsSearchParams) ([]domain.Hall, time.Duration, error) {
	return nil, 0, nil
}

func (s *storage) InsertRandomisedMovies(movieAmount int) error {
	return nil
}

func (s *storage) InsertRandomisedSessions(amount int) error {
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
