package licensevalidator

import (
	"fmt"

	"github.com/denisbrodbeck/machineid"
)

// GetId returns the protected id of the machine
// recieves the aplication Name as a parameter
func GetId(appName string) (string, error) {
	protectedId, err := machineid.ProtectedID(appName)
	if err != nil {
		return "", fmt.Errorf("failed to get protected id: %v", err)
	}

	return protectedId, nil
}
