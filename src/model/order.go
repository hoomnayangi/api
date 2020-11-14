package model

// Order data model
type Order struct {
	Base
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Description string  `json:"description"`
}
