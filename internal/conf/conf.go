package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Port    string `json:"port"`
	Address string `json:"address"`
	Cycles  int    `json:"cycles"`
	Route   string `json:"route"`
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	confFile   = filepath.Join(basepath+"/../../config", "conf.json")
)

// function reads configs
func NewConf() (*Config, error) {
	f, err := os.Open(confFile)
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
