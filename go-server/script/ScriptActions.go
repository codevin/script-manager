package script 

import (
   log "github.com/sirupsen/logrus"
   // "time"
   // "github.com/matryer/runner"
   // "encoding/json"
)

func (script *ScriptInfo) Reset() {
   script.Status.mu.Lock()
   script.Status.RunStatus=""
   script.Status.RunSubStatus=""
   script.Status.LogStatus="Inactive"
   script.Status.UIStatus = "Inactive"
   script.Status.Logs = nil
   script.Status.UIRequest=""
   script.Status.Pid=-1
   script.Status.mu.Unlock()
   script.Status.ActionOptions = make(map[string]bool, 10)
   script.Status.ActionOptions["StartAction"]=true
   script.Status.ActionOptions["StopAction"]=false
   script.Status.ActionOptions["ResetAction"]=false
   script.Status.ActionOptions["ShowLogsAction"]=false
   script.Status.ActionOptions["ShowUIAction"]=false
}




func (script *ScriptInfo) SetCannotStart(substatus string, reason string) {
   log.Info("StatusUpdate: script:", script.Name, " Set to CannotStart", " Reason:", reason)
   script.Status.mu.Lock()
   script.Status.RunStatus="CannotStart"
   script.Status.ActionOptions["StartAction"]=false
   script.Status.ActionOptions["StopAction"]=false
   script.Status.ActionOptions["ResetAction"]=true
   script.Status.RunSubStatus=substatus
   script.Status.mu.Unlock()
}

func (script *ScriptInfo) SetRunning(substatus string, reason string) {
   log.Info("StatusUpdate: script:", script.Name, " Set to Running", " Reason:", reason)
   script.Status.mu.Lock()
   script.Status.RunStatus="Running"
   script.Status.ActionOptions["StartAction"]=false
   script.Status.ActionOptions["StopAction"]=true
   script.Status.RunSubStatus=substatus
   script.Status.mu.Unlock()
}

func (script *ScriptInfo) SetStopped(substatus string, reason string) {
   log.Info("StatusUpdate: script:", script.Name, " Set to Stopped", " Reason:", reason)
   script.Status.mu.Lock()
   script.Status.RunStatus="Stopped"
   script.Status.ActionOptions["StopAction"]=false
   script.Status.ActionOptions["ResetAction"]=true
   script.Status.ActionOptions["ShowLogsAction"]=true
   script.Status.ActionOptions["ShowUIAction"]=false
   script.Status.RunSubStatus=substatus
   script.Status.mu.Unlock()
}

func (script  *ScriptInfo) SetError(substatus string, reason string) {
   log.Info("StatusUpdate: script:", script.Name, " Set to Error", " Reason:", reason)
   script.Status.mu.Lock()
   script.Status.RunStatus="Error"
   script.Status.ActionOptions["StartAction"]=false
   script.Status.ActionOptions["StopAction"]=false
   script.Status.ActionOptions["ResetAction"]=true
   script.Status.RunSubStatus=substatus
   script.Status.mu.Unlock()
}

func (script *ScriptInfo) IsRunning() bool {
  return script.Status.RunStatus == "Running"
}

func (script *ScriptInfo) HasStopped() bool {
  return script.Status.RunStatus == "Stopped"
}

func (script *ScriptInfo) HasError() bool {
  return script.Status.RunStatus == "Error"
}

func (script *ScriptInfo) GetStatus() string {
   return script.Status.RunStatus
}

func (script *ScriptInfo) GetSubStatus() string {
   return script.Status.RunSubStatus
}

