// Package config used to load config from json file.
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
	var err error
	path, err = absPath(path)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(path)
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

// Dump .
func Dump(config interface{}, path string) error {
	var err error
	path, err = absPath(path)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, bytes, os.ModeAppend); err != nil {
		return err
	}

	return nil
}

var dir *string

// SetDir set the directory for relative path find.
func SetDir(d string) {
	dir = &d
}
func absPath(p string) (string, error) {
	if !filepath.IsAbs(p) {
		var (
			d   string
			err error
		)
		if dir == nil {
			d, err = os.Getwd()
			if err != nil {
				return "", nil
			}
		} else {
			d = *dir
		}
		p = filepath.Join(d, p)
	}
	return p, nil
}
