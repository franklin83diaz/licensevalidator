package internal

import (
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v4"
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
	token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicHJvdGVjdGVkSWQiOiJYMTIzIiwiaWF0IjoxNTE2MjM5MDIyfQ.SABagqyys2MhgvIUjybOyFkH2XBu-9HxGSDgRgkg3LmDXj94Yn1Eu5q1J-_NwYihEKf8obopNINCp2mtFx4myE6ZXWl6H7gOOjO0Z3D8rb1ngKjQpRt-z36KvaL3twVMGUAHzSBFdZZhQgjPKn004BRyMP83-tRf_NpZRYdwUuTETi7RZFDLgMOKJaLhequO_Dd3my_G59JaYb9eGPnpzPZIKNuYs0O3utYFHd6EfZVzTSXY-wav7IfIIxeKUe693XL8r2M0SKeH9lj1xpTXeZuzIPdGy5rtZYAWuYA3LiDFeIBipB49kLNo2o68eQcROPc5yfDG7X__3k7NHkcRsA`
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
		internalFile string
		tokenString  string
	}
	tests := []struct {
		name    string
		args    args
		want    jwt.MapClaims
		wantErr bool
	}{
		{"Test 1", args{publicKey, token}, jwt.MapClaims{
			"sub":         "1234567890",
			"protectedId": "X123",
			"iat":         1516239022,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckJWT(tt.args.internalFile, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got["sub"], tt.want["sub"]) {
				t.Errorf("CheckJWT() = %v, want %v", got["sub"], tt.want["sub"])
			}
			if !reflect.DeepEqual(got["protectedId"], tt.want["protectedId"]) {
				t.Errorf("CheckJWT() = %v, want %v", got["protectedId"], tt.want["protectedId"])
			}
			//DOTO: Fix this test
			// if !reflect.DeepEqual(got["iat"], tt.want["iat"]) {
			// 	t.Errorf("CheckJWT() = %v, want %v", got["iat"], tt.want["iat"])
			// }
		})
	}
}
