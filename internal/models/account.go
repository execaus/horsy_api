package models

import (
	"horsy_api/internal/gen/schema"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             uuid.UUID
	Email          string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastActivityAt time.Time
}

func (a *Account) LoadFromDB(from schema.Account) error {
	a.ID = from.ID
	a.Email = from.Email
	a.Password = from.Password
	a.CreatedAt = from.CreatedAt
	a.UpdatedAt = from.UpdatedAt
	a.LastActivityAt = from.LastActivityAt

	return nil
}
