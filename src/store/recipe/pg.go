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

func (s recipePGStore) GetHighLight(c echo.Context, lat, long float64) (*model.Recipe, error) {
	tx := db.GetTxFromCtx(c)
	var res model.Recipe
	return &res, tx.Where("id = ?", id).First(&res).Error
}
