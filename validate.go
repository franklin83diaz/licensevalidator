package licensevalidator

import (
	"fmt"
	"licensevalidator/internal"
)

// Validate validates the license
// receives:
// the server address,
// the serial number,
// the app name,
// the path to the license file,
// the public key and the private key
// returns:
// a boolean indicating if the license is valid
// an error if the validation fails
func Validate(server string, serial string, appNAme string, pathLicenseFile string, ServerPubKey string, AppPrivKey string) (licOk bool, err error) {

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
	protectedId, err := GetId(appNAme)
	if err != nil {
		return
	}

	//file exists
	fileExists := internal.FileExists(pathLicenseFile)
	if fileExists {
		licOk, err = internal.CheckLicenseFromFile(protectedId, pathLicenseFile, ServerPubKey)
	}

	// If the license is not valid or the file does not exist, proccess to validate from the server
	// if !licOk {
	// 	licOk, err = internal.CheckLicenseFromServer(protectedId, server, serial)
	// }

	return
}
