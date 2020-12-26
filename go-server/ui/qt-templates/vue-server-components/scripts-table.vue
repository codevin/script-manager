<template>
  <v-container>
   <v-data-table
            :headers="headers"
            :items="scripts"
            item-key="Name"
            :items-per-page="10"
            class="elevation-1 mytable"
  >

    <template v-slot:top>
          <v-toolbar flat >
               <v-toolbar-title>Scripts Manager</v-toolbar-title>
               <v-divider
                 class="mx-4"
                 inset
                 vertical
               ></v-divider>
               <v-spacer></v-spacer>
      <v-btn
        color="primary"
        @click="updateScripts"
      >
        Refresh
      </v-btn>

          </v-toolbar>
    </template>

    <template v-slot:item.messages="{item}">
        <strong>{{ item.RunStatus }}</strong> <br>
        {{ item.RunSubStatus }} <br>
        {{ callsCount }} 
    </template>
    
    <template v-slot:item.actions="{item}">
      <span :key="item.renderKey" >
        <script-action v-for="action in action_keys" :p-script="item" :action="action"></script-action> &nbsp;
     </span>
    </template>

    <template v-slot:no-data>
      <v-btn
        color="primary"
        @click="updateScripts"
      >
        Update Table Data 
      </v-btn>
    </template>
  </v-data-table>
  </v-container>
</template>

<script>
module.exports = {
    data () { return {
      refreshTableKey: 0, 
      uidialog: false,
      updateUrl: "/update",
      loading: true,
      headers: [
          {
            text: 'Script Name', align: 'start', sortable: false, value: 'Name',
          }, 
          {
            text: 'Command Path', align: 'start', sortable: false, value: 'CmdPath'
          },
          { text: 'Messages', value: 'messages', sortable: false },
          { text: 'Actions', value: 'actions', sortable: false },
          { text: 'Duration', value: 'Duration' },
      ],
      } 
    },
    methods: {
       ...Vuex.mapActions(['s_UpdateScripts', 's_UpdateScript']),
       updateScripts() {
           this.s_UpdateScripts()
           this.refreshScripts()
       },
       updateScript(script_id) {
           this.s_UpdateScript({id: script_id})
       },
       openScriptUI (item) {
           this.editedIndex = this.scripts.indexOf(item)
           this.editedItem = Object.assign({}, item)
           this.dialog = true
       },
       openLogUI (item) {
           this.editedIndex = this.scripts.indexOf(item)
           this.editedItem = Object.assign({}, item)
           this.dialogLogUI = true
       },
       refreshScripts () {
          console.log("Hope to refresh scripts")
          // this.refreshTableKey++
          this.$forceUpdate()
       }
    },
    components: {
       'script-action': httpVueLoader('/vue/components/script-action.vue'),
    },
    computed: {
       ...Vuex.mapGetters(['scripts', 'callsCount']),
       ...Vuex.mapState({mytotalScripts: 'totalScripts'}),  // Any variable from store.state,
       action_keys: {
         get: function() {
              console.log("action_keys:", this.$store.state.script_action_keys)
              return this.$store.state.script_action_keys
         }
       },
       refreshTableKey2: {
         get: "actions" + this.refreshTableKey,
       }
    },
    mounted () {
       // This is where initialization of scripts happens.
       this.s_UpdateScripts();
       self=this;
       this.$EventBus.$on("refreshScript", function(req){
             console.log("Updating script: req=", req)
             self.updateScript(req.script_id);
       });
    }
}
</script>




