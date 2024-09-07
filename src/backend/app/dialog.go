package Application

import (
	customError "gopdf/backend/app/Error"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func OpenSingleFileDialog(a *App) (string, error) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choose a file",
	})
	if err != nil {
		return "", err
	}

	if filePath == "" {
		return "", customError.ErrorNoFile()
	}
	return filePath, nil
}

func OpenMultipleFileDialog(a *App) ([]string, error) {
	filePath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choose a file",
	})
	if err != nil {
		return nil, err
	}

	if len(filePath) == 0 {
		return nil, customError.ErrorNoFile()
	}
	return filePath, nil
}
