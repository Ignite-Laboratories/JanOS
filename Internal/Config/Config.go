package Config

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"os"
)

var Current Configuration

type Configuration struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Network     string `json:"network"`
	Address     string `json:"address"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Initialize(filename string) {
	if !fileExists(filename) {
		Current = Configuration{
			ID:          uuid.New().String(),
			Title:       "Unknown",
			Description: "This is a newly instantiated component",
			Network:     "tcp",
			Address:     "0.0.0.0:0",
		}
		createConfigFile(filename)
	}

	loadConfigFile(filename)
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

	var config Configuration

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&config)

	if err != nil {
		log.Println(err)
	}

	Current = config
}
