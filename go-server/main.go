package main

import (
      log "github.com/sirupsen/logrus"
      "github.com/spf13/viper"
      "github.com/fsnotify/fsnotify"
      "./script"
      "fmt"
      "time"
      "./server"
      "./ui"
      // "flag"
)


var config_ver int=0

type ScriptConfig struct {
   Program string
   Args string
   Duration int64
   Enabled  bool
}

type Config struct {
  Port string
  StaticDir string
  Scripts map[string]ScriptConfig
  ScriptDir string
}

var config Config

func main() {
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
        viper.AddConfigPath(".")
        //viper.AddConfigPath("/etc/script-manager/")
        //viper.AddConfigPath("$HOME/.script-manager")
        viper.SetEnvPrefix("SCRIPTMANAGER_")

        err := viper.ReadInConfig()
        if err != nil { // Handle errors reading the config file
	   panic(fmt.Errorf("Fatal error config file: %s \n", err))
        }
        err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
        viper.WatchConfig()
        viper.OnConfigChange(func(e fsnotify.Event) {
	     fmt.Println("Config file changed:", e.Name, " Reloading and restarting the server.")
             config_ver++
        })

        // fmt.Println("Port=", config.Port)
        // fmt.Println("Example Script=", config.Scripts["test1-yesno"].Program)

        server.InitConfig()
        server.SetPort(config.Port)
        server.SetStaticDir(config.StaticDir)

        script.InitConfig()
        script.SetScriptDir(config.ScriptDir)

        for name, scr := range(config.Scripts) {
           script.RegisterScript(name, scr.Program, scr.Args, scr.Duration)
        }

        // script.RegisterScript("test1-yesno", "test1-yesno.sh", "", 3600)
        // script.RegisterScript("test2-inputbox", "test2-inputbox.sh", "", 3600)
        // script.RegisterScript("test3-menu", "test3-menu.sh", "", 3600)

        ui.InitWidgets()

	for {
           log.Info("Main: Started http Server.") 
           server.Begin()
           log.Info("Main: http server has quit... (Error) ") 
           time.Sleep(2 * time.Second)
	}
}

