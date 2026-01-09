package models

import (
	"horsy_api/internal/gen/schema"
	"horsy_api/pkg"
)

type Color struct {
	ID          string
	Name        string
	Description *string
}

func (h *Color) LoadFromDB(dbHorseColor schema.Color) *Color {
	h.ID = dbHorseColor.ID.String()
	h.Name = dbHorseColor.Name

	h.Description = pkg.ConvertNullableVarToPtr(dbHorseColor.Description)

	return h
}
