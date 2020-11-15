package model

// OrderRecipe data model
type OrderRecipe struct {
	OrderID  int    `json:"order_id"`
	RecipeID int    `json:"recipe_id"`
	Amount   string `json:"amount"`
}
