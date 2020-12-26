// Code generated by qtc from "DropDownMenu.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line widgets/DropDownMenu.qtpl:1
package widgets

//line widgets/DropDownMenu.qtpl:1
import . "../types"

//line widgets/DropDownMenu.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line widgets/DropDownMenu.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line widgets/DropDownMenu.qtpl:3
func StreamDropDownMenu(qw422016 *qt422016.Writer, opt *WidgetInfo) {
//line widgets/DropDownMenu.qtpl:3
	qw422016.N().S(`
  <div class="dropdown">
    <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown">
      `)
//line widgets/DropDownMenu.qtpl:6
	qw422016.E().S(opt.ButtonText)
//line widgets/DropDownMenu.qtpl:6
	qw422016.N().S(` 
    </button>
    <div class="dropdown-menu">
`)
//line widgets/DropDownMenu.qtpl:9
	for _, item := range opt.Items {
//line widgets/DropDownMenu.qtpl:9
		qw422016.N().S(`
      <div class="dropdown-item" item="`)
//line widgets/DropDownMenu.qtpl:10
		qw422016.E().S(item.Value)
//line widgets/DropDownMenu.qtpl:10
		qw422016.N().S(`">`)
//line widgets/DropDownMenu.qtpl:10
		qw422016.E().S(item.Value)
//line widgets/DropDownMenu.qtpl:10
		qw422016.N().S(`</div>
`)
//line widgets/DropDownMenu.qtpl:11
	}
//line widgets/DropDownMenu.qtpl:11
	qw422016.N().S(`
    </div>
  </div>
`)
//line widgets/DropDownMenu.qtpl:14
}

//line widgets/DropDownMenu.qtpl:14
func WriteDropDownMenu(qq422016 qtio422016.Writer, opt *WidgetInfo) {
//line widgets/DropDownMenu.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line widgets/DropDownMenu.qtpl:14
	StreamDropDownMenu(qw422016, opt)
//line widgets/DropDownMenu.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line widgets/DropDownMenu.qtpl:14
}

//line widgets/DropDownMenu.qtpl:14
func DropDownMenu(opt *WidgetInfo) string {
//line widgets/DropDownMenu.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line widgets/DropDownMenu.qtpl:14
	WriteDropDownMenu(qb422016, opt)
//line widgets/DropDownMenu.qtpl:14
	qs422016 := string(qb422016.B)
//line widgets/DropDownMenu.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line widgets/DropDownMenu.qtpl:14
	return qs422016
//line widgets/DropDownMenu.qtpl:14
}