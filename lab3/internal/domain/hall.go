package domain

import (
	"gorm.io/gorm"
)

type Hall struct {
	gorm.Model
	Sessions    []Session
	Title       string
	Description string
	Rows        []Row
	Capacity    *uint
}

type MovieSearchParams struct {
	CreatedAtGt string
	CreatedAtLt string
}
