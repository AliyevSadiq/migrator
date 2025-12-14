package migrator

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ModelsFolder string = "models"        // default
var MigrationsFolder string = "migration" // default

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	if folder := os.Getenv("MODELS_FOLDER"); folder != "" {
		ModelsFolder = folder
	}

	if folder := os.Getenv("MIGRATIONS_FOLDER"); folder != "" {
		MigrationsFolder = folder
	}

	// Create migrations folder if not exists
	if _, err := os.Stat(MigrationsFolder); os.IsNotExist(err) {
		os.Mkdir(MigrationsFolder, os.ModePerm)
	}
}
