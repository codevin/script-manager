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

      <v-list dense>
          <v-subheader>{{request.Message}}</v-subheader>
          <v-list-item-group
            v-model="selectedItemIdx"
            color="primary"
          >
            <v-list-item
              v-for="(item, i) in request.Items"
              :key="i"
              @click="menuOKCallback"
            >
              <v-list-item-content>
                <v-list-item-title v-text="item.Label"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
      </v-list>

      <v-card-actions class="justify-center" >
        <v-btn
          color="blue"
          text
          @click="menuCancelCallback"
        >
          Cancel
        </v-btn>

        <v-btn
          color="blue"
          text
          @click="menuOKCallback"
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
       selectedItemIdx: -1,
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
       menuOKCallback: function() {
           var selectedItem 
           if (this.selectedItemIdx >= 0) {
              selectedItem = this.request.Items[this.selectedItemIdx]
           } else {
              selectedItem = { Value: "" }
           }
           console.log("widget-menu-box: Selected Item is", this.selectedItemIdx)
           var response={
               command: ":reply", 
               content: selectedItem.Value + "\n:exit 0\n",
               keepWidget: false,
           }
           this.$EventBus.$emit('finishedUI', response)
       },
       menuCancelCallback: function() {
           console.log("widget-menu-box: No Selected Item. Esc pressed.")
           var response = ":exit 1\n" 
           this.$EventBus.$emit('finishedUI', response)
       },
    }
}

</script>
