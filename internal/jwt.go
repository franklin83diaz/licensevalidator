package internal

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateJWT creates a JWT token
// receives the path to the private key and the [claimsData]
//
// [claimsData] is a structure that implements one of the following:
// ServerRequest
func CreateJWT(privateKeyPath string, claimsData interface{}) (string, error) {
	// Read the private key
	privateKeyData, err := ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("error reading the private key: %v", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyData))
	if err != nil {
		return "", fmt.Errorf("error parsing the private key: %v", err)
	}

	// Implement the ServerRequest interface
	claimsDataImp, ok := claimsData.(ServerRequest)
	if !ok {
		//TODO: Implement other interfaces

		//For now, return an error
		return "", fmt.Errorf("error: claimsData does not implement any of the interfaces")
	}

	// claims
	claims := jwt.MapClaims{
		"SerialNumber": claimsDataImp.GetSerialNumber(),
		"protectedID":  claimsDataImp.GetProtectedID(),
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	// create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing the token: %v", err)
	}

	return tokenString, nil
}
