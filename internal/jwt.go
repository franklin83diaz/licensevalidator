package internal

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateJWT creates a JWT token
// receives the name internal file key key and the [claimsData]
//
// [claimsData] is a structure that implements one of the following:
// ServerRequest
func CreateJWT(internalFile string, claimsData interface{}) (string, error) {
	// Read the private key
	privateKeyData, err := ReadInternalfile(internalFile)
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

// CheckJWT verifies a JWT token
// receives the name internal file key key and the token string
// returns the claims if the token is valid
func CheckJWT(internalFile string, tokenString string) (jwt.MapClaims, error) {
	// Read the public key
	publicKeyData, err := ReadInternalfile(internalFile)
	if err != nil {
		return nil, fmt.Errorf("error reading the public key: %v", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyData))
	if err != nil {
		return nil, fmt.Errorf("error parsing the public key: %v", err)
	}

	// check the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//using RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing the token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("error: token is not valid")
}
