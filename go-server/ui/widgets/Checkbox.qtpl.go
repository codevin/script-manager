// Code generated by qtc from "Checkbox.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line widgets/Checkbox.qtpl:1
package widgets

//line widgets/Checkbox.qtpl:1
import . "../types"

//line widgets/Checkbox.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line widgets/Checkbox.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line widgets/Checkbox.qtpl:3
func StreamCheckBox(qw422016 *qt422016.Writer, opt *WidgetInfo) {
//line widgets/Checkbox.qtpl:3
	qw422016.N().S(`
`)
//line widgets/Checkbox.qtpl:4
	if len(opt.Items) == 0 {
//line widgets/Checkbox.qtpl:4
		qw422016.N().S(`
<div>No Options Provided </div>

`)
//line widgets/Checkbox.qtpl:7
	} else {
//line widgets/Checkbox.qtpl:7
		qw422016.N().S(`
`)
//line widgets/Checkbox.qtpl:8
		for _, item := range opt.Items {
//line widgets/Checkbox.qtpl:8
			qw422016.N().S(` 
<div class="form-check `)
//line widgets/Checkbox.qtpl:9
			if opt.ShowInline {
//line widgets/Checkbox.qtpl:9
				qw422016.N().S(` form-check-inline `)
//line widgets/Checkbox.qtpl:9
			}
//line widgets/Checkbox.qtpl:9
			qw422016.N().S(`">
  <input class="form-check-input" type="checkbox" value="`)
//line widgets/Checkbox.qtpl:10
			qw422016.E().S(item.Option)
//line widgets/Checkbox.qtpl:10
			qw422016.N().S(`" id="`)
//line widgets/Checkbox.qtpl:10
			qw422016.E().S(item.Id)
//line widgets/Checkbox.qtpl:10
			qw422016.N().S(`" `)
//line widgets/Checkbox.qtpl:10
			if item.Checked {
//line widgets/Checkbox.qtpl:10
				qw422016.N().S(`checked`)
//line widgets/Checkbox.qtpl:10
			}
//line widgets/Checkbox.qtpl:10
			qw422016.N().S(` `)
//line widgets/Checkbox.qtpl:10
			if item.Disabled {
//line widgets/Checkbox.qtpl:10
				qw422016.N().S(`disabled`)
//line widgets/Checkbox.qtpl:10
			}
//line widgets/Checkbox.qtpl:10
			qw422016.N().S(` >
  <label class="form-check-label" for="`)
//line widgets/Checkbox.qtpl:11
			qw422016.E().S(item.Id)
//line widgets/Checkbox.qtpl:11
			qw422016.N().S(`">`)
//line widgets/Checkbox.qtpl:11
			qw422016.E().S(item.Label)
//line widgets/Checkbox.qtpl:11
			qw422016.N().S(`</label>
</div>
`)
//line widgets/Checkbox.qtpl:13
		}
//line widgets/Checkbox.qtpl:13
		qw422016.N().S(`

`)
//line widgets/Checkbox.qtpl:15
	}
//line widgets/Checkbox.qtpl:15
	qw422016.N().S(`

`)
//line widgets/Checkbox.qtpl:17
}

//line widgets/Checkbox.qtpl:17
func WriteCheckBox(qq422016 qtio422016.Writer, opt *WidgetInfo) {
//line widgets/Checkbox.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line widgets/Checkbox.qtpl:17
	StreamCheckBox(qw422016, opt)
//line widgets/Checkbox.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line widgets/Checkbox.qtpl:17
}

//line widgets/Checkbox.qtpl:17
func CheckBox(opt *WidgetInfo) string {
//line widgets/Checkbox.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line widgets/Checkbox.qtpl:17
	WriteCheckBox(qb422016, opt)
//line widgets/Checkbox.qtpl:17
	qs422016 := string(qb422016.B)
//line widgets/Checkbox.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line widgets/Checkbox.qtpl:17
	return qs422016
//line widgets/Checkbox.qtpl:17
}
