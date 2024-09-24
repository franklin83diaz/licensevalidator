package entities

type License struct {
	SerialNumber string //This is the serial number
	ProtectedID  string //This is the protected id is sub in the jwt
	Iat          int64  //This is the time the token was issued
	Exp          int64  //This is the expiration time for the license
}
