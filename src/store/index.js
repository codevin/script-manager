import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

var store = new Vuex.Store({
      state: {
        _scripts: [],
        _scripts_by_id: {},
        uimodal_active: false,
        uimodal_changed: 0,
        uimodal: {
          script_id: ""
        },
        totalScripts: 0,
        script_action_keys: ['refresh', 'reset', 'start', 'stop', 'showLogs', 'showUI' ], 
        script_actions: {
                refresh: { key: 'RefreshAction', text: 'Refresh', command: ":refresh", ui: null, icon: "refresh", local: true },
                reset: { key: 'ResetAction', text: 'Reset', command: ":reset", ui: null, icon: "transit_enterexit" },
                start: { key: 'StartAction', text: 'Start', command: ":start", ui: null, icon: "launch" },
                stop: { key: 'StopAction', text: 'Stop', command: ":stop", ui: null, icon: "cancel"},
                showLogs: { key: 'ShowLogsAction', text: 'Logs', command: null, ui_modal: true, icon: "file-outline"},
                showUI: { key: 'ShowUIAction', text: 'Interact', icon: "account", command: null, ui_modal: true},
        },
      },
      getters: {
        // returns a function
        getScript: (state) => (script_id) => {
           var scripts_by_id = state._scripts_by_id
           return scripts_by_id[script_id] 
        },
        scripts(state) {
           return state._scripts
        },
        runStatus: (state) => (script_id) => {
          if (state) ;
          var icons = {
          }
          var actions = {
             "Inactive": "Start",
             "CannotRun": "Reset",
             "Running": "Stop", 
             "Stopped": "Reset", 
             "Error": "Reset", 
          }
          var statusTexts = {
             "Inactive": "Not running.",
             "CannotRun": "Error starting. Can't Run.",
             "Running": "Script is Running", 
             "Stopped": "Script has Stopped.", 
             "Error": "Reset To Start.", 
          }
          var s = {}
          s.RunStatus = store.getters.getScript(script_id).RunStatus
          s.StatusText = statusTexts[s.StatusText] 
          s.action = actions[s.RunStatus]
          s.icon = icons[s.action]
          s.disabled = false
          s.depressed = false
          return s
       },
       logStatus: (state) => (script_id) => {
          if (state) ;
          var icons = {
             "No Logs": "fa fa-square-o",
             "See Logs": "fa fa-square",
          }
          var actions = {
             "Inactive": "No Logs",
             "HaveLogs": "See Logs", 
             "NoLogs": "No Logs", 
             "Error": "No Logs", 
          }
          var s = {}
          s.LogStatus = store.getters.getScript(script_id).LogStatus
          s.action = actions[s.LogStatus]
          s.icon = icons[s.action]
          s.disabled = (s.LogStatus != "HaveLogs") 
          s.depressed = false
          return s
       },
       uiStatus: (state) => (script_id) => {
          if (state) ;
          var icons = {
             "No UI": "fa fa-ban",
             "Open UI": "fa fa-user-o",
             "Error UI": "fa fa-ban",
          }
          var actions = {
             "Inactive": "No UI",
             "HaveUIRequest": "Open UI",
             "NoRequest": "No UI", 
             "Error": "Error UI", 
          }
          var s = {}
          s.UIStatus = store.getters.getScript(script_id).UIStatus
          s.action = actions[s.Status]
          s.icon = icons[s.action]
          s.click = "uiAction"
          s.disabled = (s.Status != "HaveUIRequest") 
          s.depressed = false
          return s
        },
      },
      actions: {      // use store.dispatch('...', {}) for calling actions.
        activate_uimodal({commit}, obj) {
            (commit);
            if ( ! obj || ! obj.script_id)  {
                console.error("activate_uimodal: script_id is null.")
                return
            }
            store.state.uimodal = {
                script_id: obj.script_id,
                action: obj.action, 
                script: store.getters.getScript(obj.script_id), 
            }
            store.state.uimodal_changed++
            store.state.uimodal_active = true
        },
        deactivate_uimodal({commit}) {
            (commit);
            store.state.uimodal = {}
            store.state.uimodal_active = false
            store.state.uimodal_changed=0
        },
        s_UpdateScript({ commit }, obj) {
              axios.post('/updateScript', {Id: obj.id})
              .then(function(r){
                  commit('_saveScript', r.data) 
                  if (obj.nextFn) {
                     obj.nextFn()
                  }
              }).catch(function (error) {
                  console.log(error);

              })
       },
       s_UpdateScripts({ commit }) {
            // store is defined here.
            return  axios.get("/update") 
               .then(function(response) {

                  var data = response.data
                  var keys = Object.keys(data)
                  keys.forEach(key => {
                      commit('_saveScript', data[key]) 
                  })
                  return 
                }).catch(function (error) {
                  console.log(error);
                })
      },
      s_Interact({ commit }, request) {
                (commit);
                var params = new FormData();
                params.append('ScriptName', request.script_id)
                params.append('Command', request.command)
                params.append('Content', request.content)
                const config = {
                  headers: {
                    'content-type': 'multipart/form-data'
                  }
                }
                if (request.command == "") {
                   console.log("Store interact: Command is null. Not interacting.")
                   alert("Store interact: Command is null. Not doing ajax interact.")
                   request.nextFn(request.script_name, request.command, {}, "Error: Bad command")
                }
                axios.post('/interact', params, config) 
                  .then(function(response){
                       var r = response.data 
                       if ( r.StatusCode != 0 ) {
                           alert("Something went wrong. message=" + r.Message)
                           return
                       }
                       if (request.nextFn) {
                           request.nextFn(request.script_name, request.command, r)
                       } else if ( r.MessageType != "AckResponseOK" ) {
                           alert("Message:" + r.Message + "/" + r.MessageSubType)
                       }
                       return
                }).catch(function (err) {
                   // handle error
                  console.log("Error with /interact:", err)
                })
           },
           s_Logs({ commit }, request) {
                (commit);
                var params = new FormData();
                params.append('ScriptName', request.script_id)
                const config = {
                  headers: {
                    'content-type': 'multipart/form-data'
                  }
                }
                axios.post('/logs', params, config) 
                  .then(function(response){
                       var r = response.data 
                       if ( r.StatusCode != 0 ) {
                           alert("Something went wrong. message=" + r.Message)
                           return
                       }
                       if (request.nextFn) {
                           request.nextFn(request.script_name, request.command, r)
                       } else if ( r.MessageType != "AckResponseOK" ) {
                           alert("Message:" + r.Message + "/" + r.MessageSubType)
                       }
                       return
                }).catch(function (err) {
                   // handle error
                  console.log("Error with /logs:", err)
                })
           },
      },
      mutations: {
        _saveScript(state, r) {

           var script = state._scripts_by_id[ r.Name ] 
           if ( ! script ) {
              script = {}
              script.id = r.Name  // Note: Not script.Id
              script.renderKey= 0
              state._scripts.push(script)
              state._scripts_by_id[ script.id ] = script 
           }
           script.renderKey++ 
           script.Name = r.Name
           script.CmdPath = r.CmdPath 
           script.Duration = r.Duration 
           script.RunStatus = r.Status.RunStatus || "Inactive" 
           script.RunSubStatus = r.Status.RunSubStatus || ""
           script.LogStatus = r.Status.LogStatus || "Inactive"
           script.UIStatus = r.Status.UIStatus || "Inactive"
           script.Actions = r.Status.ActionOptions

        },
      }
});

export default store
