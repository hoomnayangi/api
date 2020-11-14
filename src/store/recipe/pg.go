package recipe

import (
	"fmt"

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
		Table("recipes").
		Joins("inner join recipes_ingredients on recipes_ingredients.recipe_id = recipes.id").
		Where("recipes_ingredients.ingredient_id in (?)", ingredients).
		Limit(5).
		Find(&res).
		Error
}

func (s recipePGStore) GetHighLight(c echo.Context, lat, long float64) ([]*model.Recipe, error) {
	tx := db.GetTxFromCtx(c)
	var res []*model.Recipe
	return res, tx.
		Table("recipes").
		Joins("inner join recipes_ingredients on recipes_ingredients.recipe_id = recipes.id").
		Joins("inner join ingredients on recipes_ingredients.ingredient_id = ingredients.id").
		Joins("inner join vendors_ingredients on vendors_ingredients.ingredient_id = ingredients.id").
		Where(fmt.Sprintf("vendors_ingredients.vendor_id = (select id from vendors order by distance(%v, %v, lat, long) desc limit 1)", lat, long)).
		Group("recipes.id, vendors_ingredients.vendor_id").
		Limit(10).
		Find(&res).
		Error
}
