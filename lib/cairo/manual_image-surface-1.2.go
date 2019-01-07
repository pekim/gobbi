// +build cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

//unsigned char *
//cairo_image_surface_get_data (cairo_surface_t *surface);

func (surface *Surface) GetFormat() Format {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_image_surface_get_format(c_surface)
	return Format(retC)
}

func (surface *Surface) GetStride() Format {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_image_surface_get_stride(c_surface)
	return Format(retC)
}
