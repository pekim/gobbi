// This is a generated file - DO NOT EDIT

package cairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

// ImageSurfaceCreate is a wrapper around the C function cairo_image_surface_create.
func ImageSurfaceCreate() {
	C.cairo_image_surface_create()

	return
}
