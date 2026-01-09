package models

import (
	"horsy_api/internal/gen/schema"
	"horsy_api/pkg"
)

type GeneticMarker struct {
	ID          string
	Name        string
	Description *string
}

func (h *GeneticMarker) LoadFromDB(dbGeneticMarker schema.GeneticMarker) *GeneticMarker {
	h.ID = dbGeneticMarker.ID.String()
	h.Name = dbGeneticMarker.Name

	h.Description = pkg.ConvertNullableVarToPtr(dbGeneticMarker.Description)

	return h
}
