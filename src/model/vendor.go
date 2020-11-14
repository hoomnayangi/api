package model

// Vendor data model
type Vendor struct {
	Base
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	District string  `json:"district"`
	Ward     string  `json:"ward"`
}
