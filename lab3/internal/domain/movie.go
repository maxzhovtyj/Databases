package domain

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Sessions    []Session
	Title       string
	Description string
	Duration    string
}
