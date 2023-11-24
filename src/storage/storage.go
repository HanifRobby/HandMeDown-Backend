package storage

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ImagesDir = "images"
)

// SaveImage saves the image to the specified directory.
func SaveImage(imageData []byte, fileName string) error {
	path := fmt.Sprintf("%s/%s", ImagesDir, fileName)
	err := ioutil.WriteFile(path, imageData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetImage retrieves the image from the specified directory.
func GetImage(fileName string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s", ImagesDir, fileName)
	imageData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}

// DeleteImage deletes the image from the specified directory.
func DeleteImage(fileName string) error {
	path := fmt.Sprintf("%s/%s", ImagesDir, fileName)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
