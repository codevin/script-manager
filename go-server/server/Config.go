package server

import (
  "../config"
)

type Config struct {
   Port string
   StaticDir string
}

var _config Config

func InitConfig() {
   config.SetServerConfig(_config)
}

func GetConfig() *Config {
  return &_config
}


func SetPort(port string) {
   _config.Port = port
}

func SetStaticDir(dir string) {
   _config.StaticDir = dir
}
