package ingredient

import (
	"github.com/labstack/echo"

	"github.com/hoomnayangi/api/src/model"
	"github.com/hoomnayangi/api/src/server"
)

// GetIngredients - get ingredients handler
func GetIngredients(c echo.Context) ([]*model.Ingredient, error) {
	srv := server.GetServerCfg()
	return nil, nil
}
