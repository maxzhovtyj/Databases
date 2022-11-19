package domain

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	CustomerID uint `gorm:"default:null"`
	SessionID  uint
	RowID      uint `gorm:"constraint:OnDelete:SET NULL;"`
	PositionID uint `gorm:"constraint:OnDelete:SET NULL;"`
	Price      float64
}

type TicketsSearchParams struct {
	PriceGt         float64
	PriceLt         float64
	MovieDurationGt string
	MovieDurationLt string
}
