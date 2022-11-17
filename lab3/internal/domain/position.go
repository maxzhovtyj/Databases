package domain

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Tickets     []Ticket
	RowID       uint
	NumberInRow uint
}
