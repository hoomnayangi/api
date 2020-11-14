package recipe

import (
	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"

	"github.com/hoomnayangi/api/src/model"
	"github.com/hoomnayangi/api/src/server"
)

// GetHighLightRecipes - get high light recipe for main page
func GetHighLightRecipes(c echo.Context, lat, long float64) ([]*model.Recipe, error) {
	srv := server.GetServerCfg()

	res, err := srv.Store().Recipe.GetHighLight(c, lat, long)
	if err != nil {
		return nil, errors.Customize(err, 404, "recipe not found")
	}

	return res, nil
}

// GetFromIngredients - get recipe from provided ingredients
func GetFromIngredients(c echo.Context, ingredients []string) ([]*model.Recipe, error) {
	srv := server.GetServerCfg()
	res, err := srv.Store().Recipe.GetFromIngredients(c, ingredients)
	if err != nil {
		return nil, errors.Customize(err, 404, "recipe not found")
	}

	return res, nil
}
