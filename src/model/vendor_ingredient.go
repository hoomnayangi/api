package model

// VendorIngredient data model
type VendorIngredient struct {
	Base
	VendorID     int    `json:"recipe_id"`
	IngredientID int    `json:"ingredient_id"`
	Amount       string `json:"amount"`
	Unit         string `json:"unit"`
}
