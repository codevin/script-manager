import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import store from './store'
import axios from 'axios'
import moment from 'moment'
import Vuex from 'vuex'

Vue.config.productionTip = false


Vue.prototype.$EventBus=new Vue()
Vue.prototype.$axios = axios
Vue.prototype.$methods = {
	formatTime: function(value) {
             if (! value) value = (new Date()).getTime(); 
             if (value) {
               return moment(String(value)).format('HH:MM:SS')
             }
          },
          formatDate: function(value) {
             if (! value) value = (new Date()).getTime(); 
             if (value) {
               return moment(String(value)).format('MM/DD/YYYY')
             }
          },
          now: function() {
              return new Date()
          },
          timeout: function(value, msec) {
              var now = new Date()
              if (now - value > msec) return true
              return false 
          }
};

new Vue({
    render: h => h(App),
    store: store,
    data: () => ({ 
    }),
    Vuex, 
    vuetify,
    methods: {
          systemTimer1sec: function() {
             // This is called every 1 seconds. Any component can do their timed events.
             var now = (new Date()).getTime();
             this.$EventBus.$emit('timer1sec', {now: now})
          },
          systemTimer5sec: function() {
             var now = (new Date()).getTime();
             this.$EventBus.$emit('timer5sec', {now: now})
          },
          systemTimer30sec: function() {
             var now = (new Date()).getTime();
             this.$EventBus.$emit('timer30sec', {now: now})
          }
   },
   created: function() {
          setInterval(this.systemTimer1sec, 1000)
          setInterval(this.systemTimer5sec, 5000)
          setInterval(this.systemTimer30sec, 30000)
   }
}).$mount('#vueApp')


