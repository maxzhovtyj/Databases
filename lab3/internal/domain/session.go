package domain

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	Tickets []Ticket
	MovieID uint
	HallID  uint
	StartAt time.Time
}

type SessionsSearchParams struct {
	MovieName string
	StartAtGt time.Time
	StartAtLt time.Time
}
