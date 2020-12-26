// Code generated by qtc from "ScriptsTable.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Old/ScriptsTable.qtpl:1
package Old

//line views/Old/ScriptsTable.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Old/ScriptsTable.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Old/ScriptsTable.qtpl:1
func streammodal_invoke_button(qw422016 *qt422016.Writer, name string, modal_id string, icon string) {
//line views/Old/ScriptsTable.qtpl:1
	qw422016.N().S(`
            <button type="button" class="btn btn-sky  btn-action" data-toggle="modal" data-target="#`)
//line views/Old/ScriptsTable.qtpl:2
	qw422016.E().S(modal_id)
//line views/Old/ScriptsTable.qtpl:2
	qw422016.N().S(`"><i class="fa `)
//line views/Old/ScriptsTable.qtpl:2
	qw422016.E().S(icon)
//line views/Old/ScriptsTable.qtpl:2
	qw422016.N().S(`"></i> `)
//line views/Old/ScriptsTable.qtpl:2
	qw422016.E().S(name)
//line views/Old/ScriptsTable.qtpl:2
	qw422016.N().S(`</button>
`)
//line views/Old/ScriptsTable.qtpl:3
}

//line views/Old/ScriptsTable.qtpl:3
func writemodal_invoke_button(qq422016 qtio422016.Writer, name string, modal_id string, icon string) {
//line views/Old/ScriptsTable.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/ScriptsTable.qtpl:3
	streammodal_invoke_button(qw422016, name, modal_id, icon)
//line views/Old/ScriptsTable.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line views/Old/ScriptsTable.qtpl:3
}

//line views/Old/ScriptsTable.qtpl:3
func modal_invoke_button(name string, modal_id string, icon string) string {
//line views/Old/ScriptsTable.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/ScriptsTable.qtpl:3
	writemodal_invoke_button(qb422016, name, modal_id, icon)
//line views/Old/ScriptsTable.qtpl:3
	qs422016 := string(qb422016.B)
//line views/Old/ScriptsTable.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/ScriptsTable.qtpl:3
	return qs422016
//line views/Old/ScriptsTable.qtpl:3
}

//line views/Old/ScriptsTable.qtpl:5
func streamsimple_button(qw422016 *qt422016.Writer, name, button_id, buttonFn, icon string) {
//line views/Old/ScriptsTable.qtpl:5
	qw422016.N().S(`
            <button type="button" id="`)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.E().S(button_id)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.N().S(`" class="btn btn-sky  btn-`)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.E().S(icon)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.N().S(`" onclick='`)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.E().S(buttonFn)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.N().S(`'><i class="fa `)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.E().S(icon)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.N().S(`"></i> `)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.E().S(name)
//line views/Old/ScriptsTable.qtpl:6
	qw422016.N().S(`</button>
`)
//line views/Old/ScriptsTable.qtpl:7
}

//line views/Old/ScriptsTable.qtpl:7
func writesimple_button(qq422016 qtio422016.Writer, name, button_id, buttonFn, icon string) {
//line views/Old/ScriptsTable.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/ScriptsTable.qtpl:7
	streamsimple_button(qw422016, name, button_id, buttonFn, icon)
//line views/Old/ScriptsTable.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line views/Old/ScriptsTable.qtpl:7
}

//line views/Old/ScriptsTable.qtpl:7
func simple_button(name, button_id, buttonFn, icon string) string {
//line views/Old/ScriptsTable.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/ScriptsTable.qtpl:7
	writesimple_button(qb422016, name, button_id, buttonFn, icon)
//line views/Old/ScriptsTable.qtpl:7
	qs422016 := string(qb422016.B)
//line views/Old/ScriptsTable.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/ScriptsTable.qtpl:7
	return qs422016
//line views/Old/ScriptsTable.qtpl:7
}

//line views/Old/ScriptsTable.qtpl:9
func StreamScriptsTable(qw422016 *qt422016.Writer) {
//line views/Old/ScriptsTable.qtpl:9
	qw422016.N().S(`
`)
//line views/Old/ScriptsTable.qtpl:10
}

//line views/Old/ScriptsTable.qtpl:10
func WriteScriptsTable(qq422016 qtio422016.Writer) {
//line views/Old/ScriptsTable.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/ScriptsTable.qtpl:10
	StreamScriptsTable(qw422016)
//line views/Old/ScriptsTable.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line views/Old/ScriptsTable.qtpl:10
}

//line views/Old/ScriptsTable.qtpl:10
func ScriptsTable() string {
//line views/Old/ScriptsTable.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/ScriptsTable.qtpl:10
	WriteScriptsTable(qb422016)
//line views/Old/ScriptsTable.qtpl:10
	qs422016 := string(qb422016.B)
//line views/Old/ScriptsTable.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/ScriptsTable.qtpl:10
	return qs422016
//line views/Old/ScriptsTable.qtpl:10
}

//line views/Old/ScriptsTable.qtpl:12
func StreamScriptsTable_DISABLE(qw422016 *qt422016.Writer) {
//line views/Old/ScriptsTable.qtpl:12
	qw422016.N().S(`
          <div class="row align-items-center">
            <h3>Scripts
            `)
//line views/Old/ScriptsTable.qtpl:15
	streamsimple_button(qw422016, "Refresh", "refreshAllScripts", "updateScriptsTable();", "fa-plus")
//line views/Old/ScriptsTable.qtpl:15
	qw422016.N().S(`
            </h3>
            <div class="input-group">
            	<div class="input-group-prepend">
		    <div class="input-group-text"><i class="fa fa-search"></i></div>
		    </div>
          	</div>
          	<table id="scriptsTable" class="table table-condensed">
			  <thead>
			    <tr>
			      <th scope="col"></th>
			      <th scope="col">Name</th>
			      <th scope="col">Actions</th>
			      <th scope="col">Status/Substatus</th>
			      <th scope="col">Duration (sec)</th>
			    </tr>
			  </thead>
			  <tbody id="scriptsTableBody">
			  </tbody>
			</table>
          </div>
`)
//line views/Old/ScriptsTable.qtpl:36
}

//line views/Old/ScriptsTable.qtpl:36
func WriteScriptsTable_DISABLE(qq422016 qtio422016.Writer) {
//line views/Old/ScriptsTable.qtpl:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/ScriptsTable.qtpl:36
	StreamScriptsTable_DISABLE(qw422016)
//line views/Old/ScriptsTable.qtpl:36
	qt422016.ReleaseWriter(qw422016)
//line views/Old/ScriptsTable.qtpl:36
}

//line views/Old/ScriptsTable.qtpl:36
func ScriptsTable_DISABLE() string {
//line views/Old/ScriptsTable.qtpl:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/ScriptsTable.qtpl:36
	WriteScriptsTable_DISABLE(qb422016)
//line views/Old/ScriptsTable.qtpl:36
	qs422016 := string(qb422016.B)
//line views/Old/ScriptsTable.qtpl:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/ScriptsTable.qtpl:36
	return qs422016
//line views/Old/ScriptsTable.qtpl:36
}

//line views/Old/ScriptsTable.qtpl:39
func StreamJavascript_ScriptsTable_ORIG(qw422016 *qt422016.Writer) {
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`

let activeScript = ""
let _scripts = new Map() 

let html_templates = {
    refreshButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="refreshScriptStatus('{{:Name}}');" type="button" class="btn btn-info btn-sm">Status Update</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    startButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="startScript('{{:Name}}');" type="button" class="btn btn-warning btn-sm">Start</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    interactButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="showScriptUI('{{:Name}}');" type="button" class="btn btn-warning btn-sm">UI</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    forcestopButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="forcestopScript('{{:Name}}');" type="button" class="btn btn-danger btn-sm">Force Stop</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    stopButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="stopScript('{{:Name}}');" type="button" class="btn btn-danger btn-sm">Stop</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    resetButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="resetScript('{{:Name}}');" type="button" class="btn btn-danger btn-sm">Reset</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    logButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="logFromScript('{{:Name}}');" type="button" class="btn btn-info btn-sm">Logs</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),
    clearlogsButton: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<button onclick="clearlogsFromScript('{{:Name}}');" type="button" class="btn btn-info btn-sm">ClearLogs</button>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`),

    script_row: $.templates(`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`<tr id="{{:ScriptRowId}}">
               <td class="align-middle"></td>
               <td class="align-middle">{{:Name}}</td>
               <td class="align-middle">
                  {{:refreshButton}}
                  {{:resetButton}}
                  {{:startButton}}
                  {{:interactButton}}
                  {{:stopButton}}
                  {{:forcestopButton}}
                  {{:logButton}}
                  {{:clearlogsButton}}
               </td>
               <td class="align-left">{{:Status.Status}} - {{:Status.SubStatus}}</td>
               <td class="align-middle">{{:Duration}}</td>
          </tr>`)
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S("`")
//line views/Old/ScriptsTable.qtpl:39
	qw422016.N().S(`)
}


function refreshScriptRow(script_name) {
       var button_renders = {}
       var s = _scripts.get(script_name)

       button_renders.startButton = "" 
       button_renders.stopButton = "" 
       button_renders.forcestopButton = "" 
       button_renders.resetButton = ""
       button_renders.interactButton = "" 
       if ( s.Status.Status == "" || s.Status.Status == "CannotStart") {
          button_renders.startButton = html_templates['startButton'].render(s)

       } else if ( s.Status.Status == "CannotStart") {
          button_renders.resetButton = html_templates['resetButton'].render(s)

       } else if ( s.Status.Status == "Stopped" ) {
          button_renders.resetButton = html_templates['resetButton'].render(s)

       } else if (s.Status.Status == "Running" ) {
          button_renders.stopButton = html_templates['stopButton'].render(s)
          button_renders.forcestopButton = html_templates['forcestopButton'].render(s)
          if ( s.Interaction.RequestState == "HaveRequest" )  {
              button_renders.interactButton = html_templates['interactButton'].render(s)
          }

       } else if (s.Status.Status == "Error" ) {
          button_renders.stopButton = html_templates['stopButton'].render(s)
          button_renders.forcestopButton = html_templates['forcestopButton'].render(s)
          button_renders.interactButton = html_templates['interactButton'].render(s)
       }
       button_renders.refreshButton = html_templates['refreshButton'].render(s)

       if ( s.Logs || s.Interaction.HaveLogs ) { 
           button_renders.logButton = html_templates['logButton'].render(s)
       }
       if ( s.Logs ) { 
           button_renders.clearlogButton = html_templates['clearlogButton'].render(s)
       }

       var row_id = "#ScriptRowId_"+script_name

       var contents = { 
            ScriptRowId : row_id, 
            ...s, 
            ...button_renders 
       }
       var rowHTML = html_templates['script_row'].render(contents)

       var el = document.createElement("tr")
       el.innerHTML = rowHTML 

       console.log("Updating Table row: ", row_id)

       // Find and replace displayed row if there is one.
       if (  $(row_id).val()  ) {
          console.log("Found row in table.")
          var row = $("table tbody").find(row_id);
          if (row) {
            console.log("Replacing.")
            row.replaceWith(rowHTML);
          }  else {
            console.log("Did not replace.")
          }
       } else {
          console.log("Row not found in table. Appending new.")
          console.log($("#scriptsTableBody").innerHtml)
          $("#scriptsTableBody").append(el)
       }
}


var updateFlag = false 

function refreshRow(script_name) {
   _scripts.forEach(function(script, s_name) {
      if (s_name == script_name) { 
         refreshScriptRow(script_name)
      }
   })
}

function refreshScriptRowsOnUpdate() {
    if (updateFlag) { 
       refreshScriptRows()
    }
}

function refreshScriptRows() {
   console.log("Refresh of Scripts:")

   $("#scriptsTableBody").html("")

   _scripts.forEach(function(script, script_name) {
       console.log("Refresh of Script:", script_name, script)
       refreshScriptRow(script_name)
   })
}


function updateScriptsTable() {
    var successFn = function(data) {

	var updated_scripts = JSON.parse(data)

        script_names = Object.keys(updated_scripts) 
        if ( script_names.length == 0) {
            $("#scriptsTableBody").html("<div>No Scripts Available</div>")
            return
        }
        for (i in script_names) {
            var s = updated_scripts[script_names[i]]
            // New script?
            if ( ! _scripts.get(s.Name) ) { 
                var script = {
                   Name: s.Name
                }
                _scripts.set(s.Name, script)
            } 
            s2 = _scripts.get(s.Name)
            s2.CmdPath = s.CmdPath,
            s2.Args =  s.Args,
            s2.Status = s.Status,
            s2.Interaction = s.Interaction
            s2.Duration = s.Duration, 
            s2.ErrorMessage = s.ErrorMessage
        }
        refreshScriptRows()
    }
    var errorFn = function(errMsg) {
       // systemAlert("Error in calling update:" + errMsg)
    }

    formData=new FormData()
    // Ajax Submit
    $.ajax({url: "/update", type: "POST", cache: false, data: formData,
      processData: false, contentType: false, forceSync: false, cache: false,
      success: function(data) { 
          successFn(data); 
      },
      error: function(xhr, status, errMsg) { 
          errorFn(errMsg + ": " + xhr.ResponseText); 
      },
      complete: function() {},
    });
}

$(document).ready(function() {
    updateScriptsTable()

    setInterval(updateScriptsTable, 10000)   // 10 sec
    setInterval(refreshScriptRowsOnUpdate(), 1000)  // 1 sec only if anything changed.
});


`)
//line views/Old/ScriptsTable.qtpl:233
}

//line views/Old/ScriptsTable.qtpl:233
func WriteJavascript_ScriptsTable_ORIG(qq422016 qtio422016.Writer) {
//line views/Old/ScriptsTable.qtpl:233
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/ScriptsTable.qtpl:233
	StreamJavascript_ScriptsTable_ORIG(qw422016)
//line views/Old/ScriptsTable.qtpl:233
	qt422016.ReleaseWriter(qw422016)
//line views/Old/ScriptsTable.qtpl:233
}

//line views/Old/ScriptsTable.qtpl:233
func Javascript_ScriptsTable_ORIG() string {
//line views/Old/ScriptsTable.qtpl:233
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/ScriptsTable.qtpl:233
	WriteJavascript_ScriptsTable_ORIG(qb422016)
//line views/Old/ScriptsTable.qtpl:233
	qs422016 := string(qb422016.B)
//line views/Old/ScriptsTable.qtpl:233
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/ScriptsTable.qtpl:233
	return qs422016
//line views/Old/ScriptsTable.qtpl:233
}