package internal

import (
	"os"
)

// FileExists checks if a file exists
// receives the path to the file as a parameter
// returns a true if the file exists
func FileExists(pathLicenseFile string) bool {
	_, err := os.Stat(pathLicenseFile)
	return !os.IsNotExist(err)
}

// Read file reads a file and returns its content
// receives the path to the file as a parameter
// returns the content of the file and an error
func ReadFile(path string) (content string, err error) {

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file
	buffer := make([]byte, 2024)
	len, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:len]), nil

}
