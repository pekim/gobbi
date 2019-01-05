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
		o.Destroy()
	})

	return retGo
}

func ImageSurfaceCreateForData(data []byte, format Format, width int, height int, stride int) *Surface {
	panic("Need to figure out a sensible API for the data.")

	//c_data := (*C.uchar)(&data[0])
	//c_format := (C.cairo_format_t)(format)
	//c_width := (C.gint)(width)
	//c_height := (C.gint)(height)
	//c_stride := (C.gint)(stride)
	//
	//retC := C.cairo_image_surface_create_for_data(c_data, c_format, c_width, c_height, c_stride)
	//retGo := SurfaceNewFromC(unsafe.Pointer(retC))
	//
	//runtime.SetFinalizer(retGo, func(o *Surface) {
	//	o.Destroy()
	//})
	//
	//return retGo
}

func FormatStridForWidth(format Format, width int) int {
	c_format := (C.cairo_format_t)(format)
	c_width := (C.gint)(width)

	retC := C.cairo_format_stride_for_width(c_format, c_width)
	return int(retC)
}

func (surface *Surface) GetReferenceCount() int {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_surface_get_reference_count(c_surface)
	return int(retC)
}

func (surface *Surface) Destroy() {
	c_surface := (*C.cairo_surface_t)(surface.ToC())
	C.cairo_surface_destroy(c_surface)
}

//unsigned char *
//cairo_image_surface_get_data (cairo_surface_t *surface);
