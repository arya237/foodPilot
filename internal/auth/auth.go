package auth

import (
	"errors"
	"time"

	"github.com/arya237/foodPilot/internal/config"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET", "I just want you to know that i love you"))

type Claims struct {
	UserID string          `json:"userid"`
	Token  string          `json:"token"`
	Role   models.UserRole `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, token string, role models.UserRole, duration time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		Token:  token,
		Role:   role,
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
