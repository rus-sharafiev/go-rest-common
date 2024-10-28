package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	common "github.com/rus-sharafiev/go-rest-common"
)

type Claims struct {
	UserId     int    `json:"userId"`
	UserAccess string `json:"userAccess"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(id int, userAccess string) (string, error) {
	claims := Claims{
		id,
		userAccess,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(*common.Config.JwtKey))
}

func GenerateRefreshToken(id int, userAccess string) (string, error) {
	claims := Claims{
		id,
		userAccess,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(*common.Config.JwtKey))
}

func Validate(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(*common.Config.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}
