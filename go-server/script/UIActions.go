package script

import (
   log "github.com/sirupsen/logrus"
   "strings"
   . "../ui/types"
   "../ui"
   // "time"
)

func (script *ScriptInfo) HaveUIRequest() bool {
    return script.Status.UIStatus == "HaveUIRequest"
}

// This will not change state.
func (script *ScriptInfo) GetUIRequest() (bool, *WidgetInfo) {

   status := script.Status.UIStatus
   if  status  == "HaveUIRequest" {

        cmdline := script.Status.UIRequest
        var wi *WidgetInfo

        if  strings.HasPrefix(cmdline, ":cmdline ") {
             wi = ui.CreateNewWidgetRequest(script.Name, cmdline)

        } else if strings.HasPrefix(cmdline, ":update ") {
             wi = ui.CreateUpdateForWidget(script.Name, cmdline)
        }

        return true, wi
   }
   return false, nil 
}

// Called from server side.
func (script *ScriptInfo) SendResponseToScript(response_str string) (bool, string) {
    status := script.Status.UIStatus
    if ( status == "HaveUIRequest" ) {

       result := true
       script.Status.mu.Lock()
       script.Status.UIStatus = "NoRequest"
       script.Status.ActionOptions["ShowUIAction"]=false
       script.Status.mu.Unlock()

       log.Info("Script UIAction: Writing to script now. Look for 'DONE'") 
       response_str += "\n"
       _ , err := script._InPipe.Write([]byte(response_str))
       if (err == nil) {
            // Can't close here! Script's stdin is only getting switched.
            // log.Info("Closing script inpipe.")
            // err = script._InPipe.Close()
            // Closing the pipe here will let peer stop all communication...
       }
       if (err != nil) {
           script.Status.UIStatus = "Error"
           log.Info("Script UIAction: Error while Closing Inpipe. error=", err)
           result=false
       }
       log.Info("Script UIAction: DONE. See error if any above.") 

       if (result) {
           log.Info("ToScript: Wrote to script stdin: ", response_str)
           return true, ""
       } else {
           log.Error("ToScript: Unable to Write to Script. err=", err)
           return false, "Write Error:" + error.Error(err)
       }

   } else {
       log.Error("SendResponseToScript: Script was not waiting for response.")
       return false, "Script not waiting for response."
   }
}

