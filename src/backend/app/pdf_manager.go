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

	filePath, err := OpenMultipleFileDialog(a)

	if err != nil {
		fmt.Println(err)
	}

	NewfilePath := utils.GetTempFile(".pdf")

	api.MergeCreateFile(filePath, "./temp/"+NewfilePath, false, nil)
	return NewfilePath
}

// Optimize a PDF File
func (a *App) OptimizePdf() string {
	filePath, err := OpenSingleFileDialog(a)

	if err != nil {
		fmt.Println(err)
	}

	inputFile, err := os.Open(filePath)
	if err != nil {
		return "An error occurred during opening"
	}
	defer inputFile.Close()

	NewfilePath := utils.GetTempFile(".pdf")

	// Create the output file
	outputFile, err := os.Create("./temp/" + NewfilePath)
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

	return NewfilePath
}

func (a *App) ImgToPdf() string {
	filesPath, err := OpenMultipleFileDialog(a)

	if err != nil {
		fmt.Println(err)
	}

	config := model.NewDefaultConfiguration()

	outputPath := utils.GetTempFile(".pdf")

	//Image will be on an A5 format
	imp, _ := api.Import("f:A5, pos:c", types.POINTS)

	err = api.ImportImagesFile(filesPath, "./temp/"+outputPath, imp, config)
	if err != nil {
		return "an error has occured"
	}
	return outputPath
}

func (a *App) OpenSinglePdf() string {
	filePath, err := OpenSingleFileDialog(a)

	if err != nil {
		fmt.Println(err)
	}
	tempFilePath, err := utils.CopyFileToTemp(filePath)
	if err != nil {
		log.Println(err)
		return "Failed to copy file to temp directory"
	}

	// Return the temp file name
	return tempFilePath
}

func (a *App) OpenSinglePdfFromPath(filePath string) string {
	if filePath == "" {
		return "Select a file"
	}
	fmt.Println("the path in go is : " + filePath)
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

func (a *App) MoovePdfPage(filePath string, page int) string {
	moovedPath, _ := utils.MoovePage("./temp/"+filePath, page)

	return moovedPath
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

func (a *App) ReorderPdf(filePath string, order []int) {

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

	utils.ReorderPages(filePath, order, outputPath)
}
