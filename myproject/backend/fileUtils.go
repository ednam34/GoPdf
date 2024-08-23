package backend

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
	// Créer un fichier temporaire dans le répertoire temporaire du système
	// tempDir := "./temp/"
	name := GetTempFile(".pdf")

	tempFile, err := os.Create("./temp/" + name) // Modifier le suffixe selon le type de fichier
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

	// Copier le contenu du fichier source dans le fichier temporaire
	_, err = io.Copy(tempFile, sourceFile)
	if err != nil {
		return "", fmt.Errorf("échec de la copie du fichier: %w", err)
	}
	fmt.Println("le nome 2 : " + tempFile.Name())
	return name, nil
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

	randId := strconv.Itoa(rand.Intn(10000000000))

	fileName = strings.Replace(fileName, "*", randId, -1)

	return fileName
}
