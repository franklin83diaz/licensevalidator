package internal

import (
	"fmt"
	"time"

	"licensevalidator/internal/entities"

	"github.com/golang-jwt/jwt/v4"
)

// CreateJWT creates a JWT token
// receives the string private key and the [claimsData]
//
// [claimsData] is a structure that implements one of the following:
// ServerRequest
func CreateJWT(privateKeyData string, claimsData interface{}) (string, error) {

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
		"protectedId":  claimsDataImp.GetProtectedId(),
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
// receives the string public key and the token string
// returns the claims if the token is valid
func CheckJWT(publicKeyData string, tokenString string) (lic entities.License, err error) {

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyData))
	if err != nil {
		return lic, fmt.Errorf("error parsing the public key: %v", err)
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
		return lic, fmt.Errorf("error parsing the token: %v", err)
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return lic, fmt.Errorf("error: token is not valid")
	}

	lic = entities.License{
		Sub:         int64(claims["sub"].(float64)),
		ProtectedID: claims["protectedId"].(string),
		Iat:         int64(claims["iat"].(float64)),
	}

	return lic, nil
}
