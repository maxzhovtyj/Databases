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
	SearchSessions(params domain.SearchSessionsParams) ([]domain.SelectSessionDTO, time.Duration, error)
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

func (s *service) SearchSessions(params domain.SearchSessionsParams) ([]domain.SelectSessionDTO, time.Duration, error) {
	return s.repo.SearchSessions(params)
}
