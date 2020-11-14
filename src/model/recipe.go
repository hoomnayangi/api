package model

import (
	"encoding/json"
	"fmt"
)

// Recipe data model
type Recipe struct {
	Base
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Steps       string `json:"-" sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// MarshalJSON - Marshaler implementation of Recipe model
func (r *Recipe) MarshalJSON() ([]byte, error) {
	resp := struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}{r.ID, r.Name, r.Description}
	b, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	jsonSteps := fmt.Sprintf(`,"steps": %s }`, r.Steps)
	return append(b[:len(b)-1], []byte(jsonSteps)...), nil
}
