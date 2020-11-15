package handler

import (
	"strconv"
	"strings"

	"github.com/hoomnayangi/api/src/handler/recipe"
	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"
)

func recipeRoutes(r *echo.Echo) {
	g := r.Group("/recipe")
	{
		g.GET("", getFromIngredients)
		g.GET("/highlight", getHighLight)
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

func getHighLight(c echo.Context) error {
	lat, err := strconv.ParseFloat(c.QueryParam("lat"), 64)
	if err != nil {
		return errors.Customize(err, 400, "we need your latitude")
	}
	long, err := strconv.ParseFloat(c.QueryParam("long"), 64)
	if err != nil {
		return errors.Customize(err, 400, "we need your longitude")
	}
	recipes, err := recipe.GetHighLightRecipes(c, lat, long)
	if err != nil {
		return err
	}
	return JSON(c, 200, recipes)
}
