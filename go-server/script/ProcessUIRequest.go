package script

import (
  log "github.com/sirupsen/logrus"
)

func ProcessMessageFromUI(script_name string, command string, content string) *MessageToUI {

    if ( script_name == "" ) {
       log.Error("Received empty script name.") 
       return AckUserWithError(script_name, "BadScriptName", "Bad/Empty script name")
    }

    script := LookupScript(script_name) 

    if ( script == nil ) {
       log.Error("script_name not found in registry. name=", script_name)
       return AckUserWithError(script_name, "ScriptNotFound", "Script not found")
    }

    switch {
      case command == ":reply" :
         fallthrough
      case command == ":respond" :
         // Reply for one of our scripts.
         if ( ! script.IsRunning()) {
            return AckUserWithError(script.Name, "NotRunning", "Script Not Running") 
         }
         if  ! script.HaveUIRequest()  {
            return AckUserWithError(script.Name, "ScriptHasNoRequest", "Script has no request.")
         }
         ok, err_str := script.SendResponseToScript(content) 
         if (ok ) {
            return AckUserWithOK(script.Name, "Response sent to Script.")
         } else {
            return AckUserWithError(script.Name, "CommunicationError", err_str) 
         }


      case command == ":reset" :
         if (script.HasStopped() || script.HasError()) {
            script.Reset()
            return AckUserWithOK(script.Name, "Script state set back to Init state.")
         } else {
            return AckUserWithError(script.Name, "NotInStoppedOrErrorState", "Script is not in Stopped state.")
         }

      case command == ":start":
         if ( script.IsRunning() ) {
            return AckUserWithError(script.Name, "AlreadyRunning", "App is already Running")
         }
         script.Start()
         return AckUserWithOK(script.Name, "Initiated Script Execution. Check the status/logs.")

      case command == ":ui":
         ok, wi := script.GetUIRequest()
         if ok {
            return SendUIRequestToUser(script.Name, wi)
         }
         return AckUserWithOK(script.Name, "No UI requests pending.")

      case command == ":logs" :
         if ( script.HaveLogs() ) {
            return SendLogsToUser(script.Name, script.GetLogs())
         } else {
            return AckUserWithError(script.Name, "NoLogs", "No Logs for the script")
         }


      case command == ":stop" :
         if ( script.IsRunning() ) {
            script.Stop()
            return AckUserWithOK(script.Name, "Stop Initiated. Refresh and Check.")
         }
         return AckUserWithError(script.Name, "NotRunning", "Script is NOT running. Refresh and Check.")

      case command == ":forcestop" :
         if ( script.IsRunning() || script.HasError() ) {
            script.ForceStop()
            return AckUserWithOK(script.Name, "ForceStop Initiated. Refresh and Check.")
         }
         return AckUserWithError(script.Name, "NotRunning", "Script is NOT running. Refresh and Check.")

      case command == ":status" :
         // var status map[string]interface{}
         status := make(map[string]interface{}, 10)
         status["LogStatus"] = script.Status.LogStatus
         status["UIStatus"] = script.Status.UIStatus
         status["Status"] = script.Status.RunStatus
         status["SubStatus"] = script.Status.RunSubStatus

         log.Info("Status =", status)

         return AckUserWithStatus(script.Name, status)

      case command == ":clearlogs" :
         if ( script.HaveLogs() ) {
            script.ClearLogs()
         }
         return AckUserWithOK(script.Name, "Logs Cleared.")

      default:
         return AckUserWithError(script.Name, "BadCommand", "ScriptManager: Unknown command received:" + command)
    }
}


