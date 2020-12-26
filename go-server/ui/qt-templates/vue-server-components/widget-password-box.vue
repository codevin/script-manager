// Input box
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

      <v-card-text class="body-1 text-body-1 py-3">
      <v-form
          ref="form"
          v-model="valid"
          lazy-validation  >

          <v-text-field
            v-model="passwordValue"
            :append-icon="showPassword ? 'fa fa-eye' : 'fa fa-eye-slash'"
            :rules="[rules.required, rules.min]"
            :type="showPassword ? 'text' : 'password'"
            name="input-10-1"
            label="Normal with hint text"
            hint="At least 8 characters"
            counter
            @click:append="showPassword = !showPassword"
          ></v-text-field>

      <v-form>
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
       passwordValue: "", 
       showPassword: false,
       rules: {
          required: value => !!value || 'Required.',
          min: v => v.length >= 8 || 'Min 8 characters',
          emailMatch: () => (`The email and password you entered don't match`),
       },
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
           var response={
               command: ":reply", 
               content: this.passwordValue + "\n:exit 0\n",
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
