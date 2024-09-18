package licensevalidator

import (
	"fmt"
	"licensevalidator/internal"
)

func Validate(server string, serial string, pathLicenseFile string) (licOk bool, err error) {

	//Validate pathLicenseFile
	if pathLicenseFile == "" {
		return false, fmt.Errorf("path to license file is empty")
	}

	//Validate serial
	if serial == "" {
		return false, fmt.Errorf("serial number is empty")
	}

	//Validate server
	if server == "" {
		return false, fmt.Errorf("server address is empty")
	}

	//Get protected id
	protectedId, err := GetId("myapp")
	if err != nil {
		return
	}

	//file exists
	if internal.FileExists(pathLicenseFile) {
		licOk, err = internal.CheckLicenseFromFile(protectedId, pathLicenseFile)
	}

	// If the license is not valid or the file does not exist, proccess to validate from the server
	if !licOk {
		licOk, err = internal.CheckLicenseFromServer(protectedId, server, serial)
	}

	return
}
