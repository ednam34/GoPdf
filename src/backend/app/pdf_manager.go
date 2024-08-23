package Application

import (
	"fmt"
	utils "gopdf/backend/Utils"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// OpenFile ouvre une boîte de dialogue pour sélectionner un fichier et lit son contenu
func (a *App) MergePdf() string {
	filePath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choose a file",
	})
	if err != nil {
		return "an error"
	}

	if filePath == nil {
		return "select a file"
	}

	NewfilePath, err2 := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save file",
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

// Optimize a PDF File
func (a *App) OptimizePdf() string {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choose a file",
	})
	if err != nil {
		return "An error occurred during opening"
	}

	if filePath == "" {
		return "Select a file"
	}

	inputFile, err := os.Open(filePath)
	if err != nil {
		return "An error occurred during opening"
	}
	defer inputFile.Close()

	NewfilePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "pdf file",
				Pattern:     "*.pdf",
			},
		},
		DefaultFilename: "compressed.pdf",
	})
	if err != nil {
		fmt.Println(err)
	}

	// Create the output file
	outputFile, err := os.Create(NewfilePath)
	if err != nil {
		return "an error has occurred"
	}
	defer outputFile.Close()

	// TODO try to change the configuration
	conf := model.NewDefaultConfiguration()

	// Optimise the pdf
	err = api.Optimize(inputFile, outputFile, conf)
	if err != nil {
		return "an error has occurred"
	}

	return "The pdf has been compressed"
}

func (a *App) ImgToPdf() string {
	filePath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choose a file",
	})
	if err != nil {
		return "an error has occurred during opening"
	}

	if filePath == nil {
		return "Select a file"
	}

	config := model.NewDefaultConfiguration()

	outputPath, err2 := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save file",
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

	//Image will be on an A5 format
	imp, _ := api.Import("f:A5, pos:c", types.POINTS)

	err = api.ImportImagesFile(filePath, outputPath, imp, config)
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
		Title: "Choose a file",
	})

	if err != nil {
		return "an error has occurred"
	}

	if filePath == "" {
		return "Select a file"
	}

	tempFilePath, err := utils.CopyFileToTemp(filePath)
	if err != nil {
		log.Println(err)
		return "Failed to copy file to temp directory"
	}

	// Return the temp file name
	return tempFilePath
}

func (a *App) RemovePages(pagesList []int, filePath string) string {
	fileName := utils.GetTempFile(".pdf")
	outputFile := "./temp/" + fileName

	fmt.Println("ici on a : ", filePath)
	arr := utils.IntSliceToStr(pagesList)
	api.RemovePagesFile("./temp/"+filePath, outputFile, arr, nil)

	return fileName
}

func (a *App) SaveModifiedPDF(path string) {
	outputPath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "pdf file",
				Pattern:     "*.pdf",
			},
		},
		DefaultFilename: "merge.pdf",
	})
	if err != nil {
		fmt.Println(err)
	}

	err = utils.CopyFile("./temp/"+path, outputPath)

	if err != nil {
		fmt.Println(err)
	}
}
