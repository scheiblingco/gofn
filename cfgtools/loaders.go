// Load configurations from different sources into a struct
//
// This package provides helpers for loading configurations from different sources into a struct.
// You can use multiple functions to load configurations from different sources, such as JSON, YAML, XML and more.
//
// Example:
//
//	 package main
//
//	 import "github.com/scheiblingco/gofn/cfgtools"
//
//	 type Config struct {
//	 	Host string `json:"host"`
//	 	Port int    `json:"port"`
//	 }
//
//	 func main() {
//	 	config := Config{}
//	 	err := cfgtools.LoadJsonConfig("config.json", &config)
//	 	if err != nil {
//	 		panic(err)
//	 	}
//	}
package cfgtools

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"github.com/BurntSushi/toml"
	ini "github.com/subpop/go-ini"
	"gopkg.in/yaml.v2"
)

// Load configuration from Json into a struct
func LoadJsonConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, config)
	return err
}

// Load configuration from a yaml file into a struct
func LoadYamlConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, config)
	return err
}

// Load configuration from an xml file into a struct
func LoadXmlConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, config)
	return err
}

// Load configuration from a TOML file into a struct
func LoadTomlConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = toml.Decode(string(data), config)
	return err
}

// Load configuration from an INI-like file into a struct
func LoadIniConfig(path string, config interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = ini.Unmarshal(data, config)
	return err
}
