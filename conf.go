package conf

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port    string `json:"port"`
	Address string `json:"address"`
	Cycles  int    `json:"cycles"`
	Route   string `json:"route"`
}

func NewConf() (*Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
