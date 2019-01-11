// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

func ImageSurfaceCreate(format Format, width int, height int) *Surface {
	c_format := (C.cairo_format_t)(format)
	c_width := (C.gint)(width)
	c_height := (C.gint)(height)

	retC := C.cairo_image_surface_create(c_format, c_width, c_height)
	retGo := SurfaceNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Surface) {
		o.destroy()
	})

	return retGo
}

//func ImageSurfaceCreateForData(data []byte, format Format, width int, height int, stride int) *Surface {
//	panic("Need to figure out a sensible API for the data.")
//
//	//c_data := (*C.uchar)(&data[0])
//	//c_format := (C.cairo_format_t)(format)
//	//c_width := (C.gint)(width)
//	//c_height := (C.gint)(height)
//	//c_stride := (C.gint)(stride)
//	//
//	//retC := C.cairo_image_surface_create_for_data(c_data, c_format, c_width, c_height, c_stride)
//	//retGo := SurfaceNewFromC(unsafe.Pointer(retC))
//	//
//	//runtime.SetFinalizer(retGo, func(o *Surface) {
//	//	o.Destroy()
//	//})
//	//
//	//return retGo
//}

func (surface *Surface) GetWidth() int {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_image_surface_get_width(c_surface)
	return int(retC)
}

func (surface *Surface) GetHeight() int {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_image_surface_get_height(c_surface)
	return int(retC)
}

func (surface *Surface) destroy() {
	c_surface := (*C.cairo_surface_t)(surface.ToC())
	C.cairo_surface_destroy(c_surface)
}

func (surface *Surface) reference() {
	c_surface := (*C.cairo_surface_t)(surface.ToC())
	C.cairo_surface_reference(c_surface)
}

// The user_data functions would be a little more complicatedd
// to support than most other cairo functions.
// And there should be little need for them in a Go application.
// Use of a closure will usually suffice.
//
// cairo_image_surface_get_data (cairo_surface_t *surface);
