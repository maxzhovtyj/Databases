package service

import (
	"lab2/internal/domain"
	"lab2/internal/repository"
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
	return s.repo.SelectCustomers()
}

func (s *service) SelectMovies() (movies []domain.Movie, err error) {
	return s.repo.SelectMovies()
}

func (s *service) SelectHalls() ([]domain.Hall, error) {
	return s.repo.SelectHalls()
}

func (s *service) SelectSessions() ([]domain.Session, error) {
	return s.repo.SelectSessions()
}

func (s *service) SelectTickets() ([]domain.Ticket, error) {
	return s.repo.SelectTickets()
}

func (s *service) CreateMovie(movie domain.Movie) (int, error) {
	return s.repo.InsertMovie(movie)
}

func (s *service) CreateCustomer(customer domain.Customer) (int, error) {
	return s.repo.InsertCustomer(customer)
}

func (s *service) CreateSession(session domain.Session) (int, error) {
	return s.repo.InsertSession(session)
}

func (s *service) CreateTicket(ticket domain.Ticket) (int, error) {
	return s.repo.InsertTicket(ticket)
}

func (s *service) SearchSessions(params domain.SessionsSearchParams) ([]domain.SelectSessionDTO, time.Duration, error) {
	return s.repo.SearchSessions(params)
}

func (s *service) SearchTickets(params domain.TicketsSearchParams) ([]domain.SelectTicketDTO, time.Duration, error) {
	return s.repo.SearchTickets(params)
}

func (s *service) SearchHalls(params domain.HallsSearchParams) ([]domain.SelectHallDTO, time.Duration, error) {
	return s.repo.SearchHalls(params)
}

func (s *service) CreateRandomMovies(amount int) error {
	return s.repo.InsertRandomisedMovies(amount)
}

func (s *service) CreateRandomSessions(amount int) error {
	return s.repo.InsertRandomisedSessions(amount)
}

func (s *service) DeleteCustomer(id int) error {
	return s.repo.DeleteCustomer(id)
}
func (s *service) DeleteMovie(id int) error {
	return s.repo.DeleteMovie(id)
}
func (s *service) DeleteSession(id int) error {
	return s.repo.DeleteSession(id)
}
func (s *service) DeleteTicket(id int) error {
	return s.repo.DeleteTicket(id)
}

//

func (s *service) UpdateCustomer(customer domain.Customer) error {
	return s.repo.UpdateCustomer(customer)
}

func (s *service) UpdateMovie(movie domain.Movie) error {
	return s.repo.UpdateMovie(movie)
}

func (s *service) UpdateSession(session domain.Session) error {
	return s.repo.UpdateSession(session)
}

func (s *service) UpdateTicket(ticket domain.Ticket) error {
	return s.repo.UpdateTicket(ticket)
}
