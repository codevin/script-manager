package script 

import (
   // log "github.com/sirupsen/logrus"
   // "time"
   // "github.com/matryer/runner"
   // "encoding/json"
)

func (script *ScriptInfo) ClearLogs() {
   script.Status.mu.Lock()
   script.Status.LogStatus = "NoLogs"
   script.Status.ActionOptions["ShowLogsAction"]=false
   script.Status.Logs = nil
   script.Status.mu.Unlock()
}


func (script *ScriptInfo) GetLogs() *LogInfo {
   if (script.Status.LogStatus == "HaveLogs" ) {
       return script.Status.Logs
   } else {
       return nil
   }
}

func (script *ScriptInfo) HaveLogs() bool {
    return script.Status.LogStatus == "HaveLogs"
}


