package model

// Recipe data model
type Recipe struct {
	Base
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Steps       string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}
