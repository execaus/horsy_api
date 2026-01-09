package dto

import (
	"horsy_api/internal/models"
	"time"
)

type HerdOutput struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	AccountID   string    `json:"accountId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (h *HerdOutput) Load(herd models.Herd) {
	h.ID = herd.ID
	h.Name = herd.Name
	h.Description = herd.Description
	h.AccountID = herd.AccountID
	h.CreatedAt = herd.CreatedAt
	h.UpdatedAt = herd.UpdatedAt
}

type CreateHerdRequest struct {
	Name        string  `json:"name" validate:"required,min=4,max=64"`
	Description *string `json:"description" validate:"max=1024"`
}

type CreateHerdResponse struct {
	Herd HerdOutput `json:"herd"`
}

type GetHerdsResponse struct {
	Herds      []HerdOutput `json:"herds"`
	TotalCount int32        `json:"totalCount"`
}

type GetHerdByIDResponse struct {
	Herd HerdOutput `json:"herd"`
}

type GetHerdHorsesResponse struct {
	Horses     []HorseOutput `json:"horses"`
	TotalCount int32         `json:"totalCount"`
}
