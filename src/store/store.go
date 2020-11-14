package store

import "github.com/hoomnayangi/api/src/store/ingredient"

// Store - server store struct
type Store struct {
	Ingredient ingredient.Store
}

// New - create new store variable
func New() *Store {
	return &Store{
		ingredient.NewStore()
	}
}
