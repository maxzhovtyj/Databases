package repository

import (
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

type Repository interface {
}

func NewRepository(db *gorm.DB) Repository {
	return &storage{db: db}
}

