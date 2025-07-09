package ordering

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App struct {
		WorkerCount int `yaml:"workerCount"`
	} `yaml:"app"`
}

// Load reads the file at path and decodes it into Config.
func Load(path string) (*Config, error) {
	// set config into object
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
