package models

import (
	"horsy_api/internal/gen/schema"
	"horsy_api/pkg"
)

type Birthplace struct {
	ID          string
	Name        string
	Description *string
}

func (h *Birthplace) LoadFromDB(dbHorseBirthplace schema.Birthplace) *Birthplace {
	h.ID = dbHorseBirthplace.ID.String()
	h.Name = dbHorseBirthplace.Name
	h.Description = pkg.ConvertNullableVarToPtr(dbHorseBirthplace.Description)

	return h
}
