package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthClaims struct {
	AccountID uuid.UUID
	jwt.RegisteredClaims
}
