package dto

import "horsy_api/internal/models"

type BirthplaceOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (b *BirthplaceOutput) Load(birthplace models.Birthplace) {
	b.ID = birthplace.ID
	b.Name = birthplace.Name
	b.Description = birthplace.Description
}

type GetBirthplacesResponse struct {
	Birthplaces []*BirthplaceOutput `json:"birthplaces"`
}
