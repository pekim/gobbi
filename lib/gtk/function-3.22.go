// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// This is a convenience function for launching the default application
// to show the uri. The uri must be of a form understood by GIO (i.e. you
// need to install gvfs to get support for uri schemes such as http://
// or ftp://, as only local files are handled by GIO itself).
// Typical examples are
// - `file:///home/gnome/pict.jpg`
// - `http://www.gnome.org`
// - `mailto:me@gnome.org`
//
// Ideally the timestamp is taken from the event triggering
// the gtk_show_uri() call. If timestamp is not known you can take
// %GDK_CURRENT_TIME.
//
// This is the recommended call to be used as it passes information
// necessary for sandbox helpers to parent their dialogs properly.
/*

C function

gtk_show_uri_on_window
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
