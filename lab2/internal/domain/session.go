package domain

import "time"

type Session struct {
	Id      int
	MovieId int
	HallId  int
	StartAt time.Time
}
