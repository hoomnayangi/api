package model

// RecipeIngredient data model
type RecipeIngredient struct {
	Base
	RecipeID     int    `json:"recipe_id"`
	IngredientID int    `json:"ingredient_id"`
	Amount       string `json:"amount"`
	Unit         string `json:"unit"`
}
