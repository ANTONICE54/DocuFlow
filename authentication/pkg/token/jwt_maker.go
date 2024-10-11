package token

import (
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
			"user_id": userID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
