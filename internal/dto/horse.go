package dto

import (
	"horsy_api/internal/models"
	"time"

	"github.com/google/uuid"
)

type HorseOutput struct {
	ID            string            `json:"id"`
	Herd          HerdOutput        `json:"herd"`
	Gender        *GenderOutput     `json:"gender"`
	Name          *string           `json:"name"`
	Sire          *HorseOutput      `json:"sire"`
	Dam           *HorseOutput      `json:"dam"`
	BirthDay      *int32            `json:"birthDay"`
	BirthMonth    *int32            `json:"birthMonth"`
	BirthYear     *int32            `json:"birthYear"`
	BirthPlace    *BirthplaceOutput `json:"birthPlace"`
	WithersHeight *int32            `json:"withersHeight"`
	IsPregnant    bool              `json:"isPregnant"`
	Description   *string           `json:"description"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

func (o *HorseOutput) Load(horse *models.Horse) {
	var gender *GenderOutput
	if horse.Gender != nil {
		gender = &GenderOutput{}
		gender.Load(*horse.Gender)
	}

	var birthplace *BirthplaceOutput
	if horse.BirthPlace != nil {
		birthplace = &BirthplaceOutput{}
		birthplace.Load(*horse.BirthPlace)
	}

	herd := HerdOutput{}
	herd.Load(horse.Herd)

	if horse.Sire != nil {
		sire := &HorseOutput{}
		sire.Load(horse.Sire)
		o.Sire = sire
	}

	if horse.Dam != nil {
		dam := &HorseOutput{}
		dam.Load(horse.Dam)
		o.Dam = dam
	}

	o.ID = horse.ID
	o.Herd = herd
	o.Gender = gender
	o.Name = horse.Name
	o.BirthDay = horse.BirthDay
	o.BirthMonth = horse.BirthMonth
	o.BirthYear = horse.BirthYear
	o.BirthPlace = birthplace
	o.WithersHeight = horse.WithersHeight
	o.IsPregnant = horse.IsPregnant
	o.Description = horse.Description
	o.CreatedAt = horse.CreatedAt
	o.UpdatedAt = horse.UpdatedAt
}

type HorseBreed struct {
	ID      uuid.UUID `json:"id" binding:"required,uuid4"`
	Percent int32     `json:"percent" binding:"required,min=0,max=10000"`
}

type CreateHorseRequest struct {
	Herd           uuid.UUID    `json:"herd" binding:"required,uuid4"`
	Gender         *uuid.UUID   `json:"gender"`
	Name           *string      `json:"name"`
	BirthDay       *int32       `json:"birthDay"`
	BirthMonth     *int32       `json:"birthMonth"`
	BirthYear      *int32       `json:"birthYear"`
	BirthPlace     *uuid.UUID   `json:"birthPlace"`
	WithersHeight  *int32       `json:"withersHeight"`
	Sire           *uuid.UUID   `json:"sire"`
	Dam            *uuid.UUID   `json:"dam"`
	IsPregnant     bool         `json:"isPregnant"`
	Description    *string      `json:"description"`
	Breeds         []HorseBreed `json:"breeds" binding:"required"`
	Color          uuid.UUID    `json:"color" binding:"required,uuid4"`
	GeneticMarkers []uuid.UUID  `json:"geneticMarkers" binding:"required"`
}

type CreateHorseResponse struct {
	Horse HorseOutput `json:"horse"`
}

type GetHorseResponse struct {
	Horse          HorseOutput   `json:"horse"`
	RelativeHorses []HorseOutput `json:"relativeHorses"`
}
