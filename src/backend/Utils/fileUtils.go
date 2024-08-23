package utils

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CopyFileToTemp(src string) (string, error) {

	name := GetTempFile(".pdf")

	tempFile, err := os.Create("./temp/" + name)
	if err != nil {
		return "", fmt.Errorf("échec de la création du fichier temporaire: %w", err)
	}
	defer tempFile.Close()

	// Ouvrir le fichier source
	sourceFile, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("échec de l'ouverture du fichier source: %w", err)
	}
	defer sourceFile.Close()

	_, err = io.Copy(tempFile, sourceFile)
	if err != nil {
		return "", fmt.Errorf("échec de la copie du fichier: %w", err)
	}
	fmt.Println("le nome 2 : " + tempFile.Name())
	return name, nil
}

func CopyFile(source string, destination string) error {

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("échec de l'ouverture du fichier source: %w", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("échec de la création du fichier de destination: %w", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("échec de la copie du fichier: %w", err)
	}

	err = destinationFile.Sync()
	if err != nil {
		return fmt.Errorf("échec de la synchronisation du fichier de destination: %w", err)
	}

	return nil
}

func DeleteAllTempFiles() error {
	tempDir := "./temp/"
	files, err := os.ReadDir(tempDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return err
	}

	// Iterate over files and delete each one
	for _, file := range files {
		filePath := filepath.Join(tempDir, file.Name())

		// Check if it is a file (and not a directory)
		if !file.IsDir() {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println("Error deleting file:", err)
				return err

			} else {
				fmt.Println("Deleted file:", filePath)
			}
		}
	}
	return nil
}

func GetTempFile(ext string) string {
	fileName := "temp-*" + ext

	randId := strconv.Itoa(rand.Intn(1000000000000))

	fileName = strings.Replace(fileName, "*", randId, -1)

	return fileName
}
