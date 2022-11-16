package service

import (
	"lab3/internal/domain"
	"lab3/internal/repository"
	"time"
)

type service struct {
	repo repository.Repository
}

type Service interface {
	SelectCustomers() ([]domain.Customer, error)
	SelectMovies() (movies []domain.Movie, err error)
	SelectHalls() ([]domain.Hall, error)
	SelectSessions() ([]domain.Session, error)
	SelectTickets() ([]domain.Ticket, error)
	CreateMovie(movie domain.Movie) (int, error)
	CreateCustomer(customer domain.Customer) (int, error)
	CreateSession(session domain.Session) (int, error)
	CreateTicket(ticket domain.Ticket) (int, error)
	SearchSessions(params domain.SessionsSearchParams) ([]domain.SelectSessionDTO, time.Duration, error)
	SearchTickets(params domain.TicketsSearchParams) ([]domain.SelectTicketDTO, time.Duration, error)
	SearchHalls(params domain.HallsSearchParams) ([]domain.SelectHallDTO, time.Duration, error)
	CreateRandomMovies(amount int) error
	CreateRandomSessions(amount int) error
	DeleteCustomer(id int) error
	DeleteMovie(id int) error
	DeleteSession(id int) error
	DeleteTicket(id int) error
	UpdateCustomer(customer domain.Customer) error
	UpdateMovie(movie domain.Movie) error
	UpdateSession(session domain.Session) error
	UpdateTicket(ticket domain.Ticket) error
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) SelectCustomers() ([]domain.Customer, error) {
	return nil, nil
}

func (s *service) SelectMovies() (movies []domain.Movie, err error) {
	return nil, err
}

func (s *service) SelectHalls() ([]domain.Hall, error) {
	return nil, nil
}

func (s *service) SelectSessions() ([]domain.Session, error) {
	return nil, nil
}

func (s *service) SelectTickets() ([]domain.Ticket, error) {
	return nil, nil
}

func (s *service) CreateMovie(movie domain.Movie) (int, error) {
	return 0, nil
}

func (s *service) CreateCustomer(customer domain.Customer) (int, error) {
	return 0, nil
}

func (s *service) CreateSession(session domain.Session) (int, error) {
	return 0, nil
}

func (s *service) CreateTicket(ticket domain.Ticket) (int, error) {
	return 0, nil
}

func (s *service) SearchSessions(params domain.SessionsSearchParams) ([]domain.SelectSessionDTO, time.Duration, error) {
	return nil, 0, nil
}

func (s *service) SearchTickets(params domain.TicketsSearchParams) ([]domain.SelectTicketDTO, time.Duration, error) {
	return nil, 0, nil
}

func (s *service) SearchHalls(params domain.HallsSearchParams) ([]domain.SelectHallDTO, time.Duration, error) {
	return nil, 0, nil
}

func (s *service) CreateRandomMovies(amount int) error {
	return nil
}

func (s *service) CreateRandomSessions(amount int) error {
	return nil
}

func (s *service) DeleteCustomer(id int) error {
	return nil
}
func (s *service) DeleteMovie(id int) error {
	return nil
}
func (s *service) DeleteSession(id int) error {
	return nil
}
func (s *service) DeleteTicket(id int) error {
	return nil
}

//

func (s *service) UpdateCustomer(customer domain.Customer) error {
	return nil
}

func (s *service) UpdateMovie(movie domain.Movie) error {
	return nil
}

func (s *service) UpdateSession(session domain.Session) error {
	return nil
}

func (s *service) UpdateTicket(ticket domain.Ticket) error {
	return nil
}
