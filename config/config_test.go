package config

import (
	"fmt"
	"testing"
)

type TestConfig struct {
	Name string `json:"name"`
}

func TestLoad(t *testing.T) {
	config := TestConfig{}
	Load("test.json", &config)
	fmt.Println("name=" + config.Name)
}
func TestLoadQuiet(t *testing.T) {
	fmt.Println("name=" + LoadQ("test.json", &TestConfig{}).(*TestConfig).Name)
}
