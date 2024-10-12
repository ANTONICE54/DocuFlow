package token

import (
	"auth_service/internal/apperrors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secretKey []byte
}

func NewJWTMaker(secret string) *JWTMaker {
	return &JWTMaker{
		secretKey: []byte(secret),
	}
}

func (m *JWTMaker) GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": uint(userID),
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (m *JWTMaker) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return m.secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return apperrors.ErrUnauthorized("invalid token")
	}

	return nil
}

func (m *JWTMaker) ExtractClaims(tokenString string) (*uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return m.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, apperrors.ErrInternalServer("Failed to extract claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)

	if !ok {
		return nil, apperrors.ErrInternalServer("Failed to convert interface to float")
	}

	userIDUint := uint(userIDFloat)

	return &userIDUint, nil

}
