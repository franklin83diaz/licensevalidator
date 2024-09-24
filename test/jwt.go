package test

import (
	"fmt"
	"licensevalidator/internal/dto"

	"github.com/golang-jwt/jwt/v4"
)

func CheckJWT(publicKeyData string, tokenString string) (dtoReq dto.ServerRequest, err error) {

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyData))
	if err != nil {
		return dtoReq, fmt.Errorf("error parsing the public key: %v", err)
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
		return dtoReq, fmt.Errorf("error parsing the token: %v", err)
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return dtoReq, fmt.Errorf("error: token is not valid")
	}

	// Get values from the token
	protectedID, ok := claims["sub"].(string)
	if !ok {
		return dtoReq, fmt.Errorf("error: protected id not found")
	}

	serialNumber, ok := claims["serialNumber"].(string)
	if !ok {
		return dtoReq, fmt.Errorf("error: serialNumber not found")
	}

	dtoReq = dto.NewServerRequest(serialNumber, protectedID)

	return dtoReq, nil
}
