// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// ShowUriOnWindow is a wrapper around the C function gtk_show_uri_on_window.
func ShowUriOnWindow(parent *Window, uri string, timestamp uint32) (bool, error) {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_timestamp := (C.guint32)(timestamp)

	var cThrowableError *C.GError

	retC := C.gtk_show_uri_on_window(c_parent, c_uri, c_timestamp, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}
