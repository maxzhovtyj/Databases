package domain

import "time"

type Movie struct {
	Id          int
	Title       string
	Description string
	Duration    time.Duration
}
