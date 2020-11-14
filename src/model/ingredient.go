package model

// Ingredient data model
type Ingredient struct {
	Base
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

// IngredientResponse - JSON response model
type IngredientResponse struct {
	*Ingredient
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
}
