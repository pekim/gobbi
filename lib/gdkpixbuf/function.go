// This is a generated file - DO NOT EDIT

package gdkpixbuf

import glib "github.com/pekim/gobbi/lib/glib"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufErrorQuark is a wrapper around the C function gdk_pixbuf_error_quark.
func PixbufErrorQuark() glib.Quark {
	retC := C.gdk_pixbuf_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}
