// Code generated by qtc from "MessageBox.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// whiptail --title "Example Title" --msgbox "This is an example message box. Press OK to continue." 8 70
//

//line widgets/MessageBox.qtpl:3
package widgets

//line widgets/MessageBox.qtpl:3
import . "../types"

//line widgets/MessageBox.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line widgets/MessageBox.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line widgets/MessageBox.qtpl:5
func StreamMessageBox(qw422016 *qt422016.Writer, opt *WidgetInfo) {
//line widgets/MessageBox.qtpl:5
	qw422016.N().S(`
		      <div class="modal-header">
		        <h5 class="modal-title" id="`)
//line widgets/MessageBox.qtpl:7
	qw422016.E().S(opt.Id + "_title")
//line widgets/MessageBox.qtpl:7
	qw422016.N().S(`">`)
//line widgets/MessageBox.qtpl:7
	qw422016.E().S(opt.Title)
//line widgets/MessageBox.qtpl:7
	qw422016.N().S(`</h5>
		      </div>
		      <div class="modal-body">
		          <div class="row">
                               <div id="`)
//line widgets/MessageBox.qtpl:11
	qw422016.E().S(opt.Id)
//line widgets/MessageBox.qtpl:11
	qw422016.N().S(`" class="text-center">`)
//line widgets/MessageBox.qtpl:11
	qw422016.E().S(opt.Message)
//line widgets/MessageBox.qtpl:11
	qw422016.N().S(`</div>
                          </div>
		     </div> <!-- modal-body -->
                     <div class="modal-footer">
                          <button type="button" class="btn btn-success ok">`)
//line widgets/MessageBox.qtpl:15
	qw422016.E().S(opt.YesButtonText)
//line widgets/MessageBox.qtpl:15
	qw422016.N().S(`</button>
                     </div>
`)
//line widgets/MessageBox.qtpl:17
}

//line widgets/MessageBox.qtpl:17
func WriteMessageBox(qq422016 qtio422016.Writer, opt *WidgetInfo) {
//line widgets/MessageBox.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line widgets/MessageBox.qtpl:17
	StreamMessageBox(qw422016, opt)
//line widgets/MessageBox.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line widgets/MessageBox.qtpl:17
}

//line widgets/MessageBox.qtpl:17
func MessageBox(opt *WidgetInfo) string {
//line widgets/MessageBox.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line widgets/MessageBox.qtpl:17
	WriteMessageBox(qb422016, opt)
//line widgets/MessageBox.qtpl:17
	qs422016 := string(qb422016.B)
//line widgets/MessageBox.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line widgets/MessageBox.qtpl:17
	return qs422016
//line widgets/MessageBox.qtpl:17
}
