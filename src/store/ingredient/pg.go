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
	return res, tx.
		Where(fmt.Sprintf("lower(name) like lower('%%%%%s%%%%') and is_hidden <> true", key)).
		Order("NULLIF(picture, '') DESC NULLS LAST").
		Limit(7).
		Find(&res).
		Error
}
