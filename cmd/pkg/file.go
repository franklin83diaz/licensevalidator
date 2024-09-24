package pkg

import (
	"embed"
)

//go:embed secret/*
var secretFiles embed.FS

func ReadInternalfile(fileName string) (string, error) {
	fs, err := secretFiles.Open("secret/" + fileName)
	if err != nil {
		return "", err
	}
	defer fs.Close()

	// Read the file
	buffer := make([]byte, 2048)
	len, err := fs.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:len]), nil

}
