package domain

import "gorm.io/gorm"

type Row struct {
	gorm.Model
	Tickets      []Ticket
	Positions    []Position
	HallID       uint
	NumberInHall uint
}
