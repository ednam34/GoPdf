package main

import (
	"context"
	"fmt"
	"gopdf/backend"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {

	if _, err := os.Stat("./temp/"); os.IsNotExist(err) {
		// Si le répertoire n'existe pas, le crée
		err := os.MkdirAll("./temp/", os.ModePerm)
		if err != nil {
			fmt.Println("échec de la création du répertoire %s: %w", "./temp/", err)
		}
		fmt.Printf("Répertoire créé: %s\n", "./temp/")
	} else {
		fmt.Printf("Le répertoire existe déjà: %s\n", "./temp/")
	}

	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Println("test")
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// OpenFile ouvre une boîte de dialogue pour sélectionner un fichier et lit son contenu
func (a *App) OpenFile() string {
	filePath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choisir un fichier",
	})
	if err != nil {
		return "an error"
	}

	if filePath == nil {
		return "select a file"
	}

	NewfilePath, err2 := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Enregistrer le fichier",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "pdf file",
				Pattern:     "*.pdf",
			},
		},
		DefaultFilename: "merge.pdf",
	})
	if err2 != nil {
		fmt.Println(err)
	}

	api.MergeCreateFile(filePath, NewfilePath, false, nil)
	return "The Pdf's have been mixed"
}

func (a *App) OptimizeFile() string {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choisir un fichier",
	})
	if err != nil {
		return "an error has occurred"
	}

	if filePath == "" {
		return "Select a file"
	}

	inputFile, err := os.Open(filePath)
	if err != nil {
		return "an error has occurred"
	}
	defer inputFile.Close()

	NewfilePath, err2 := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Enregistrer le fichier",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "pdf file",
				Pattern:     "*.pdf",
			},
		},
		DefaultFilename: "compressed.pdf",
	})
	if err2 != nil {
		fmt.Println(err)
	}

	// Créez le fichier de sortie
	outputFile, err := os.Create(NewfilePath)
	if err != nil {
		return "an error has occurred"
	}
	defer outputFile.Close()

	// Créez une nouvelle configuration
	conf := model.NewDefaultConfiguration()

	// Optimisez le PDF
	err = api.Optimize(inputFile, outputFile, conf)
	if err != nil {
		return "an error has occurred"
	}

	return "The pdf has been compressed"
}

func (a *App) ImgToPdf() string {
	filePath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choisir un fichier",
	})
	if err != nil {
		return "an error has occurred"
	}

	if filePath == nil {
		return "Select a file"
	}

	config := model.NewDefaultConfiguration()

	outputPath, err2 := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Enregistrer le fichier",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "pdf file",
				Pattern:     "*.pdf",
			},
		},
		DefaultFilename: "merge.pdf",
	})
	if err2 != nil {
		fmt.Println(err)
	}

	err = api.ImportImagesFile(filePath, outputPath, nil, config)
	if err != nil {
		return "an error has occured"
	}
	if len(filePath) == 1 {
		return "the image has been converted"
	} else {
		return "the images have been converted"

	}
}

func (a *App) OpenSinglePdf() string {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choisir un fichier",
	})

	if err != nil {
		return "an error has occurred"
	}

	if filePath == "" {
		return "Select a file"
	}

	tempFilePath, err := backend.CopyFileToTemp(filePath)
	if err != nil {
		log.Println(err)
		return "Failed to copy file to temp directory"
	}

	// Retourne le chemin du fichier temporaire à utiliser dans le frontend
	return tempFilePath
}

func (a *App) PrintT(an any) {
	fmt.Println(an)
}

func (a *App) PrintString(s string) {
	fmt.Println(s)
}
func (a *App) PrintArr(s []int) {
	fmt.Println(s)
}

func (a *App) RemovePages(pagesList []int, filePath string) string {
	fileName := backend.GetTempFile(".pdf")
	outputFile := "./temp/" + fileName

	fmt.Println("ici on a : ", filePath)
	arr := backend.IntSliceToStr(pagesList)
	api.RemovePagesFile("./temp/"+filePath, outputFile, arr, nil)

	return fileName
}

func (a *App) beforeClosing(ctx context.Context) bool {

	err := backend.DeleteAllTempFiles()

	if err != nil {
		log.Fatal(err)
		return true
	}

	return false
}
