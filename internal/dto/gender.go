package dto

import "horsy_api/internal/models"

type GenderOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (g *GenderOutput) Load(gender models.Gender) {
	g.ID = gender.ID
	g.Name = gender.Name
	g.Description = gender.Description
}

type GetHorseGenderListResponse struct {
	Genders []*GenderOutput `json:"genders"`
}
