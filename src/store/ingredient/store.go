package ingredient

import (
	"github.com/labstack/echo"

	"github.com/hoomnayangi/api/src/model"
)

// Store - ingredient store interface
type Store interface {
	Search(c echo.Context, key string) ([]*model.Ingredient, error)
}
