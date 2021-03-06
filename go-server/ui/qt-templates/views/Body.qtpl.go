// Code generated by qtc from "Body.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Body.qtpl:1
package views

//line views/Body.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Body.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Body.qtpl:1
func StreamBody(qw422016 *qt422016.Writer) {
//line views/Body.qtpl:1
	qw422016.N().S(`
  <body>
    <!-- ===============================================-->
    <!--    Main Content-->
    <!-- ===============================================-->
    <main class="main">
      <h1>Hello - Testing icon. <i class="fa fa-bullhorn"></i>
       <i class="material-icons">face</i></h1> 

      `)
//line views/Body.qtpl:10
	StreamTopNavigation(qw422016)
//line views/Body.qtpl:10
	qw422016.N().S(`
      <div id="vueApp">
         <v-app>
            <v-main>
               <v-container>
	           <div class="row align-items-center">
	        		<h2>Script Vue Table Interaction</h2>
                   </div>
	           <div class="row align-items-center">
                                `)
//line views/Body.qtpl:19
	StreamScriptsTableVue(qw422016)
//line views/Body.qtpl:19
	qw422016.N().S(`
	           </div>
               </v-container>
            </v-main>
         </v-app>
      </div>

      <!-- ============================================-->
      <section class="text-center border-top footer-text py-3">

        <!-- <div class="container">Made with ❤️ by <a class="text-500 text-decoration-none">Me </a> -->
        <!-- </div> -->
        <!-- end of .container-->

      </section>
      <!-- <section> close ============================-->
      <!-- ============================================-->


    </main>

    <!-- ===============================================-->
    <!--    JavaScripts-->
    <!-- ===============================================-->
    <!-- JavaScript and dependencies -->
    <script src="_assets/js/jquery-3.5.1.min.js"></script>
    <script src="_assets/js/popper.min.js"></script>
    <script src="_assets/bootstrap-5.0.0-alpha1-dist/js/bootstrap.min.js"></script>
    <!-- <script src="https://stackpath.bootstrapcdn.com/bootstrap/5.0.0-alpha1/js/bootstrap.min.js" integrity="sha384-oesi62hOLfzrys4LxRF63OJCXdXDipiYWBnvTl9Y9/TRlw5xlKIEHpNyvvDShgf/" crossorigin="anonymous"></script> 
    <script src="_assets/js/bootstrap-5.0.0.min.js"></script> -->
    <script src="_assets/js/plugins.js"></script>
    <script src="_assets/lib/is/is.min.js"></script>
    <script src="_assets/lib/stickyfilljs/stickyfill.min.js"></script>
    <script src="_assets/js/theme.js"></script>
    <script src="_assets/js/jsrender.min.js"></script>
    <script src="_assets/js/moment.min.js"></script>

    <!-- For Vue Modules. -->
    <script type="text/javascript" src="vue/libs/vue2.js"></script>
    <script type="text/javascript" src="vue/libs/vuex-3.6.0.js"></script>
    <script type="text/javascript" src="vue/libs/vuetify.min.js"></script>
    <script type="text/javascript" src="vue/libs/axios.min.js"></script>
    <script type="text/javascript" src="vue/libs/gcare-httpVueLoader-simplified.js"></script>
    <script type="text/javascript" src="vue/libs/gcare-httpVueLoader-addon.js"></script>
    <script type="text/javascript" src="vue/vuejs-dialog-orig/dist/vuejs-dialog.min.js"></script>

   <!-- Finally our standard code if any -->
   <script src="_assets/js/index-cleaned.js"></script>

   <!-- Whatever is generated by template system. -->
   <script src="_assets/js/index-autogen.js"></script> 

   <script>

    `)
//line views/Body.qtpl:73
	StreamVueComponent_ScriptsTable(qw422016)
//line views/Body.qtpl:73
	qw422016.N().S(`
    `)
//line views/Body.qtpl:74
	StreamJS_StoreVue(qw422016)
//line views/Body.qtpl:74
	qw422016.N().S(`

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
          now: function(value, msec) {
              return new Date()
          },
          timeout: function(value, msec) {
              var now = new Date()
              if (now - value > msec) return true
              return false 
          }
    };

    new Vue({
      el: '#vueApp',
      store: store,
      data: () => ({ 
      }),
      vuetify: new Vuetify({
         icons: {
            iconfont: 'fa'
         }
      }),
      methods: {
          systemTimer1sec: function() {
             // This is called every 1 seconds. Any component can do their timed events.
             now = (new Date()).getTime();
             this.$EventBus.$emit('timer1sec', {now: now})
          },
          systemTimer5sec: function() {
             now = (new Date()).getTime();
             this.$EventBus.$emit('timer5sec', {now: now})
          },
          systemTimer30sec: function() {
             now = (new Date()).getTime();
             this.$EventBus.$emit('timer30sec', {now: now})
          }
      },
      created: function() {
          setInterval(this.systemTimer1sec, 1000)
          setInterval(this.systemTimer5sec, 5000)
          setInterval(this.systemTimer30sec, 30000)
      }
    })

   </script>

  </body>
`)
//line views/Body.qtpl:136
}

//line views/Body.qtpl:136
func WriteBody(qq422016 qtio422016.Writer) {
//line views/Body.qtpl:136
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Body.qtpl:136
	StreamBody(qw422016)
//line views/Body.qtpl:136
	qt422016.ReleaseWriter(qw422016)
//line views/Body.qtpl:136
}

//line views/Body.qtpl:136
func Body() string {
//line views/Body.qtpl:136
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Body.qtpl:136
	WriteBody(qb422016)
//line views/Body.qtpl:136
	qs422016 := string(qb422016.B)
//line views/Body.qtpl:136
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Body.qtpl:136
	return qs422016
//line views/Body.qtpl:136
}

//line views/Body.qtpl:139
func StreamJSCodeForBody(qw422016 *qt422016.Writer) {
//line views/Body.qtpl:139
	qw422016.N().S(`
 <!-- All JS functions which are related to main body. -->

`)
//line views/Body.qtpl:142
}

//line views/Body.qtpl:142
func WriteJSCodeForBody(qq422016 qtio422016.Writer) {
//line views/Body.qtpl:142
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Body.qtpl:142
	StreamJSCodeForBody(qw422016)
//line views/Body.qtpl:142
	qt422016.ReleaseWriter(qw422016)
//line views/Body.qtpl:142
}

//line views/Body.qtpl:142
func JSCodeForBody() string {
//line views/Body.qtpl:142
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Body.qtpl:142
	WriteJSCodeForBody(qb422016)
//line views/Body.qtpl:142
	qs422016 := string(qb422016.B)
//line views/Body.qtpl:142
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Body.qtpl:142
	return qs422016
//line views/Body.qtpl:142
}

//line views/Body.qtpl:145
func StreamCSSCodeForBody(qw422016 *qt422016.Writer) {
//line views/Body.qtpl:145
	qw422016.N().S(`
`)
//line views/Body.qtpl:146
}

//line views/Body.qtpl:146
func WriteCSSCodeForBody(qq422016 qtio422016.Writer) {
//line views/Body.qtpl:146
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Body.qtpl:146
	StreamCSSCodeForBody(qw422016)
//line views/Body.qtpl:146
	qt422016.ReleaseWriter(qw422016)
//line views/Body.qtpl:146
}

//line views/Body.qtpl:146
func CSSCodeForBody() string {
//line views/Body.qtpl:146
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Body.qtpl:146
	WriteCSSCodeForBody(qb422016)
//line views/Body.qtpl:146
	qs422016 := string(qb422016.B)
//line views/Body.qtpl:146
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Body.qtpl:146
	return qs422016
//line views/Body.qtpl:146
}
