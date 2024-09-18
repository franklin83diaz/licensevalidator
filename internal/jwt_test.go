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
		{"Test 1", args{"private_key_app", token}, jwt.MapClaims{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckJWT(tt.args.internalFile, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
