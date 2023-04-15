package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type conf struct {
	Budget         int `yaml:"budget"`
	NbOuvrier      int `yaml:"nbOuvriers"`
	TravailOuvrier int `yaml:"travailOuvriers"`
	NbCitoyen      int `yaml:"nbCitoyens"`
	NbJour         int `yaml:"nbJours"`
	DayTime        int `yaml:"dayTime"`
}

func LoadConfig(filename string) (*conf, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &conf{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}
