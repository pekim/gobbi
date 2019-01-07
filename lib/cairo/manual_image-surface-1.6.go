// +build cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

func FormatStrideForWidth(format Format, width int) int {
	c_format := (C.cairo_format_t)(format)
	c_width := (C.gint)(width)

	retC := C.cairo_format_stride_for_width(c_format, c_width)
	return int(retC)
}
