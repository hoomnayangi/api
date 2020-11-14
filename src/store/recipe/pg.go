package recipe

import (
	"github.com/hoomnayangi/api/src/model"
	"github.com/labstack/echo"
	"github.com/phuwn/tools/db"
)

type recipePGStore struct{}

// NewStore - create new recipe store
func NewStore() Store {
	return &recipePGStore{}
}

func (s recipePGStore) GetFromIngredients(c echo.Context, ingredients []string) ([]*model.Recipe, error) {
	tx := db.GetTxFromCtx(c)
	var res []*model.Recipe
	return res, tx.
		Joins("inner join recipes_ingredients on recipes_ingredients.recipe_id = recipes.id").
		// Preload("")
		Where("recipes_ingredients.ingredient_id in (?)", ingredients).
		Find(&res).
		Error
}
