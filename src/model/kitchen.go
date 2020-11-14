package model

// Kitchen data model
type Kitchen struct {
	Base
	ID          int     `json:"id"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Description string  `json:"description"`
}
