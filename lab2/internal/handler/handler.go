package handler

import "lab2/internal/service"

type handler struct {
	service service.Service
}

type Handler interface {
	Search() error
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Search() error {
	return nil
}
