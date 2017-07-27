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
	if err := Load("test.json", &config); err != nil {
		t.Fatal(err)
	}
	fmt.Println("name=" + config.Name)
}
func TestLoadQuiet(t *testing.T) {
	config := LoadQ("test.json", &TestConfig{}).(*TestConfig)
	if config == nil {
		t.Fatal("load config in quiet mode failed.")
	}
	fmt.Println("name=" + config.Name)
}

func TestDump(t *testing.T) {
	config := TestConfig{"test"}
	if err := Dump(config, "test-dump.json"); err != nil {
		t.Fatal(err)
	}
}
