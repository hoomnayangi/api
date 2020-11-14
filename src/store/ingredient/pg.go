package ingredient

import (
	"github.com/labstack/echo"
	"github.com/phuwn/tools/db"

	"github.com/hoomnayangi/api/src/model"
)

type ingredientPGStore struct{}

// NewStore - create new ingredient store
func NewStore() Store {
	return &ingredientPGStore{}
}

func (s ingredientPGStore) Get(c echo.Context, id int) (*model.Ingredient, error) {
	tx := db.GetTxFromCtx(c)
	var res model.Ingredient
	return &res, tx.Where("id = ?", id).First(&res).Error
}
