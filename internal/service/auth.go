package service

import (
	"context"
	"crypto/rand"
	"errors"
	"horsy_api/config"
	"horsy_api/internal/models"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultJWTExpireDuration = time.Hour * 24
	passwordLength           = 16
	chars                    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type AuthService struct {
	secretKey string
}

func (s *AuthService) GenerateToken(ctx context.Context, accountID uuid.UUID) (string, error) {
	claims := models.AuthClaims{
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(DefaultJWTExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) GetClaimsFromToken(ctx context.Context, tokenString string) (*models.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			zap.L().Error(jwt.ErrTokenSignatureInvalid.Error())
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			zap.L().Warn(err.Error())
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(*models.AuthClaims); ok && token.Valid {
		return claims, nil
	}

	zap.L().Error(ErrTokenInvalid.Error())
	return nil, ErrTokenInvalid
}

func (s *AuthService) ComparePassword(ctx context.Context, hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		zap.L().Error(err.Error())
		return false
	}

	return true
}

func (s *AuthService) hashPassword(ctx context.Context, password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}
	return string(hashedBytes), nil
}

func (s *AuthService) generatePassword() (string, error) {
	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		indexBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			zap.L().Error(err.Error())
			return "", err
		}
		password[i] = chars[indexBig.Int64()]
	}

	return string(password), nil
}

func NewAuthService(cfg config.AuthConfig) *AuthService {
	return &AuthService{secretKey: cfg.Key}
}
