package migrator

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ModelsFolder string = "models" // default

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	if folder := os.Getenv("MODELS_FOLDER"); folder != "" {
		ModelsFolder = folder
	}
}
