package model

// Ingredient data model
type Ingredient struct {
	Base
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}
