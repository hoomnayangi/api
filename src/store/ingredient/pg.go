package ingredient

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/phuwn/tools/db"

	"github.com/hoomnayangi/api/src/model"
)

type ingredientPGStore struct{}

// NewStore - create new ingredient store
func NewStore() Store {
	return &ingredientPGStore{}
}

func (s ingredientPGStore) Search(c echo.Context, key string) ([]*model.Ingredient, error) {
	tx := db.GetTxFromCtx(c)
	var res []*model.Ingredient
	return res, tx.Where("name like ?", fmt.Sprintf("%%%s%%", key)).Limit(7).Find(&res).Error
}
