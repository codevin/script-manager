// Checklist box
<template>
    <v-card class="justify-center" >
      <v-toolbar v-if="Boolean(request.Title)" dark :color="toolbarColor" 
        dense 
        elevation="4" 
      >
        <v-icon v-if="Boolean(icon)" left>stop</v-icon>
        <v-toolbar-title class="white--text">{{request.Title}}</v-toolbar-title>
      </v-toolbar>

      <v-card-text class="body-1 text-body-1 py-3" v-html="request.Message">
      </v-card-text>

      <v-card-text><p>Selected: {{ selectedItems }}</p></v-card-text>
      <v-list dense>
          <v-checkbox
              v-model="selectedItems"
              on-icon="fa fa-check-square"
              off-icon="fa fa-square-o"
              :label="item.Label"
              :value="item.Value"
              v-for="(item, i) in request.Items"
              :key="i"
              hide-details
          >
          </v-checkbox>
      </v-list>

      <v-card-actions class="justify-center" >
        <v-btn
          color="blue"
          text
          @click="checkboxCancelCallback"
        >
          Cancel
        </v-btn>

        <v-btn
          color="blue"
          text
          @click="checkboxOKCallback"
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
       selectedItems: [],
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
       checkboxOKCallback: function() {
           console.log("widget-checkbox: Sending response. Selected Item is", this.selectedItems)
           var response={
               command: ":reply", 
               content: this.selectedItems.join(",") + "\n:exit 0\n",
               keepWidget: false,
           }
           this.$EventBus.$emit('finishedUI', response)
       },
       checkboxCancelCallback: function() {
           console.log("widget-checkbox: No Selected Item. Esc pressed.")
           var response = ":exit 1\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
    }
}

</script>
