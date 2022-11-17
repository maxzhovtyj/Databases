package domain

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	CustomerID uint
	SessionID  uint
	RowID      uint
	PositionID uint
	Price      float64
}

type TicketsSearchParams struct {
	PriceGt         float64
	PriceLt         float64
	MovieDurationGt string
	MovieDurationLt string
}
