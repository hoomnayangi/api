package ingredient

import (
	"github.com/labstack/echo"

	"github.com/hoomnayangi/api/src/model"
)

// Store - ingredient store interface
type Store interface {
	Get(c echo.Context, id int) (*model.Ingredient, error)
}
