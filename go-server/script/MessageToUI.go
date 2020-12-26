package script 

import (
  // log "github.com/sirupsen/logrus"
  // "strconv"
  . "../ui/types"
)

type MessageToUI struct {
  ScriptName string
  StatusCode int
  MessageType string     // "Response", "NewWidget", "UpdateWidget", "Logs" ...
  MessageSubType  string
  Error string
  Message  interface{}
}

// Script is running fine.
func AckUserWithOK(script_name string, message string)  *MessageToUI {
     m := &MessageToUI{}
     m.ScriptName = script_name
     m.MessageType = "AckResponseOK"
     m.MessageSubType = "OK"
     m.StatusCode = 0
     m.Message = message

     // log.Info("MessageToUI: OK. script=", script_name, " message=", message)

     return m
}

// Script is running fine.
func AckUserWithError(script_name string, err string, message string)  *MessageToUI {
     m := &MessageToUI{}
     m.ScriptName = script_name
     m.MessageType = "AckResponseError"
     m.MessageSubType = err // Any string 
     m.StatusCode = 0
     m.Message = message

     // log.Info("MessageToUI: Error. script=", script_name, " message=", message)

     return m
}

// Script is NOT running or has error state. 
func AckUserWithInternalError(script_name string, status_code int, message string) *MessageToUI {
     m := &MessageToUI{}
     m.ScriptName = script_name
     m.MessageType = "AckResponseInternalError"
     m.MessageSubType = "InternalError"
     m.StatusCode = status_code
     m.Message = message

     // log.Info("MessageToUI: InternalError script=", script_name, " ERROR. message=", message)
     return m
}

// Script is running fine.
func AckUserWithStatus(script_name string, message interface{})  *MessageToUI {
     m := &MessageToUI{}
     m.ScriptName = script_name
     m.MessageType = "Status"
     m.MessageSubType = "OK"
     m.StatusCode = 0
     m.Message = message

     // log.Info("MessageToUI: Status script=", script_name, " Status message=", message)

     return m
}

func SendLogsToUser(script_name string, logs *LogInfo) *MessageToUI  {
     m := &MessageToUI{}
     m.ScriptName = script_name
     m.StatusCode = 0
     m.MessageType = "Logs"
     m.MessageSubType = "Logs"
     m.Message = *logs

     // log.Info("MessageToUI: Logs script=", script_name, " Logs message=", m)
     return m
}

func SendUIRequestToUser(script_name string, wi *WidgetInfo) *MessageToUI {
       if (wi == nil) { 
           return AckUserWithError(script_name, "BadCmdFromScript", "Command Request is badly formed. Check the script. ")
       }

       message := &MessageToUI{}
       message.ScriptName = script_name
       message.StatusCode = 0
       message.MessageType = "UIRequest"
       message.MessageSubType = wi.WidgetType
       message.Message = wi

       // log.Info("MessageToUI: UIRequest. script=", script_name, " UI Request message=", message)
       return message
}

