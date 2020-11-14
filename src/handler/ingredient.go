package handler

import (
	"github.com/labstack/echo"
)

func ingredientRoutes(r *echo.Echo) {
	g := r.Group("/ingredient")
	{
		g.GET("", getIngredients)
	}
}

func getIngredients(c echo.Context) error {
	ingredients, err := ingredient.GetIngredients(c)
	if err != nil {
		return err
	}
	return JSON(c, 200, ingredients)
}
