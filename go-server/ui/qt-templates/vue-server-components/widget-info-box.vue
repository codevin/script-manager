// message box (only ok button)
<template>
    <v-card class="justify-center">
      <v-toolbar v-if="Boolean(request.Title)" dark :color="toolbarColor" 
        dense 
        elevation="4" 
      >
        <v-icon v-if="Boolean(icon)" left>stop</v-icon>
        <v-toolbar-title class="white--text">{{request.Title}}</v-toolbar-title>
      </v-toolbar>

      <v-card-text class="body-1 text-body-1 py-3" v-html="request.Message">
      </v-card-text>

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
       timestampCreated: "",
       activeState: true,
      }
    },
    computed: {
      icon: function() {
          return this.$vuetify.icons.values.warning
      },
    },
    methods: {
       // Since it is updateable widget, there are no callbacks.
    },
}

</script>
