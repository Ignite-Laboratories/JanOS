package common

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
)

var Current Config

type Config struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Initialize() {
	fullPath := filepath.Join(GetSystemDirectory(), GetComponentName()) + ".cmp"

	if !FileExists(fullPath) {
		Current = Config{
			ID:    uuid.New().String(),
			Title: GetComponentName(),
		}
		createConfigFile(fullPath)
	}

	loadConfigFile(fullPath)
}

func createConfigFile(filename string) {
	jsonFile, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(Current)

	if err != nil {
		log.Println(err)
	}
}

func loadConfigFile(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	var config Config

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&config)

	if err != nil {
		log.Println(err)
	}

	Current = config
}
