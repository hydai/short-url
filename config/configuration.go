package config

import ()

type ServerConfiguration struct {
	Port        string `yaml:"port"`
	CaFilePath  string `yaml:"caFilePath"`
	KeyFilePath string `yaml:"keyFilePath"`
}

type DatabaseConfiguration struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Configuration struct {
	Server   ServerConfiguration   `yaml:"server"`
	Database DatabaseConfiguration `yaml:"database"`
}
