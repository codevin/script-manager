package script

import (
   "sync"
   "io"
   "os/exec"
   "encoding/json"
   "time"
)

type ScriptInfo struct {
   Id  string  `json:"Id"`
   Name  string  `json:"Name"`
   CmdPath string `json:"CmdPath"`
   Args string    `json:"Args"`
   Duration int64 `json:"Duration"`
   Status StatusT
   ErrorMessage string `json:"ErrorMessage"`
   _InPipe io.WriteCloser `json:"-"`
   _OutPipe io.ReadCloser `json:"-"`
   _ErrPipe io.ReadCloser `json:"-"`
}


// var statuses map[string]bool{"":true,"Running":true,"Killed":true}

type StatusT struct {
  mu sync.Mutex `json:"-"`
  RunStatus  string    // "Inactive", CannotStart, Running, Stopped, Error
  RunSubStatus  string
  LogStatus string     // "Inactive", "HaveLogs", "NoLogs"
  UIStatus string      // "Inactive", "HaveUIRequest", "HaveNoRequest", "Error"
  UIRequest string  `json:"-"`
  ActionOptions map[string]bool
  Logs  *LogInfo `json:"-"`
  Pid	int `json:"-"`
  Cmd   *exec.Cmd   `json:"-"`
}

type LogInfo struct {
   Logs  string
   Timestamp time.Time
}

func (status StatusT) MarshalJSON() ([]byte, error) {
    var s map[string]interface{}
    s = make(map[string]interface{}, 10)
    s["RunStatus"]=status.RunStatus
    s["RunSubStatus"]=status.RunSubStatus
    s["LogStatus"]=status.LogStatus
    s["UIStatus"]=status.UIStatus
    s["ActionOptions"]=status.ActionOptions
    return json.Marshal(s)
}

func (script *ScriptInfo) MarshalJSON() ([]byte, error) {
    var s  map[string]interface{}
    s = make(map[string]interface{}, 10)
    if (script != nil) {
       s["Name"]=script.Name
       s["CmdPath"]=script.CmdPath
       s["Args"]=script.Args
       s["Duration"]=script.Duration
       s["Status"]=script.Status
    }
    return json.Marshal(s)
}

func (logs *LogInfo) MarshalJSON() ([]byte, error) {
    var s  map[string]interface{}
    if (logs != nil) {
       s["Logs"]=logs.Logs
       s["Timestamp"]=logs.Timestamp
    }
    return json.Marshal(s)
}

