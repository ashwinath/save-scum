package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Files []FileConfig `yaml:"Files"`
}

type FileConfig struct {
	From           string   `yaml:"From"`
	To             string   `yaml:"To"`
	Flags          []string `yaml:"Flags"`
	Chown          Chown    `yaml:"Chown"`
	RemoveOriginal bool     `yaml:"RemoveOriginal"`
}

type Chown struct {
	Enabled bool   `yaml:"Enabled"`
	User    string `yaml:"User"`
	Group   string `yaml:"Group"`
}

func New(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
