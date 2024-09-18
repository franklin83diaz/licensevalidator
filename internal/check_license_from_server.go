package internal

import (
	"net/http"
)

// FileExists checks if a file exists
// receives the path to the file as a parameter
// returns a true if the file exists
func CheckLicenseFromServer(protectedId string, server string, serial string) (bool, error) {

	// Create the request
	req, err := http.NewRequest("GET", server, nil)
	if err != nil {
		return false, err
	}

	// Add the protected id and the serial number to the request
	q := req.URL.Query()
	q.Add("protectedId", protectedId)
	q.Add("serial", serial)
	req.URL.RawQuery = q.Encode()

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil

}
