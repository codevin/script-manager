// Code generated by qtc from "InputBoxWithOkCancel.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// COLOR=$(whiptail --inputbox "What is your favorite Color?" 8 78 Blue --title "Example Dialog" 3>&1 1>&2 2>&3)
// Here Blue is default value.
//

//line widgets/InputBoxWithOkCancel.qtpl:4
package widgets

//line widgets/InputBoxWithOkCancel.qtpl:4
import . "../types"

//line widgets/InputBoxWithOkCancel.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line widgets/InputBoxWithOkCancel.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line widgets/InputBoxWithOkCancel.qtpl:6
func StreamInputBoxWithOkCancel(qw422016 *qt422016.Writer, opt *WidgetInfo) {
//line widgets/InputBoxWithOkCancel.qtpl:6
	qw422016.N().S(`
		          <div class="row">
                               <div id="widget_infobox_message" class="text-center">`)
//line widgets/InputBoxWithOkCancel.qtpl:8
	qw422016.E().S(opt.Message)
//line widgets/InputBoxWithOkCancel.qtpl:8
	qw422016.N().S(`</div>
                          </div>
		          <div class="row">
                             <input type="text" id="widget_InputBox_input" onkeyup="widgetInputBox();" value="`)
//line widgets/InputBoxWithOkCancel.qtpl:11
	qw422016.E().S(opt.Value)
//line widgets/InputBoxWithOkCancel.qtpl:11
	qw422016.N().S(`" placeholder="`)
//line widgets/InputBoxWithOkCancel.qtpl:11
	qw422016.E().S(opt.Placeholder)
//line widgets/InputBoxWithOkCancel.qtpl:11
	qw422016.N().S(`"></input>
                          </div>
		          <div class="row text-right">
                              <button type="button" class="btn btn-success ok" onclick="widgetInput_okPressed();">`)
//line widgets/InputBoxWithOkCancel.qtpl:14
	qw422016.E().S(opt.YesButtonText)
//line widgets/InputBoxWithOkCancel.qtpl:14
	qw422016.N().S(`</button>
                              <button type="button" class="btn btn-default cancel" onclick="widgetInput_cancelPressed();">`)
//line widgets/InputBoxWithOkCancel.qtpl:15
	qw422016.E().S(opt.NoButtonText)
//line widgets/InputBoxWithOkCancel.qtpl:15
	qw422016.N().S(`</button>
                         </div>
		</div>
	    </div>
         </div>
`)
//line widgets/InputBoxWithOkCancel.qtpl:20
}

//line widgets/InputBoxWithOkCancel.qtpl:20
func WriteInputBoxWithOkCancel(qq422016 qtio422016.Writer, opt *WidgetInfo) {
//line widgets/InputBoxWithOkCancel.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line widgets/InputBoxWithOkCancel.qtpl:20
	StreamInputBoxWithOkCancel(qw422016, opt)
//line widgets/InputBoxWithOkCancel.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line widgets/InputBoxWithOkCancel.qtpl:20
}

//line widgets/InputBoxWithOkCancel.qtpl:20
func InputBoxWithOkCancel(opt *WidgetInfo) string {
//line widgets/InputBoxWithOkCancel.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line widgets/InputBoxWithOkCancel.qtpl:20
	WriteInputBoxWithOkCancel(qb422016, opt)
//line widgets/InputBoxWithOkCancel.qtpl:20
	qs422016 := string(qb422016.B)
//line widgets/InputBoxWithOkCancel.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line widgets/InputBoxWithOkCancel.qtpl:20
	return qs422016
//line widgets/InputBoxWithOkCancel.qtpl:20
}
