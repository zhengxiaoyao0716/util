// Package config used to load config from json file.
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load load a json file to config from give path.
//
// config is an instance of Json Struct, such as:
//
// type Config struct {
// 	Name string `json:"name"`
// }
//
func Load(path string, config interface{}) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(dir + "/" + path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, config); err != nil {
		return err
	}

	return nil
}

// LoadQ ignored the error and return nil instead.
func LoadQ(path string, config interface{}) interface{} {
	if err := Load(path, config); err != nil {
		return nil
	}
	return config
}
