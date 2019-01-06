// +build cairo_1.4 cairo_1.6 cairo_1.10

package cairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

func (surface *Surface) GetReferenceCount() int {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_surface_get_reference_count(c_surface)
	return int(retC)
}
