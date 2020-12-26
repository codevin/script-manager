// Textfile display 
<template>
    <v-card class="justify-center" >
      <v-toolbar v-if="Boolean(request.Title)" dark :color="toolbarColor" 
        dense 
        elevation="4" 
      >
        <v-icon v-if="Boolean(icon)" left>stop</v-icon>
        <v-toolbar-title class="white--text">{{request.Title}}</v-toolbar-title>
      </v-toolbar>

      <v-card-text class="body-1 text-body-1 py-3" v-html="getTextFileContent()">
      </v-card-text>

      <v-card-text> 
            <v-progress-linear
              :value="progressPercent"
              :color="color"
              absolute
              height="7"
            ></v-progress-linear>
      </v-card-text>

    </v-card>
</template>

<script>
module.exports = {
    props: {
      request: Object,
    },
    data: function(){ return {
       progressPercent: 5,
       buttonTrueText: "Yes", 
       buttonFalseText: "No", 
       buttonTrueColor: "primary", 
       buttonFalseColor: "grey", 
       buttonFalseFlat: true, 
       buttonTrueFlat: true, 
       toolbarColor: "primary", 
       message: "Enter Yes or No", 
       persistent: Boolean,
       title: "Please Provide a Title", 
      }
    },
    computed: {
      icon: function() {
          return this.$vuetify.icons.values.warning
      },
      getTextFileContent: function() {
          textFile=this.request.TextFileName
          console.log("TODO: Fetch the text file from backend and display.")
          // Fetch it from backend.
      }
    },
    beforeUpdate: function() {
    },
    methods: {
       doneCallback: function() {
           var response = ":exit 0\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
    },
}

</script>
