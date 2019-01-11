// +build cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-text.html.

func (ctx *Context) GetScaledFont() *ScaledFont {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_scaled_font(c_ctx)
	retGo := ScaledFontNewFromC(unsafe.Pointer(retC))

	retGo.reference()
	runtime.SetFinalizer(retGo, func(o *ScaledFont) {
		o.destroy()
	})

	return retGo
}
