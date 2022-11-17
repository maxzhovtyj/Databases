package domain

import "gorm.io/gorm"

type Hall struct {
	gorm.Model
	Sessions    []Session
	Title       string
	Description string
	Rows        []Row
	Capacity    *uint
}

type HallsSearchParams struct {
	HallTitle  string
	CapacityGt int
	CapacityLt int
}
