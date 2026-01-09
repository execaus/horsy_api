package service

import "github.com/google/uuid"

type CreateHorseBreedParams struct {
	ID      uuid.UUID
	Percent int32
}
