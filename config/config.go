package config

import (
	"github.com/BurntSushi/toml"
)

func LoadConfig(fpath string) (*Config, error) {
	if len(fpath) == 0 {
		fpath = "./config.toml"
	}
	var config Config
	_, err := toml.DecodeFile(fpath, &config)
	return &config, err
}

type Config struct {
	IsDebug bool
	Server  Server
	Log     Log
}

type Log struct {
	Dir string
}

type Server struct {
	Address  string
	HttpPort int32
}
