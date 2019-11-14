// Code generated - DO NOT EDIT.

package gdkpixbuf

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'gdk_pixbuf_error_quark' : non trivial function

var pixbufErrorQuarkInvoker *gi.FunctionInvoker

// PixbufErrorQuark is a representation of the C type gdk_pixbuf_error_quark.
func PixbufErrorQuark() {
	if pixbufErrorQuarkInvoker == nil {
		pixbufErrorQuarkInvoker = gi.FunctionInvokerNew("GdkPixbuf", "pixbuf_error_quark")
	}

	pixbufErrorQuarkInvoker.Call()
}
