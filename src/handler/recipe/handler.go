package recipe

import (
	"github.com/labstack/echo"

	"github.com/hoomnayangi/api/src/model"
	"github.com/hoomnayangi/api/src/server"
)

// GetHighLightRecipes - get high light recipe for main page
func GetHighLightRecipes(c echo.Context, lat, long float64) ([]*model.Recipe, error) {
	srv := server.GetServerCfg()

	return nil, nil
}
