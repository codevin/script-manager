const managerModule={
   namespaced: true,
   state: () => {
     polled_scripts: {},
   },
   mutations: {
   }, 
   actions: {
      poll: (state) => ({
         var self = this
         by_id = {}
         return  axios.get("/update") 
             .then(function(response) {
                  var data = response.data
                  var keys = Object.keys(data)
                  keys.forEach(key => {
                     if ( !scriptExists(key) {
                        createScript(key, data)
                     }
                     commit('scripts/' + key + '/updateState', data[key]) 
                  })
                  return
             }).catch(function (error) {
                  console.log(error);
             })
      }),
      processScriptEvent: (script_id, event, view) => {
          // Analyze server info and carry out relevant actions. 
          if ( event.HasUIRequest || event.HasUIUpdateRequest ) {

              // First update script state
              this.scripts[event.script_id].updateUIState()

              // Let modal refresh.
              modal_id = this.scripts[event.script_id].getModalId()
              if (! modal_id) { 
                   // User has not chosen script's UI.
               } else {
                  this.modals(modal_id).refreshUI()
                  return
              } 
              var modal = this.assignModalToScript(script_id)
              if ( ! modal ) return 
              refreshModalUI(modal.id)
              return

          } else if ( event.ScriptStopped ) {

              modal_id = this.getModalId(script_id) 
              if ( ! modal_id ) return 
              freeScriptModal(modal_id)
          }
        }, 
        processModalEvent: () => {

        }
      })
