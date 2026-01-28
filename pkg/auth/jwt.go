package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(
	id int, expiration time.Duration, secret_key string, role_id int,
) string {
	unixTime := time.Now().Add(expiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"role_id": role_id,
		"exp":     unixTime,
	})

	tokenString, _ := token.SignedString([]byte(secret_key))

	return tokenString
}

func CreateRefreshAccsessToken(id, role_id int) (string, string) {

	accessToken := CreateToken(id, ENV.REFRESH_TIME, ENV.ACCESS_KEY, role_id)
	refreshToken := CreateToken(id, ENV.REFRESH_TIME, ENV.REFRESH_KEY, role_id)

	return accessToken, refreshToken
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func VerifyAppleIDToken(tokenString string) (jwt.MapClaims, error) {
	keySet, err := jwk.Fetch(context.Background(), "https://appleid.apple.com/auth/keys")
	if err != nil {
		return nil, errors.New("failed to fetch apple public keys")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, errors.New("missing kid")
		}

		key, found := keySet.LookupKeyID(kid)
		if !found {
			return nil, errors.New("invalid key id")
		}

		var pubKey interface{}
		if err := key.Raw(&pubKey); err != nil {
			return nil, err
		}

		return pubKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid apple token")
	}

	claims := token.Claims.(jwt.MapClaims)

	// Optional checks
	if claims["iss"] != "https://appleid.apple.com" {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}
