package recipe

import (
	"github.com/labstack/echo"

	"github.com/hoomnayangi/api/src/model"
)

// Store - recipe store interface
type Store interface {
	GetFromIngredients(c echo.Context, ingredients []string) ([]*model.Recipe, error)
}
