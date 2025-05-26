package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(
	id int, expiration time.Duration, secret_key, role string,
) string {
	unixTime := time.Now().Add(expiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  unixTime,
	})

	tokenString, _ := token.SignedString([]byte(secret_key))

	return tokenString
}

func CreateRefreshAccsessToken(id int, role string) (string, string) {

	accessToken := CreateToken(id, ENV.REFRESH_TIME, ENV.ACCESS_KEY, role)
	refreshToken := CreateToken(id, ENV.REFRESH_TIME, ENV.REFRESH_KEY, role)

	return accessToken, refreshToken
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
