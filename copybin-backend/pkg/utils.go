package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func TimeNowInMinutes() int {
	return int(time.Now().Unix() / 60)
}

func CreateToken(userID uint, jwtSecret string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			Issuer:    "copybin",
			Subject:   strconv.FormatUint(uint64(userID), 10),
		},
	)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(token string, jwtSecret string) (*jwt.Token, error) {
	return jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
