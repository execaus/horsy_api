package models

import (
	"horsy_api/internal/gen/schema"
	"horsy_api/pkg"
)

type Gender struct {
	ID          string
	Name        string
	Description *string
}

func (h *Gender) LoadFromDB(dbHorseGender schema.Gender) *Gender {
	h.ID = dbHorseGender.ID.String()
	h.Name = dbHorseGender.Name

	h.Description = pkg.ConvertNullableVarToPtr(dbHorseGender.Description)

	return h
}
