package domain

import "time"

type Ticket struct {
	Id         int
	CustomerId int
	SessionId  int
	Price      float64
	RowId      int
	PositionId int
}

type SelectTicketDTO struct {
	Id                int
	MovieTitle        string
	SessionStartAt    time.Time
	HallTitle         string
	CustomerFirstname string
	CustomerLastname  string
	Price             float64
	MovieDuration     time.Duration
	Row               int
	Position          int
}

type TicketsSearchParams struct {
	PriceGt         float64
	PriceLt         float64
	MovieDurationGt string
	MovieDurationLt string
}
