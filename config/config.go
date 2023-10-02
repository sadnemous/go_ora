package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// driver: oracle
// service: XE
// username: demo
// server: localhost
// port: 1521
// password: demo
type Config struct {
	Driver   string `yaml:"driver"`
	Service  string `yaml:"service"`
	Username string `yaml:"username"`
	Server   string `yaml:"server"`
	Port     int32  `yaml:"port"`
	Password string `yaml:"password"`
}

func (C *Config) GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile("local.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, C)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return C
}
