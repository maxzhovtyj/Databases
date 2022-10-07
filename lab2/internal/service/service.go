package service

import (
	"lab2/internal/repository"
)

type service struct {
	repo repository.Repository
}

type Service interface {
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
