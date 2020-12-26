
// message box (only ok button)
<template>
    <v-card class="justify-center" >
      <v-toolbar v-if="Boolean(request.Title)" dark 
        dense 
        elevation="4" 
      >
        <v-icon v-if="Boolean(icon)" left>stop</v-icon>
        <v-toolbar-title class="white--text">{{request.Title}}</v-toolbar-title>
      </v-toolbar>

      <v-card-text class="body-1 text-body-1 py-3" v-html="request.Message">
      </v-card-text>

      <v-card-text> 
            <v-progress-linear
              :value="request.PercentCompleted"
              color="red"
              absolute
              height="7"
            ></v-progress-linear>
      </v-card-text>

      <v-card-actions class="justify-center" >
        <v-btn
          v-if="enableDoneButton"
          color="blue"
          text
          @click="doneCallback"
        >
           Close 
        </v-btn>
      </v-card-actions>

    </v-card>
</template>

<script>
module.exports = {
    props: {
      request: Object,
    },
    data: function(){ return {
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
       enableDoneButton: false,
      }
    },
    computed: {
      icon: function() {
          return this.$vuetify.icons.values.warning
      },
    },
    beforeUpdate: function() {
    },
    methods: {
       enableDoneCallback: function() {
           this.enableDoneButton=true
       },
       doneCallback: function() {
           var response = ":exit 0\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
       progressUpdateCallback: function() {
           var response = ":exit 0\n" 
           this.$EventBus.$emit('updatedUI', response)
      }
    },
    watch: {
    },
}

</script>
