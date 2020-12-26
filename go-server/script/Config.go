package script

import (
   "../config"
)

type Config struct {
   ScriptDir string
   Scripts map[string](*ScriptInfo)
}

var _config Config

func InitConfig() {
   _config.ScriptDir = "" 
   _config.Scripts = make(map[string](*ScriptInfo), 10)
   config.SetScriptConfig(_config)
}

func GetConfig() *Config {
   return &_config
}

func SetScriptDir(dir string) {
    _config.ScriptDir = dir
}

func RegisterScript(name, cmd, args string, duration int64) {
   script := ScriptInfo{}
   script.Reset()
   script.Id = name
   script.Name = name
   script.CmdPath = GetConfig().ScriptDir + cmd
   script.Args = args
   script.Duration = duration

   _config.Scripts[name] = &script

   script.Reset()

}

func LookupScript(name string) *ScriptInfo {
   return _config.Scripts[name]
}
