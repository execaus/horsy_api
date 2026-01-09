package models

import (
	"horsy_api/internal/gen/schema"
	"horsy_api/pkg"
)

type Breed struct {
	ID          string
	Name        string
	Description *string
}

func (h *Breed) LoadFromDB(dbHorseBreed schema.Breed) *Breed {
	h.ID = dbHorseBreed.ID.String()
	h.Name = dbHorseBreed.Name
	h.Description = pkg.ConvertNullableVarToPtr(dbHorseBreed.Description)

	return h
}
