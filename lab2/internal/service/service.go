package service

import (
	"lab2/internal/domain"
	"lab2/internal/repository"
)

type service struct {
	repo repository.Repository
}

type Service interface {
	GetCustomers() ([]domain.Customer, error)
	GetMovies() (movies []domain.Movie, err error)
	GetHalls() ([]domain.Hall, error)
	GetSessions() ([]domain.Session, error)
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetCustomers() ([]domain.Customer, error) {
	return s.repo.GetCustomers()
}

func (s *service) GetMovies() (movies []domain.Movie, err error) {
	return s.repo.GetMovies()
}

func (s *service) GetHalls() ([]domain.Hall, error) {
	return s.repo.GetHalls()
}

func (s *service) GetSessions() ([]domain.Session, error) {
	return s.repo.GetSessions()
}
