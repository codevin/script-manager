<template>
    <v-card class="justify-center" height="100%" >
      <v-toolbar dark 
        color="warning" 
        dense 
        elevation="4" 
      >
        <v-icon left>stop</v-icon>
        <v-toolbar-title class="white--text"> Logs </v-toolbar-title>
      </v-toolbar>

      <v-card-text class="body-1 text-body-1 py-3" max-height="500px">
         Status: {{logStatus}}
      </v-card-text>

      <v-card-text class="body-1 text-body-1 py-3" max-height="500px">
   <pre><code>{{logs}}</code></pre>
      </v-card-text>

      <v-card-actions class="justify-center" >
        <v-btn
          color="blue"
          text
          @click="clearLogs"
        >
          Clear 
        </v-btn>

        <v-btn
          color="blue"
          text
          @click="refreshLogs"
        >
          Refresh 
        </v-btn>

      </v-card-actions>

    </v-card>
</template>

<script>
module.exports = {
    props: {
      logs: String, 
      logStatus: String, 
    },
    data: () => ({
      noData: ""
    }),
    methods: {
       clearLogs: function() {
           var response = ":clearlogs";
           this.$EventBus.$emit('userReq_log', response)
       },
       refreshLogs: function() {
           var response = ":logs";
           this.$EventBus.$emit('userReq_log', response)
       },
    }
}
</script>
