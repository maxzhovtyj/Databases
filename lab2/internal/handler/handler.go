package handler

import (
	"lab2/internal/domain"
	"lab2/internal/service"
)

type handler struct {
	service service.Service
}

type Handler interface {
	GetCustomers() ([]domain.Customer, error)
	GetMovies() (movies []domain.Movie, err error)
	GetHalls() ([]domain.Hall, error)
	GetSessions() ([]domain.Session, error)
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) GetCustomers() ([]domain.Customer, error) {
	customers, err := h.service.GetCustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (h *handler) GetMovies() ([]domain.Movie, error) {
	movies, err := h.service.GetMovies()
	if err != nil {
		return nil, err
	}
	return movies, err
}

func (h *handler) GetHalls() ([]domain.Hall, error) {
	halls, err := h.service.GetHalls()
	if err != nil {
		return nil, err
	}

	return halls, err
}

func (h *handler) GetSessions() ([]domain.Session, error) {
	sessions, err := h.service.GetSessions()
	if err != nil {
		return nil, err
	}

	return sessions, err
}
