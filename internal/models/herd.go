package models

import (
	schema "horsy_api/internal/gen/schema"
	"horsy_api/pkg"
	"time"
)

type Herd struct {
	ID          string
	Name        string
	Description *string
	AccountID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (h *Herd) LoadFromDB(dbHerd schema.Herd) {
	h.ID = dbHerd.ID.String()
	h.Name = dbHerd.Name

	h.Description = pkg.ConvertNullableVarToPtr(dbHerd.Description)

	h.AccountID = dbHerd.AccountID.String()
	h.CreatedAt = dbHerd.CreatedAt
	h.UpdatedAt = dbHerd.UpdatedAt
}
