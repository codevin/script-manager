<template>
      <v-dialog
         v-model="uimodal_active"  
         persistent
         max-width="500px"
         :key="script_id"
      >
    <v-card v-if="script_id" height="100%" >
      <v-card-title>Script: {{script.Name}} </v-card-title>
      <v-card-subtitle><strong>{{script.RunStatus}} </strong>
      <v-btn v-if="actionbtn"
               color="blue darken-1"
               text
               @click="actionFn(actionbtn.key)"
        >
           {{actionbtn.text}} 
      </v-btn> 
      <v-card-text><strong>&nbsp;UI Status: {{uiStatus}}</strong></v-card-text>
      </v-card-subtitle>

      <v-card-text class="justify-center"
      >
         <keep-alive>
            <component 
                v-if="haveActiveWidget"
                v-bind:is="activeUIWidget" 
                :request="UIWidgetInfo" 
                :key="activeUIWidget" 
            ></component>
         </keep-alive>
      </v-card-text>

      <v-card-text class="justify-center" >
         <script-logs :logs="logs" :log-status="logStatus"></script-logs> 
      </v-card-text>

      <v-card-actions :key="script_id" >
        <v-spacer></v-spacer>
        <v-btn
            color="blue darken-1"
            text
            @click="pollForUI()"
        >
            Refresh UI 
        </v-btn>
        <v-btn
            color="blue darken-1"
            text
            @click.stop="closeModal()"
        >
           Close 
        </v-btn>
      </v-card-actions>

   </v-card>


    </v-dialog>
</template>

<script>
module.exports = {
    components: {
       'script-logs': httpVueLoader("/vue/components/script-logs.vue"),
       'widget-script-not-active': {
          template: `<p>Script is Not active.</p>`
       },
       'widget-no-ui': {
          template: `<p>Script has no UI right now.</p>`
       },
       'MessageBoxVue': httpVueLoader("/vue/components/widget-message-box.vue"),
       'InfoBoxVue': httpVueLoader("/vue/components/widget-info-box.vue"),
       'TextBoxVue': httpVueLoader("/vue/components/widget-text-box.vue"),
       'YesNoBoxVue': httpVueLoader("/vue/components/widget-yesno-box.vue"),
       'InputBoxVue': httpVueLoader("/vue/components/widget-input-box.vue"),
       'MenuBoxVue': httpVueLoader("/vue/components/widget-menu-box.vue"),
       'ChecklistBoxVue': httpVueLoader("/vue/components/widget-checklist-box.vue"),
       'RadiolistBoxVue': httpVueLoader("/vue/components/widget-radiolist-box.vue"),
       'PasswordBoxVue': httpVueLoader("/vue/components/widget-password-box.vue"),
       'TextFileBoxVue': httpVueLoader("/vue/components/widget-textfile-box.vue"),
       'ProgressBoxVue': httpVueLoader("/vue/components/widget-progress-box.vue"),
    },
    data: () => ({
      key: 0,
      activeUIWidget: "widget-script-not-active",
      haveActiveWidget: false,   // Enables or disables widget display. 
      UIWidgetInfo: {},
      UIWidgetResponseReady: false, 
      logs: "",
      uiStatus: "",
      logStatus: "",
      updatableWidget: false,
      ts_polled: {  
             // timestamps
           log: new Date(),
           ui: new Date(),
      },
    }),
    computed: {
       ...Vuex.mapGetters(['getScript', 'callsCount']),
       uimodal_active: function() {
           return this.$store.state.uimodal.script_id && this.$store.state.uimodal_active
       },
       script_id: function() {
           return this.$store.state.uimodal.script_id
       },
       script: function() {
             var s = this.$store.state.uimodal.script || {}
             return s
       },
       uiNeedsPolling: function() {
           return this.haveActiveWidget && this.updatableWidget
       },
       actionbtn: function() {
            var self=this
            var arr = ['reset', 'start', 'stop' ]
            var final_obj=null
            arr.forEach(function(a) {
                 obj=self.$store.state.script_actions[a]
                 if ( self.script.Actions[obj.key] ) {
                    final_obj=obj
                 } 
            });
            return final_obj
       },
    },
    watch: {
      uimodal_active: function(newVal, oldVal) {
          if( newVal ) {   // true means we are opening the dialog.
            this.openModal()
          }
      },
    },
    methods: {
      actionFn: function(action) {
          console.log("Action requested: ", action) 
      },
      closeModal: function () {
          console.log("Closing modal.") 
          this.$EventBus.$emit("refreshScript", {script_id: this.script_id})
          this.$store.dispatch('deactivate_uimodal')
      },
      openModal: function() {
          // Get updated UI and logs.
          this.pollForUI("force")
          this.pollForLogs("force")
          console.log("Refreshing UI on showing uimodal.")
      },
      deactivateUIWidget: function() {
          this.haveActiveWidget=false
          this.updatableWidget=false
          if (! this.script_id || this.script.RunStatus != "Running") {
             this.uiStatus = "No Script is running."
             this.activeUIWidget="widget-script-not-active"
             this.updatableWidget=false
          } else {
             this.uiStatus = "Script has no UI request."
             this.activeUIWidget="widget-no-ui"
          }
      },
      setActiveUIWidget: function(m, updatable) {
             // Note: Sometimes we may have to use this.
             // this.$nextTick(() => { this.showSomething = true });
          this.activeUIWidget = m.WidgetName 
          this.UIWidgetInfo = m
          this.haveActiveWidget=true
          this.updatableWidget = updatable 
      },
      updateWidgetUI: function(m) {
          this.UIWidgetInfo = m
          this.haveActiveWidget=true
          this.updatableWidget=true
          this.ackUIUpdate()
      },
      doneResponseFromUI: function(response, nextFn) {
         var command=response.command || ":reply"
         var content=response.content || response 
         var self=this
         this.interact(command, content, null, function(response, noresponse) {
              self.deactivateUIWidget();
              if (nextFn) nextFn();
         });
      },
      ackUIUpdate: function() {
            // same reply as done.
          var command=":reply"
          var content=":exit 0"
          this.interact(command, content);
      },
      pollForUI: function(force) {
           if (! force) {
               if ( ! this.uiNeedsPolling ) return;
               if ( ! this.$methods.timeout(this.ts_polled.ui, 2000) ) {
                   return 
               }
           }
           var self=this
           this.ts_polled.ui = this.$methods.now()

           this.interact(":ui", "", "UIRequest", function(response, noresponse) {

                console.log("pollForUI Response:", response, " NoResponse=", noresponse)
                if ( response ) {
                    if (response.MessageSubType == "NewVueWidget") {
                        self.uiStatus = "New Interaction Requested"
                        self.setActiveUIWidget(response.Message, null)
                        // Remains open till user interacts and responds. 
                        return

                    } else if (response.MessageSubType == "NewUpdatableVueWidget") {
                        self.uiStatus = "Widget Showing Updates"
                        self.setActiveUIWidget(response.Message, "updatable")
                        // Needs to send response immediately. 
                        self.ackUIUpdate()
                        return

                    } else if  (response.MessageSubType == "WidgetUpdate") {
                        console.log("Updating current widget.")
                        self.uiStatus = "Current widget updated"
                        self.updateWidgetUI(response.Message)
                        return
                    }
                }

                var response = response || noresponse
                if (response && response.MessageType == "AckResponseOK") {
                    self.uiStatus = "Refresh Status: No UI Request."

                } else if (response && response.MessageType == "AckResponseError") {
                    // TODO: Show Error. 
                    self.uiStatus = "Error: " + response.Message
                } else {
                    self.uiStatus = "Refresh: Bad Server Response." 
                }
           });
      },
      pollForLogs: function(force) {
           if (! force) {
               if ( ! this.$methods.timeout(this.ts_polled.logs, 30000) ) {
                   return 
               }
           }
           var self=this
           this.ts_polled.logs = this.$methods.now()
           this.$store.dispatch('s_Logs', { 
                script_id: self.script.id, 
                nextFn: function(script_name, command, response, error) { 
                  if ( response.MessageType == "Logs" ) {
                      var loginfo = response.Message
                      self.logs = loginfo.Logs
                      var nowFormatted = moment(self.ts_polled.log).format('hh:mm:ss')
                      if ( self.ts_polled.log!= loginfo.Timestamp) {
                           self.logStatus = "New Logs updated at " + nowFormatted
                           self.ts_polled.log = loginfo.Timestamp
                      } else {
                           self.logStatus = "No New Logs. Checked At " + nowFormatted
                      }
                      return
                  }
                  self.logStatus = "No New Logs. Checked at: " + moment().format('hh:mm:ss')
               }
          })
      },
      clearLogs: function() {
         this.logs = ""
         self.logStatus = "Clearing all Logs..."
         self.ts_polled.log = new Date(); 
         this.interact(":clearlogs", "")
      },
      onLogWidgetRequest: function(req) {
         // refresh
         if (req== ":logs" ) { 
             this.pollForLogs("force")

         } else if (req == ":clearlogs" ) {
             this.clearLogs()
         }
      },
      interact: function(command, content, respType, nextFn) {
        var self=this
        respType = respType || ""
        this.$store.dispatch('s_Interact', { 
                script_id: self.script.id, 
                command: command,  
                content: content || "", 
                nextFn: function(script_name, command, response, error) { 
                    if (error) {
                       console.error("script-modal:updated():Interact command: error response received:", error)
                    } else if ( response.MessageType == respType ) {
                         if (nextFn) nextFn(response);
                    } else {
                         if (nextFn) nextFn(null, response);
                    }
                }
        });
      },
   },
   mounted: function() {
       var self=this
       this.$EventBus.$on('timer1sec', function(data) {
           // Works. Verified. console.log("timer1sec");
               if ( self.uimodal_active == true ) {
                  self.pollForUI();
                  self.pollForLogs();
               }
       });
     /*
       this.$EventBus.$on('timer5sec', function(data) {
       }),
       this.$EventBus.$on('timer30sec', function(data) {
       });  
     */
       this.$EventBus.$on('finishedUI', function(response) {
            console.log("script-modal: component finished event. response=", response)
            self.doneResponseFromUI(response, function(){
                 self.pollForUI("force")
                 self.pollForLogs("force")
            })
       });
       this.$EventBus.$on('userReq_log', function(data) {
            self.onLogWidgetRequest(data); 
       }); 
   } 
}
</script>


