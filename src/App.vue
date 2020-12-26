<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-icon x-large >mdi-folder-open</v-icon>
        &nbsp;
        <h4>
         Script Manager
        </h4>
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/codevin/script-manager"
        target="_blank"
        text
      >
        <span class="mr-2">Github</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>

   <div id="VueScriptsTable">
       <scripts-table></scripts-table>
       <script-modal></script-modal>
   </div>

    </v-main>
  </v-app>
</template>

<script>
import ScriptsTable from './components/scripts-table.vue';
import ScriptModal from './components/script-modal.vue';

export default {
  name: 'App',

  components: {
    ScriptsTable,
    ScriptModal,
  },

  data: () => ({
    //
  }),
};
</script>
