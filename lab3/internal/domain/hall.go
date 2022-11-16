package domain

type Hall struct {
	Id          int
	Title       string
	Description string
	Capacity    *int
}

type SelectHallDTO struct {
	Id          int
	Title       string
	Description string
	Capacity    int
	Rows        []int32
}

type HallsSearchParams struct {
	HallTitle  string
	CapacityGt int
	CapacityLt int
}
