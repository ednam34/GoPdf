package utils

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
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

func MoovePage(filePath string, page int) (string, error) {

	pageMinusOne := strconv.Itoa(page - 1)
	pagePlusOne := strconv.Itoa(page + 1)
	pagePlustwo := strconv.Itoa(page + 2)

	// Fichiers temporaires
	firstPart := "./temp/" + GetTempFile(".pdf")
	movedPage := "./temp/" + GetTempFile(".pdf")
	nextPage := "./temp/" + GetTempFile(".pdf")
	finalPart := "./temp/" + GetTempFile(".pdf")
	outputFile := GetTempFile(".pdf")

	err := api.TrimFile(filePath, firstPart, []string{"1-" + pageMinusOne}, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = api.TrimFile(filePath, movedPage, []string{strconv.Itoa(page)}, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = api.TrimFile(filePath, nextPage, []string{pagePlusOne}, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = api.TrimFile(filePath, finalPart, []string{pagePlustwo + "-"}, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = api.MergeCreateFile([]string{firstPart, nextPage, movedPage, finalPart}, "./temp/"+outputFile, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Suppression des fichiers temporaires
	os.Remove(firstPart)
	os.Remove(movedPage)
	os.Remove(nextPage)
	os.Remove(finalPart)
	os.Remove(filePath)

	fmt.Println("heeeeeeeeeere : " + outputFile)

	return outputFile, nil

}
