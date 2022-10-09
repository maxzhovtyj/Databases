package domain

import "time"

type Session struct {
	Id      int
	MovieId int
	HallId  int
	StartAt time.Time
}

type SelectSessionDTO struct {
	Id      int
	Movie   string
	StartAt time.Time
	Hall    string
}

type SessionsSearchParams struct {
	MovieName string
	StartAtGt time.Time
	StartAtLt time.Time
}
