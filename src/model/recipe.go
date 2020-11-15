package model

import (
	"encoding/json"
	"fmt"
)

// Recipe data model
type Recipe struct {
	Base
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Steps       string `json:"-" sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	Picture     string `json:"picture"`
	Category    string `json:"category"`

	RecipesIngredients []*RecipeIngredient `json:"-"`
}

// MarshalJSON - Marshaler implementation of Recipe model
func (r *Recipe) MarshalJSON() ([]byte, error) {
	resp := struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
		Category    string `json:"category"`
	}{r.ID, r.Name, r.Description, r.Picture, r.Category}
	b, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	ingredientsJSON := make([]*IngredientResponse, len(r.RecipesIngredients))
	for i, v := range r.RecipesIngredients {
		ingredientsJSON[i] = &IngredientResponse{Ingredient: v.Ingredient, Amount: v.Amount, Unit: v.Unit}
	}
	ingredientsRaw, err := json.Marshal(ingredientsJSON)
	if err != nil {
		return nil, err
	}
	jsonSteps := fmt.Sprintf(`,"steps": %s, "ingredients": %v }`, r.Steps, string(ingredientsRaw))
	return append(b[:len(b)-1], []byte(jsonSteps)...), nil
}
