package config 


type Config struct {
   ServerModule interface{}  // Server related 
   ScriptModule interface{}  // Script Related 
   UIModule interface{}  // UI Related 
}

var config Config

func GetConfig() Config {
  return config
}

func SetScriptConfig(cfg interface{}) {
  config.ScriptModule = cfg
}

func SetServerConfig(cfg interface{}) {
  config.ServerModule = cfg
}

