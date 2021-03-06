// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Note: Outputs of these methods are used by Go program directly to create
// a single file for each.
//

//line views/index.qtpl:4
package views

//line views/index.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/index.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/index.qtpl:4
func StreamHtmlForIndexPage(qw422016 *qt422016.Writer, title string) {
//line views/index.qtpl:4
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en-US" dir="ltr">
   `)
//line views/index.qtpl:7
	StreamHead(qw422016, title)
//line views/index.qtpl:7
	qw422016.N().S(` 
   `)
//line views/index.qtpl:8
	StreamBody(qw422016)
//line views/index.qtpl:8
	qw422016.N().S(` 
</html>
`)
//line views/index.qtpl:10
}

//line views/index.qtpl:10
func WriteHtmlForIndexPage(qq422016 qtio422016.Writer, title string) {
//line views/index.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.qtpl:10
	StreamHtmlForIndexPage(qw422016, title)
//line views/index.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line views/index.qtpl:10
}

//line views/index.qtpl:10
func HtmlForIndexPage(title string) string {
//line views/index.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.qtpl:10
	WriteHtmlForIndexPage(qb422016, title)
//line views/index.qtpl:10
	qs422016 := string(qb422016.B)
//line views/index.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.qtpl:10
	return qs422016
//line views/index.qtpl:10
}

//line views/index.qtpl:12
func StreamJSCodeForIndexPage(qw422016 *qt422016.Writer, title string) {
//line views/index.qtpl:12
	qw422016.N().S(`
   `)
//line views/index.qtpl:13
	StreamJSCodeForHead(qw422016)
//line views/index.qtpl:13
	qw422016.N().S(`
   `)
//line views/index.qtpl:14
	StreamJSCodeForBody(qw422016)
//line views/index.qtpl:14
	qw422016.N().S(`
`)
//line views/index.qtpl:15
}

//line views/index.qtpl:15
func WriteJSCodeForIndexPage(qq422016 qtio422016.Writer, title string) {
//line views/index.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.qtpl:15
	StreamJSCodeForIndexPage(qw422016, title)
//line views/index.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line views/index.qtpl:15
}

//line views/index.qtpl:15
func JSCodeForIndexPage(title string) string {
//line views/index.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.qtpl:15
	WriteJSCodeForIndexPage(qb422016, title)
//line views/index.qtpl:15
	qs422016 := string(qb422016.B)
//line views/index.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.qtpl:15
	return qs422016
//line views/index.qtpl:15
}

//line views/index.qtpl:17
func StreamCSSForIndexPage(qw422016 *qt422016.Writer, title string) {
//line views/index.qtpl:17
	qw422016.N().S(`
   `)
//line views/index.qtpl:18
	StreamCSSCodeForHead(qw422016)
//line views/index.qtpl:18
	qw422016.N().S(`
   `)
//line views/index.qtpl:19
	StreamCSSCodeForBody(qw422016)
//line views/index.qtpl:19
	qw422016.N().S(`
`)
//line views/index.qtpl:20
}

//line views/index.qtpl:20
func WriteCSSForIndexPage(qq422016 qtio422016.Writer, title string) {
//line views/index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.qtpl:20
	StreamCSSForIndexPage(qw422016, title)
//line views/index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line views/index.qtpl:20
}

//line views/index.qtpl:20
func CSSForIndexPage(title string) string {
//line views/index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.qtpl:20
	WriteCSSForIndexPage(qb422016, title)
//line views/index.qtpl:20
	qs422016 := string(qb422016.B)
//line views/index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.qtpl:20
	return qs422016
//line views/index.qtpl:20
}
