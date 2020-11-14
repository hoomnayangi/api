package ingredient

import (
	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"

	"github.com/hoomnayangi/api/src/model"
	"github.com/hoomnayangi/api/src/server"
)

// SearchIngredients - get ingredients handler
func SearchIngredients(c echo.Context, key string) ([]*model.Ingredient, error) {
	srv := server.GetServerCfg()
	res, err := srv.Store().Ingredient.Search(c, key)
	if err != nil {
		return nil, errors.Customize(err, 404, "ingredient not found")
	}

	return res, nil
}
