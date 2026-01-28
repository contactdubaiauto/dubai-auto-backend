package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAppleClientSecret generates a JWT client secret for Apple Sign In
// using the .p8 private key file
func GenerateAppleClientSecret(teamID, keyID, clientID, keyPath string) (string, error) {
	// Read the .p8 file
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read Apple key file: %w", err)
	}

	// Parse the PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("failed to decode PEM block")
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// Assert it's an ECDSA private key
	ecdsaPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", errors.New("private key is not ECDSA")
	}

	// Create JWT claims
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": teamID,
		"iat": now.Unix(),
		"exp": now.Add(6 * 30 * 24 * time.Hour).Unix(), // Apple tokens are valid for 6 months
		"aud": "https://appleid.apple.com",
		"sub": clientID,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = keyID

	// Sign token
	tokenString, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
