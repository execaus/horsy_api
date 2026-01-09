package models

import (
	schema "horsy_api/internal/gen/schema"
	"horsy_api/pkg"
	"time"
)

type Horse struct {
	ID            string
	Herd          Herd
	Gender        *Gender
	Name          *string
	BirthDay      *int32
	BirthMonth    *int32
	BirthYear     *int32
	BirthPlace    *Birthplace
	WithersHeight *int32
	Sire          *Horse
	Dam           *Horse
	IsPregnant    bool
	IsDead        bool
	Description   *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (h *Horse) Load(dbHorse schema.Horse) *Horse {
	h.ID = dbHorse.ID.String()
	h.IsPregnant = dbHorse.IsPregnant
	h.CreatedAt = dbHorse.CreatedAt
	h.UpdatedAt = dbHorse.UpdatedAt
	h.IsDead = dbHorse.IsDead

	h.Description = pkg.ConvertNullableVarToPtr(dbHorse.Description)
	h.Name = pkg.ConvertNullableVarToPtr(dbHorse.Name)
	h.BirthDay = pkg.ConvertNullableVarToPtr(dbHorse.BirthDay)
	h.BirthMonth = pkg.ConvertNullableVarToPtr(dbHorse.BirthMonth)
	h.BirthYear = pkg.ConvertNullableVarToPtr(dbHorse.BirthYear)
	h.WithersHeight = pkg.ConvertNullableVarToPtr(dbHorse.WithersHeight)

	return h
}
