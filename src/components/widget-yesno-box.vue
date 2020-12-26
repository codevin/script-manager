<template>
    <v-card class="justify-center" >
      <v-toolbar v-if="Boolean(request.Title)" dark :color="toolbarColor" 
        dense 
        elevation="4" 
      >
        <v-icon v-if="Boolean(icon)" left>mdi-stop</v-icon>
        <v-toolbar-title class="white--text">{{request.Title}}</v-toolbar-title>
      </v-toolbar>

      <v-card-text class="body-1 text-body-1 py-3" v-html="request.Message">
      </v-card-text>

      <v-card-actions class="justify-center" >
        <v-btn
          color="blue"
          text
          @click="noCallback"
        >
          {{ request.NoButtonText }}
        </v-btn>

        <v-btn
          color="blue"
          text
          @click="yesCallback"
        >
          {{ request.YesButtonText }}
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
       yesCallback: function() {
           var response = ":exit 0\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
       noCallback: function() {
           var response = ":exit 1\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
    }
}

</script>
