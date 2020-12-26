package script 

import (
   log "github.com/sirupsen/logrus"
   // "github.com/matryer/runner"
   "time" 
   "os" 
   "os/exec" 
   "syscall"
   // "errors"
)

func forceKill(script *ScriptInfo) (bool, error) {
   pid := script.Status.Pid
   if ( pid == -1 ) {
      return false, nil
   }
   process, _ := os.FindProcess(pid)
   err := process.Signal(syscall.Signal(9))
   if (err == nil ) {
          return true, nil
   }
   return false, err
}


// Verified.
func IsProcessActive(script *ScriptInfo) (bool, error) {
     // os.FindProcess on linux just creates a structure. 
     pid := script.Status.Pid
     if pid == -1 {
       return false, nil
     }
     process, _ := os.FindProcess(pid)
     err := process.Signal(syscall.Signal(0))
     if (err == nil ) {
          return true, nil
     }
     return false, err
}

func  (script *ScriptInfo) Start() {
    cmd := exec.Command(script.CmdPath, script.Args)
    var err1, err2, err3 error
    script.Status.Cmd = cmd
    script._InPipe, err1 = cmd.StdinPipe()

    script._OutPipe, err2 = cmd.StdoutPipe()
    script._ErrPipe, err3 = cmd.StderrPipe()
    if err1 != nil || err2 != nil || err3 != nil {
          script.SetError("InitializationError", "Error creating Input/Output/Error pipe.")
          return
     }
     if err := cmd.Start(); err != nil {
	  script.SetError("CmdDidNotStart", "Command didn't start.")
          log.Error("Tasker: Error Executing script: cmd=", cmd, " Error=", err)
          return
     }
     log.Info("Tasker: Executing script: cmd=", cmd)
     var pid = -1
     if (cmd.Process != nil) {
        pid = cmd.Process.Pid
     }
     go func() {
         cmd.Wait()
	 script.SetStopped("ScriptStopped", "Script Stopped on its own or external means")
     }()
     script.Status.Pid = pid
     script.SetRunning("NormalStart", "Script Started normally.")
     go script.MonitorStderr()
     go script.MonitorStdout()
}

func (script *ScriptInfo)Stop() {

   if ( ! script.IsRunning() && ! script.HasError() ) {
      log.Info("Stop", "Script is neither running nor is in error. Status=", script.GetStatus())
      return
   }

   log.Info("Tasker: Script's defer function entered.")
   yes, _ := IsProcessActive(script)
   if ( yes ) {
            script._InPipe.Close()
            err := syscall.Kill(script.Status.Pid, syscall.SIGTERM)
            if (err != nil) {
               log.Info("Stop(): Unable to kill process: ", script.Name, " err=", err)
            }
            time.Sleep(1 * time.Second)
            yes, _ = IsProcessActive(script)
   }
   if ( yes ) {
	    script.SetError("Couldnotkill", "Couldn't kill - process still active. ")
	    log.Error("Exec: Unable to successfully kill process")
            return
   }
   log.Info("Exec: Process Stopped/Killed.") 
   script.SetStopped("NormalFinish", "Process Terminated Normally.")
}

func  (script *ScriptInfo) ForceStop() {
   script.SetError("ForceStop", "User Request for ForceStop..")
   if ( ! script.IsRunning() && ! script.HasError() ) {
      log.Info("ForceStop", "Script is neither running nor is in error. Status=", script.GetStatus())
      return
   }
   yes, _ := IsProcessActive(script)
   if yes {
       log.Info("ScriptInfo.ForceStop(): Requested. ", script.Name)
       yes, _ = forceKill(script)
       time.Sleep(1 * time.Second)
   }
   yes, _ = IsProcessActive(script)
   if ( yes  ) {
      script.SetError("CantForceKill", "Script couldn't be killed with kill -9.")
      log.Info("ScriptInfo:ForceStop(): Couldn't ForceStop the process. Status=", script.GetStatus())
      return
   }
   script.SetStopped("Force Kill Stop", "Script has been stopped with kill -9.")
   log.Info("ScriptInfo:ForceStop(): ForceStop Executed. Status=", script.GetStatus())
}

