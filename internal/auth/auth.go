package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//TODO: fix this
var jwtSecret = []byte("supersecretkey")

type Claims struct {
	UserID string `json:"userid"`
	Token  string `json:"token"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, token string, duration time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		Token:  token,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
