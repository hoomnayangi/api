package model

// RecipeIngredient data model
type RecipeIngredient struct {
	RecipeID     int    `json:"recipe_id"`
	IngredientID int    `json:"ingredient_id"`
	Amount       string `json:"amount"`
	Unit         string `json:"unit"`

	Ingredient *Ingredient `json:"ingredient"`
}

func (ri RecipeIngredient) TableName() string {
	return "recipes_ingredients"
}
