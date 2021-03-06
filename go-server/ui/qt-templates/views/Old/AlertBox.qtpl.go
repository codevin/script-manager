// Code generated by qtc from "AlertBox.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Old/AlertBox.qtpl:1
package Old

//line views/Old/AlertBox.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Old/AlertBox.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Old/AlertBox.qtpl:1
func StreamAlertBoxModal(qw422016 *qt422016.Writer) {
//line views/Old/AlertBox.qtpl:1
	qw422016.N().S(`
    <div class="fixed_alert alert-danger alert-dismissible fade in" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span>
        </button>
        <h4>Oh snap! You got an error!</h4>
        <p id="SystemAlertMessage">Fixed Alert Message</p>
   </div>
`)
//line views/Old/AlertBox.qtpl:8
}

//line views/Old/AlertBox.qtpl:8
func WriteAlertBoxModal(qq422016 qtio422016.Writer) {
//line views/Old/AlertBox.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/AlertBox.qtpl:8
	StreamAlertBoxModal(qw422016)
//line views/Old/AlertBox.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line views/Old/AlertBox.qtpl:8
}

//line views/Old/AlertBox.qtpl:8
func AlertBoxModal() string {
//line views/Old/AlertBox.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/AlertBox.qtpl:8
	WriteAlertBoxModal(qb422016)
//line views/Old/AlertBox.qtpl:8
	qs422016 := string(qb422016.B)
//line views/Old/AlertBox.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/AlertBox.qtpl:8
	return qs422016
//line views/Old/AlertBox.qtpl:8
}

//line views/Old/AlertBox.qtpl:10
func StreamOkCancelDialogModal(qw422016 *qt422016.Writer) {
//line views/Old/AlertBox.qtpl:10
	qw422016.N().S(`
    <!-- OKCancelDialogModal -->
	<div class="modal fade" id="okCancelDialogModal" tabindex="-1" role="dialog" aria-labelledby="timeSettingsModalLabel" aria-hidden="true">
	     <div class="modal-dialog" role="document">
		<div class="modal-content">
		      <div class="modal-header">
		        <h5 class="modal-title" id="okCancelDialogLabel">Confirmation</h5>
		      </div>
		      <div class="modal-body">
		          <div class="row">
                               <div id="okCancelMessage" class="text-center"> Message Comes Here</div>
                          </div>
		     </div> <!-- modal-body -->
                     <div class="modal-footer">
                          <button type="button" class="btn btn-success ok">Yes</button>
                          <button type="button" class="btn btn-default cancel">No</button>
                     </div>
		</div>
	    </div>
         </div>
`)
//line views/Old/AlertBox.qtpl:30
}

//line views/Old/AlertBox.qtpl:30
func WriteOkCancelDialogModal(qq422016 qtio422016.Writer) {
//line views/Old/AlertBox.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/AlertBox.qtpl:30
	StreamOkCancelDialogModal(qw422016)
//line views/Old/AlertBox.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line views/Old/AlertBox.qtpl:30
}

//line views/Old/AlertBox.qtpl:30
func OkCancelDialogModal() string {
//line views/Old/AlertBox.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/AlertBox.qtpl:30
	WriteOkCancelDialogModal(qb422016)
//line views/Old/AlertBox.qtpl:30
	qs422016 := string(qb422016.B)
//line views/Old/AlertBox.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/AlertBox.qtpl:30
	return qs422016
//line views/Old/AlertBox.qtpl:30
}

//line views/Old/AlertBox.qtpl:32
func StreamJavascript_AlertBox(qw422016 *qt422016.Writer) {
//line views/Old/AlertBox.qtpl:32
	qw422016.N().S(`

function systemAlert(message) {
   $("#SystemAlertMessage").text=message  
   $().alert();
   setTimeout(function(){$('.alert').alert('close')}, 3000);
}

function systemDialogOkCancel(message, callback) {
  //  put message into the body of modal
  var modal_id = '#okCancelDialogModal'

  $(modal_id).on('show.bs.modal', function (event) {
      //  store current modal reference
      var modal = $(this);
      //  update modal body help text
      modal.find('.modal-body #okCancelMessage').text(message);
  });

  //  show modal ringer custom confirmation
  $(modal_id).modal('show');

  $(modal_id + ' button.ok').off().on('click', function() {
     $(modal_id).modal('hide');
     callback(true);
  });

  $(modal_id + ' button.cancel').off().on('click', function() {
     $(modal_id).modal('hide');
     callback(false);
  });
}

`)
//line views/Old/AlertBox.qtpl:65
}

//line views/Old/AlertBox.qtpl:65
func WriteJavascript_AlertBox(qq422016 qtio422016.Writer) {
//line views/Old/AlertBox.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/AlertBox.qtpl:65
	StreamJavascript_AlertBox(qw422016)
//line views/Old/AlertBox.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line views/Old/AlertBox.qtpl:65
}

//line views/Old/AlertBox.qtpl:65
func Javascript_AlertBox() string {
//line views/Old/AlertBox.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/AlertBox.qtpl:65
	WriteJavascript_AlertBox(qb422016)
//line views/Old/AlertBox.qtpl:65
	qs422016 := string(qb422016.B)
//line views/Old/AlertBox.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/AlertBox.qtpl:65
	return qs422016
//line views/Old/AlertBox.qtpl:65
}

//line views/Old/AlertBox.qtpl:67
func StreamCSSForAlertBox(qw422016 *qt422016.Writer) {
//line views/Old/AlertBox.qtpl:67
	qw422016.N().S(`
.alert {
   position: fixed;
   top: 0px;
}
`)
//line views/Old/AlertBox.qtpl:72
}

//line views/Old/AlertBox.qtpl:72
func WriteCSSForAlertBox(qq422016 qtio422016.Writer) {
//line views/Old/AlertBox.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Old/AlertBox.qtpl:72
	StreamCSSForAlertBox(qw422016)
//line views/Old/AlertBox.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line views/Old/AlertBox.qtpl:72
}

//line views/Old/AlertBox.qtpl:72
func CSSForAlertBox() string {
//line views/Old/AlertBox.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Old/AlertBox.qtpl:72
	WriteCSSForAlertBox(qb422016)
//line views/Old/AlertBox.qtpl:72
	qs422016 := string(qb422016.B)
//line views/Old/AlertBox.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Old/AlertBox.qtpl:72
	return qs422016
//line views/Old/AlertBox.qtpl:72
}
