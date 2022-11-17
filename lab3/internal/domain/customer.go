package domain

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Tickets   []Ticket
	FirstName string
	LastName  string
}
