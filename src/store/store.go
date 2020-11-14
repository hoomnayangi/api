package store

import (
	"github.com/hoomnayangi/api/src/store/ingredient"
	"github.com/hoomnayangi/api/src/store/recipe"
)

// Store - server store struct
type Store struct {
	Ingredient ingredient.Store
	Recipe     recipe.Store
}

// New - create new store variable
func New() *Store {
	return &Store{
		ingredient.NewStore(),
		recipe.NewStore(),
	}
}
