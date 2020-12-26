<template>
  <span>
   <v-btn small 
       v-if="activateButton" 
       color="deep-purple lighten-2"
       @click="do_action"
   >
     <i :class="iconClass"></i>
     {{ actionText }}
   </v-btn>
  </span>
</template>

<script>
module.exports = {
    props: {
      action: String,
      pScript: Object,
      localAction: String,
      scriptAction: String,
      command: String,
      uidialog: Boolean, 
    },
    data: () => ({
      dialog: false,
    }),
    methods: {
      ...Vuex.mapGetters([]),
      ...Vuex.mapActions(['s_Interact', 's_UpdateScript']),
      updateScript(id) {
           this.$store.dispatch('s_UpdateScript', {
               id: id, 
               nextFn: function(){
                  console.log("Refresh command: s_UpdateScript done.")
               }
           })
      },
      do_action() {
          console.log("action click.")
          if ( ! this.pScript ) {
                   return
          }
          var self=this
          var action_obj = this.$store.state.script_actions[ this.action ]

          if (action_obj.ui_modal) {
              console.log("UI Dialog will be initiated..")
              console.log("Script id before opening dialog:", this.pScript.id)
              this.$store.dispatch('activate_uimodal', {
                  script_id: this.pScript.id,
                  action: action_obj.key, 
              })
              return
          }

          var command = action_obj.command
          if ( command == ":refresh")  {
              self.updateScript(self.pScript.id)
          } else {
              this.$store.dispatch('s_Interact', { 
                        script_id: self.pScript.id, 
                        command: command, 
                        content: "", 
                        nextFn: function() { 
                            self.updateScript(self.pScript.id)
                            console.log("Interact command: s_UpdateScript done.")
                        }
                    })
          } 
      },
      initialize () {
          buttonState = initialState
          this.stateChanged = true 
          actionFn()
      },
    },
    watch: {
      loader () {
        const l = this.loader
        this[l] = !this[l]
        setTimeout(() => (this[l] = false), 3000)
        this.loader = null
      },
    },
    computed: {
      actionText: function() {
          return this.$store.state.script_actions[ this.action ].text
      },
      iconClass: function() {
          return this.$store.state.script_actions[ this.action ].icon
      },
      activateButton: function() {
          var action_obj = this.$store.state.script_actions[ this.action ]
          if ( ! action_obj ) {
             alert("Bad action key. Check the code.")
          } 
          if ( action_obj.local ) {
             return true
          }
          var script = this.pScript
          return script.Actions[ action_obj.key ]
      },
    },
    mounted () {
    },
}
</script>

