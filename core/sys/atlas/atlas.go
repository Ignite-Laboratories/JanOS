package atlas

import (
	"encoding/json"
	"os"
)

func init() {
	configPath := "atlas.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// There's no Atlas available
		return
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	var cfg config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return
	}
	cfg.apply()
}
