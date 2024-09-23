package internal

import (
	"licensevalidator/entities"
	"reflect"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	type args struct {
		internalFileKey string
		claimsData      interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJWT(tt.args.internalFileKey, tt.args.claimsData)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckJWT(t *testing.T) {
	token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEyMzQ1Njc4OTAsInByb3RlY3RlZElkIjoiWDEyMyIsImlhdCI6MTUxNjIzOTAyMn0.E47UqLfFISTHkRAqEYd0DPUHNmH0XQY2bkMJVTvoWJkOLwi65CA5ybuCYGJJFiKA_X2LMdscQ4EiM8jlO1zExksWWTWgRu6rNtpM1XSqIpGUbA8t2yzxB88NDb65VTpOFlUN4jZjwHJqCbZ8nj0QvCv_Gyb1ECebsONpTP7bJKn-iIC_nvoBMMYdz8Caxs8cSuqlSKs6_Ozpf9N52fJ4m1DX6DqVt_Q842NOcub223bFjKwtuh2_xsMQhNJ81GUori33O6kFlnSAe5WBSW3ZtWWH0m2F_SEZZSKkpJw5GktZ0rkImXjpMdAoZcpqU1PnL51-J7DIDM_e0c67_wH4Ww`
	publicKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
mwIDAQAB
-----END PUBLIC KEY-----`
	type args struct {
		publicKey   string
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.License
		wantErr bool
	}{
		{"Test 1", args{publicKey, token}, entities.License{
			Sub:         int64(1234567890),
			ProtectedID: "X123",
			Iat:         int64(1516239022),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckJWT(tt.args.publicKey, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Sub, tt.want.Sub) {
				t.Errorf("CheckJWT() = %v, want %v", got.Sub, tt.want.Sub)
			}
			if !reflect.DeepEqual(got.ProtectedID, tt.want.ProtectedID) {
				t.Errorf("CheckJWT() = %v, want %v", got.ProtectedID, tt.want.ProtectedID)
			}

			if !reflect.DeepEqual(got.Iat, tt.want.Iat) {
				t.Errorf("CheckJWT() = %v, want %v", got.Iat, tt.want.Iat)
			}
		})
	}
}
