// Code generated by qtc from "TopNavigation.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/TopNavigation.qtpl:1
package views

//line views/TopNavigation.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/TopNavigation.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/TopNavigation.qtpl:1
func StreamTopNavigation(qw422016 *qt422016.Writer) {
//line views/TopNavigation.qtpl:1
	qw422016.N().S(`
      <nav class="navbar navbar-expand-lg navbar-light bg-white py-4 pl-0">

             <div class="container"><a class="navbar-brand text-primary font-weight-bold" href="index.html"> Script Interaction </a>

             <div class="btn-group emControl" style="">
                  <label>Start Script
	          <button class="btn btn-primary ml-auto hover-top-shadow" onclick="startScript();"><i class="fa fa-bullhorn"></i> </button> </label>

	          <button class="btn btn-primary ml-auto hover-top-shadow" onclick="stopScript();"><i class="fa fa-stop"></i> </button>
             </div>

             <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>
          
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav ml-auto">
                  <!-- <li class="nav-item text-muted"><a class="nav-link active" aria-current="page" href="index.html">Status</a></li> -->
                  <li class="nav-item"><a class="nav-link" href="#scripts">Scripts</a></li>
                  <li class="nav-item"><a class="nav-link" href="#assets">Sessions</a></li>
                  <li class="nav-item"><a class="nav-link" href="Logs.html">Logs</a></li>
                </ul>
          </div>
        </div>
      </nav>

`)
//line views/TopNavigation.qtpl:26
}

//line views/TopNavigation.qtpl:26
func WriteTopNavigation(qq422016 qtio422016.Writer) {
//line views/TopNavigation.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/TopNavigation.qtpl:26
	StreamTopNavigation(qw422016)
//line views/TopNavigation.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line views/TopNavigation.qtpl:26
}

//line views/TopNavigation.qtpl:26
func TopNavigation() string {
//line views/TopNavigation.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/TopNavigation.qtpl:26
	WriteTopNavigation(qb422016)
//line views/TopNavigation.qtpl:26
	qs422016 := string(qb422016.B)
//line views/TopNavigation.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/TopNavigation.qtpl:26
	return qs422016
//line views/TopNavigation.qtpl:26
}