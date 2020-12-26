/*  ModalMethods():
  
  modal_id -->
    modal:
       script_id: "" // If assigned
       ui_state: "" // ScriptNotActive, UINotActive, ActiveAndWaiting, ActiveAndUpdating
       component: "" // script's component
       ui_request: ""   // info sent by script.
       ui_response: ""  // info sent by user response (by modal) 
       poll_state: "" // Should poll script or not.

  polling:
    - script -- to check if there is new request or update UI.

*/
const ScriptModule={
   namespaced: true,
   state: () => {
      script_id: "",
      modal_id: "",
      status: "", // "InActive", "CantStart", "Running", "Stopped", "Error" 
      EventType: "", // "HaveUIRequest", "HaveUIUpdate", "Stopped", "Error"
      event: ""
      Logs: "",
      buttons: {},
   },
   getters: {
        getModalId: (state) => {
           return state.modal_id
        },
        getStatus: (state) => {
        }
   },
   actions: {
        setStatusAsInactive: (state) => () ({
           state.status = "InActive"
           state.buttons: {
              start: {
                 enable: true,
                 btnText: "Start",
                 btnMethod: "startScript",
              },
           }
        }),
        setStatusAsStopped: (state) => () ({
           state.status = "Stopped"
           state.buttons: {
              reset: {
                 enable: true,
                 btnText: "Reset",
                 btnMethod: "resetScript",
              },
           }
        }),
        setStatusAsRunning: (state) => () ({
           state.status = "Running"
           state.buttons: {
              stop: {
                 enable: true,
                 btnText: "Stop",
                 btnMethod: "stopScript",
              }, 
              showUI: {
                 enable: state.HaveUIRequest || state.HaveUIUpdate || state.HaveLogs,
                 btnText: "UI",
                 btnMethod: "showUI",
              },
          },
        }),
        setStatusAsError: (state) => () ({
           state.status = "Error"
           state.statusText = "Stopped due to Error. Check Logs"
           state.buttons: {
              reset: {
                 enable: true,
                 btnText: "Reset",
                 btnMethod: "resetScript",
              },
              showUI: {
                 enable: state.HasLogs,
                 btnText: "UI",
                 btnMethod: "showUI",
              },
           }
        }),
        processServerEvent: (state) => (event_type, event) ({
            switch event_type {
              case "Started": 
                     state.modal.widget_enable=false
                     state.modal.ui_request={}
                     state.modal.pollOK=true // secs. Large timeout.
                     state.modal.pollWaitTime=0 // Check after this timeout.
                     break;
              case "HaveUIRequest": 
                     state.modal.widget_enable=true
                     state.modal.ui_request=event.request
                     state.modal.pollOK=false // secs. Large timeout.
                     state.modal.refreshView();
                     break;
              case "HaveUIUpdate": 
                     state.modal.widget_enable=true
                     state.modal.ui_request=event.request 
                     state.modal.pollOK=true // Check after this timeout.
                     state.modal.pollWaitTime=300 // Check after this timeout.
                     state.modal.key++
                     break;
              case "Stopped": 
                     state.modal.widget_enable=false
                     state.modal.ui_request={}
                     state.modal.pollOK=false // Check after this timeout.
                     break;
              case "Error": 
                     state.modal.widget_enable=false
                     state.modal.error=event.error
                     state.modal.ui_request={}
                     state.modal.pollOK=false // Check after this timeout.
                     break;
              case "Reset":
                     state.modal.widget_enable=false
                     state.modal.error=""
                     state.modal.ui_request={}
                     state.modal.pollOK=false // Check after this timeout.
                     break;
            }
        }),
        processUserEvent: (state) => (event_type, event) ({
            switch event_type {
               case "Start": 
                     break;
               case "Stop":
                     break;
               case "Reset":
                     break;
               case "UIResponseReady":
                     break;
               case "GetProgress":
                     break;
               case "GetFile":
                     break;
               case "Refresh":
                     break;
            }
        }),
        assignModal: (state) => (modal_id) ({
            var modal=this.modals[modal_id]
            if (modal.script_id == script_id) {
                refreshModalView();
                return
            }
            if (modal.script_id != "") {
              freeModal(modal_id)
            }
            this.scripts_ui_state[script_id].modal=modal
            modal.script_id = script_id
            refreshModalView();
        }),
        freeAssignedModal: (state) => ({
             // Carry out actions to free this modal from assigned script ownership.
             var modal = modals[modal_id]
             var script_id = modal.script_id

             scripts[script_id].modal=null
             modals[modal_id]=null

             refreshModalView(modal_id);
        }),
    },
    mutations: {
    },
}



Vue.prototype.$ScriptWidgetsState=new Vue({
   data: () => ({
     scripts_ui: {},
     modals: {},  // Modal is where user interaction happens. Currently only one.
   }),
   methods: {
})


