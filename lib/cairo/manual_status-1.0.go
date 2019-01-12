// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

func (s Status) String() string {
	c_s := C.cairo_status_t(s)
	retC := C.cairo_status_to_string(c_s)

	return C.GoString(retC)
}
