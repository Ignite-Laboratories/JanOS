package config

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

var Current Configuration
var filename = "config.json"

type Configuration struct {
	Seed            string        `json:"seed"`
	DefaultDuration time.Duration `json:"defaultDuration"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Initialize() {
	if !fileExists(filename) {
		Current = Configuration{
			Seed:            uuid.New().String(),
			DefaultDuration: time.Millisecond * 10,
		}
		createConfigFile()
	}

	loadConfigFile()
}

func createConfigFile() {
	jsonFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(Current)

	if err != nil {
		log.Fatal(err)
	}
}

func loadConfigFile() {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	var config Configuration

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal(err)
	}

	Current = config
}
