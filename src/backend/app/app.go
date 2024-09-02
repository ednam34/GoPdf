package Application

import (
	"context"
	"fmt"
	utils "gopdf/backend/Utils"

	"log"
	"os"
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
func (a *App) Startup(ctx context.Context) {

	if _, err := os.Stat("./temp/"); os.IsNotExist(err) {
		// Si le répertoire n'existe pas, le crée
		err := os.MkdirAll("./temp/", os.ModePerm)
		if err != nil {
			fmt.Printf("échec de la création du répertoire %s: %e", "./temp/", err)
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

func (a *App) BeforeClosing(ctx context.Context) bool {

	err := utils.DeleteAllTempFiles("./temp/")

	if err != nil {
		log.Fatal(err)
		return true
	}

	return false
}
