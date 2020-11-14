package handler

import (
	"strings"

	"github.com/hoomnayangi/api/src/handler/recipe"
	"github.com/labstack/echo"
)

func recipeRoutes(r *echo.Echo) {
	g := r.Group("/recipe")
	{
		g.GET("", getFromIngredients)
	}
}

func getFromIngredients(c echo.Context) error {
	ingredientsParam := strings.ReplaceAll(c.QueryParam("ingredients"), " ", "")
	if ingredientsParam == "" {
		return JSON(c, 400, "fill up ingredients to find your perfect recipe")
	}
	ingredients := strings.Split(ingredientsParam, ",")
	if len(ingredients) == 0 {
		return JSON(c, 400, "fill up ingredients to find your perfect recipe")
	}
	recipes, err := recipe.GetFromIngredients(c, ingredients)
	if err != nil {
		return err
	}
	return JSON(c, 200, recipes)
}
