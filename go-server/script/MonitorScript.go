package script 

import (
   log "github.com/sirupsen/logrus"
   "time"
   "bufio"
   "strings"
   // "github.com/matryer/runner"
   // "encoding/json"
)

// Only for consumption of Monitor.
// When we already have UIRequest pending, we should suspend reading more.
func (script *ScriptInfo) processScriptOutput(line string) bool {

   // log.Info("Monitor: Our incoming line is:====", line, "========")
   req_type := ""
   if ( len(line)>0 && line[0] == ':') {
       req_type = "UIRequest"
       if ( script.Status.UIStatus == "HaveUIRequest" ) {
          log.Info("Monitor: Programming error: Current request is not serviced as yet. And new request has come:", line)
          return false
       }

   } else {
       req_type = "Logs"
   }
   script.Status.mu.Lock()
   if (req_type == "UIRequest") {
            script.Status.UIRequest = line
            script.Status.UIStatus = "HaveUIRequest" 
            script.Status.ActionOptions["ShowUIAction"]=true

   } else if (req_type == "Logs") {
            if (script.Status.Logs == nil ) {
                script.Status.Logs = &LogInfo{}
            }
            script.Status.Logs.Logs += line + "\n"
            script.Status.Logs.Timestamp = time.Now() 
            script.Status.LogStatus = "HaveLogs"
            script.Status.ActionOptions["ShowLogsAction"]=true
   }
   script.Status.mu.Unlock()
   return true
}

func scan_lines(lines *[]string, s string) int {
     // log.Info("Entering scanner.")
     scanner := bufio.NewScanner(strings.NewReader(s))

     scanner.Split(bufio.ScanLines)
     n:=0
     for scanner.Scan() {
         *lines = append(*lines, scanner.Text())
         n++
     }
     return n
}

// Blocks on read.
func read_from_script(script *ScriptInfo, from string) ([]string, bool) {

     readBuf := make([]byte, 1024)
     lines := make([]string, 0, 10)
     var n  int
     var err error

     // log.Info("ScriptActions: Waiting for read from outpipe")
     if (from == "Stdout") {
         n , err = script._OutPipe.Read(readBuf)
     } else if (from == "Stderr") {
         n , err = script._ErrPipe.Read(readBuf)
     } else {
         return lines, false
     }

     if (n < 0 || err != nil ) {
          log.Info("ScriptActions: Monitor(): Script: ", script.Name, " error=", err, " Monitoring stopped.")
          return lines, false
     }
     if ( n == 0 ) {
        return lines, true
     }
     scan_lines(&lines, string(readBuf[0:n-1]))
     return lines, true
}


// Waits for script to output to stdout and then makes the message available.
func (script *ScriptInfo) Monitor(from string) {

   // log.Info("ScriptInfo:Monitor(): Started for script:", script.Name)
   for {
          if ( !script.IsRunning() ) {
            log.Info("ScriptActions: Monitor(): Script: ", script.Name, " Not in Running, so stopped Polling. script.status=", script.Status.RunStatus)
            break
          }

          lines, ok := read_from_script(script, from) // blocks
          if (! ok ) {
             break
          }
          if ( len(lines) > 0 ) {
             for _, line := range(lines) {
                 for {
                    if ( script.processScriptOutput(line) ) {
                       break
                    }
                    time.Sleep(5 * time.Second)
                 }
             }
          }
       }
}


func (script *ScriptInfo) MonitorStderr() {
  script.Monitor("Stderr") 
}

func (script *ScriptInfo) MonitorStdout() {
  script.Monitor("Stdout")
}

