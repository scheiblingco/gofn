package gohelpers

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadJsonConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, config)
	return err
}

func LoadYamlConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, config)
	return err
}

func LoadXmlConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, config)
	return err
}
