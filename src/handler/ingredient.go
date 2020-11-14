package handler

import (
	"github.com/labstack/echo"

	"github.com/hoomnayangi/api/src/handler/ingredient"
)

func ingredientRoutes(r *echo.Echo) {
	g := r.Group("/ingredient")
	{
		g.GET("", searchIngredients)
	}
}

func searchIngredients(c echo.Context) error {
	key := c.QueryParam("key")
	ingredients, err := ingredient.SearchIngredients(c, key)
	if err != nil {
		return err
	}
	return JSON(c, 200, ingredients)
}
