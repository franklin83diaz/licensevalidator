package internal

import (
	"os"
	"time"
)

func CheckLicenseFromFile(protectedId string, pathLicenseFile string, ServerPubKey string) (bool, error) {
	// Open the file
	file, err := os.Open(pathLicenseFile)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the file
	buffer := make([]byte, 1024)
	len, err := file.Read(buffer)
	if err != nil {
		return false, err
	}

	lic, err := CheckJWT(ServerPubKey, string(buffer[:len]))
	if err != nil {
		return false, err
	}

	//check if the protected id is the same
	if lic.ProtectedID != protectedId {
		return false, nil
	}

	// Check if the token is expired
	timeNow := time.Now().Unix()
	if lic.Iat > timeNow {
		return false, nil
	}

	//check if sub is more a current time (not implemented for now)
	// if lic.Sub < timeNow {
	// 	return false, nil
	// }

	return true, nil

}
