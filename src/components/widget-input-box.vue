// Input box
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

      <v-card-text class="body-1 text-body-1 py-3">
      <v-form ref="form" v-model="valid" lazy-validation >
          <v-text-field
             v-model="inputValue"
             :counter="10"
             :rules="inputValueRules"
             label="Input Value"
             required
          >
          </v-text-field>
      </v-form>
      </v-card-text>

      <v-card-actions class="justify-center" >
        <v-btn
          color="blue"
          text
          @click="cancelCallback"
        >
          Cancel
        </v-btn>

        <v-btn
          color="blue"
          text
          @click="okCallback"
        >
          OK 
        </v-btn>

      </v-card-actions>

    </v-card>
</template>

<script>
module.exports = {
    props: {
      request: Object,
    },
    data: function(){
         return {
       valid: true,
       inputValue: this.request.DefaultValue, 
       inputValueRules: [
           v => !!v || 'Input Value is required',
           v => (v && v.length <= 60) || 'Value must be less than 10 characters',
       ],
       toolbarColor: "primary", 
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
       okCallback: function() {
           console.log("widget-checkbox: Selected Item is", this.selectedItems)
           var response={
               command: ":reply", 
               content: this.inputValue + "\n:exit 0\n",
               keepWidget: false,
           }
           this.$EventBus.$emit('finishedUI', response)
       },
       cancelCallback: function() {
           var response = ":exit 1\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
    }
}

</script>
