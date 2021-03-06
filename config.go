package main

import "github.com/BurntSushi/toml"

type Vno struct {
	Name string
	Path string
}

type Database struct {
	Name   string
	Host   string
	Bucket string
	Vno    []Vno
	Cnf    string
}

type Config struct {
	Database []Database
}

func ReadConfig(configFile string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
