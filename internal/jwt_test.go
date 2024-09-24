package internal

import (
	"licensevalidator/internal/entities"
	"reflect"
	"testing"
)

func TestCreateJWT(t *testing.T) {

	privateKey := `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC7VJTUt9Us8cKj
MzEfYyjiWA4R4/M2bS1GB4t7NXp98C3SC6dVMvDuictGeurT8jNbvJZHtCSuYEvu
NMoSfm76oqFvAp8Gy0iz5sxjZmSnXyCdPEovGhLa0VzMaQ8s+CLOyS56YyCFGeJZ
qgtzJ6GR3eqoYSW9b9UMvkBpZODSctWSNGj3P7jRFDO5VoTwCQAWbFnOjDfH5Ulg
p2PKSQnSJP3AJLQNFNe7br1XbrhV//eO+t51mIpGSDCUv3E0DDFcWDTH9cXDTTlR
ZVEiR2BwpZOOkE/Z0/BVnhZYL71oZV34bKfWjQIt6V/isSMahdsAASACp4ZTGtwi
VuNd9tybAgMBAAECggEBAKTmjaS6tkK8BlPXClTQ2vpz/N6uxDeS35mXpqasqskV
laAidgg/sWqpjXDbXr93otIMLlWsM+X0CqMDgSXKejLS2jx4GDjI1ZTXg++0AMJ8
sJ74pWzVDOfmCEQ/7wXs3+cbnXhKriO8Z036q92Qc1+N87SI38nkGa0ABH9CN83H
mQqt4fB7UdHzuIRe/me2PGhIq5ZBzj6h3BpoPGzEP+x3l9YmK8t/1cN0pqI+dQwY
dgfGjackLu/2qH80MCF7IyQaseZUOJyKrCLtSD/Iixv/hzDEUPfOCjFDgTpzf3cw
ta8+oE4wHCo1iI1/4TlPkwmXx4qSXtmw4aQPz7IDQvECgYEA8KNThCO2gsC2I9PQ
DM/8Cw0O983WCDY+oi+7JPiNAJwv5DYBqEZB1QYdj06YD16XlC/HAZMsMku1na2T
N0driwenQQWzoev3g2S7gRDoS/FCJSI3jJ+kjgtaA7Qmzlgk1TxODN+G1H91HW7t
0l7VnL27IWyYo2qRRK3jzxqUiPUCgYEAx0oQs2reBQGMVZnApD1jeq7n4MvNLcPv
t8b/eU9iUv6Y4Mj0Suo/AU8lYZXm8ubbqAlwz2VSVunD2tOplHyMUrtCtObAfVDU
AhCndKaA9gApgfb3xw1IKbuQ1u4IF1FJl3VtumfQn//LiH1B3rXhcdyo3/vIttEk
48RakUKClU8CgYEAzV7W3COOlDDcQd935DdtKBFRAPRPAlspQUnzMi5eSHMD/ISL
DY5IiQHbIH83D4bvXq0X7qQoSBSNP7Dvv3HYuqMhf0DaegrlBuJllFVVq9qPVRnK
xt1Il2HgxOBvbhOT+9in1BzA+YJ99UzC85O0Qz06A+CmtHEy4aZ2kj5hHjECgYEA
mNS4+A8Fkss8Js1RieK2LniBxMgmYml3pfVLKGnzmng7H2+cwPLhPIzIuwytXywh
2bzbsYEfYx3EoEVgMEpPhoarQnYPukrJO4gwE2o5Te6T5mJSZGlQJQj9q4ZB2Dfz
et6INsK0oG8XVGXSpQvQh3RUYekCZQkBBFcpqWpbIEsCgYAnM3DQf3FJoSnXaMhr
VBIovic5l0xFkEHskAjFTevO86Fsz1C2aSeRKSqGFoOQ0tmJzBEs1R6KqnHInicD
TQrKhArgLXX4v3CddjfTRJkFWDbE/CkvKZNOrcf1nhaGCPspRJj2KUkj1Fhl9Cnc
dn/RsYEONbwQSjIfMPkvxF+8HQ==
-----END PRIVATE KEY-----`

	serialNumber := "1234567890"
	protectedId := "X123"
	want := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJTZXJpYWxOdW1iZXIiOiIxMjM0NTY3ODkwIiwiZXhwIjoxNzI3MjA5NzMyLCJwcm90ZWN0ZWRJZCI6IlgxMjMifQ.QDu8PwTqW8c8k6PUdfw1Yj0pWZ4oq2EGFMFh-T1SfdIipW4JyYIlxK0_T5M_7asf99A4LUq97KGzusa6HLP4hcjSVOUAj4CM6mITWWfVYH0dtYa9a68tGxS18CLvSjLXDoW3mjSwb5hMmHQBekQoL4QOwX0GixCyOZgmH1erA1jXHcScSse561beQtHid1Pxc3_k95_ZVNb5fdFU8WKY-WZTTOnAPYKpc-FSn0ZDRRE1mU3TB0vAQGcCWGXbYXtlt0XwHupGCY9CRvQn3K7JatnvWUrPXSIG1x3SXxmNEwVOfrQmlBvvtqZdTiPwXweOYTnuWGduogOD3pYxLg7vHw`

	type args struct {
		privateKey   string
		serialNumber string
		protectedId  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{" Create JWT Test 1", args{privateKey: privateKey, serialNumber: serialNumber, protectedId: protectedId}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateJWT(tt.args.privateKey, tt.args.serialNumber, tt.args.protectedId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestCheckJWT(t *testing.T) {
	token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxNTAxZDkwNTJlMDA1OTc3YTNhMTI4MWY4ZjRhNjg5NjgxZjgyZGEwOTIyOTAwNTg4OWUwZmNjNWNmMzY1ZTU3IiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE5MTYyMzkwMjJ9.ZFRLTiOH-Ndl3U5tk6SLsOkgotv6MLw_z8J7VaJfGMqvotLC4YlG52dAizySDhhPmTsMWi_G8w8HM7fp8IjUQHCOy3EihcmXumJp06Y0k7LgqjJFLDR3S8Pmt0hOTmGB3AnNyyKHpw8g2P7w9-IjiuffuIzP2JSn-nmc_mq8q8qpaD-M5szTpPmKyvDWq1oek1ivFpHLgLDM4Oci5YxSZkan8sERZsMEPnCELq3DAub5t6-MZLWxEaTmZQtR-nQUzDh15afe5CLRFw4jhPkI8DwHGHBew_3fVVJGMziUTyeAxrS08AkhSVfsWsWlbKz4j8lu2P-hhSByonBma1unEw`
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
		{"Check JWT Test 1", args{publicKey, token}, entities.License{
			ProtectedID: "1501d9052e005977a3a1281f8f4a689681f82da09229005889e0fcc5cf365e57",
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

			if !reflect.DeepEqual(got.ProtectedID, tt.want.ProtectedID) {
				t.Errorf("CheckJWT() = %v, want %v", got.ProtectedID, tt.want.ProtectedID)
			}

			if !reflect.DeepEqual(got.Iat, tt.want.Iat) {
				t.Errorf("CheckJWT() = %v, want %v", got.Iat, tt.want.Iat)
			}
		})
	}
}
