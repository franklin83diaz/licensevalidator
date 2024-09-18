package internal

import (
	"os"
)

func CheckLicenseFromFile(protectedId string, pathLicenseFile string) (bool, error) {
	// Open the file
	file, err := os.Open(pathLicenseFile)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the file
	buffer := make([]byte, 1024)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}

	//TODO: Implement the logic to check the license from the file

	return true, nil

}
