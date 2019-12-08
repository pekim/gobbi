// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
)

var pixbufErrorQuarkFunction *gi.Function
var pixbufErrorQuarkFunction_Once sync.Once

func pixbufErrorQuarkFunction_Set() error {
	var err error
	pixbufErrorQuarkFunction_Once.Do(func() {
		pixbufErrorQuarkFunction, err = gi.FunctionInvokerNew("GdkPixbuf", "pixbuf_error_quark")
	})
	return err
}

// PixbufErrorQuark is a representation of the C type gdk_pixbuf_error_quark.
func PixbufErrorQuark() glib.Quark {

	var ret gi.Argument

	err := pixbufErrorQuarkFunction_Set()
	if err == nil {
		ret = pixbufErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}
