// +build cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"
import "unsafe"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Regions.html.

func RegionCreate() *Region {
	retC := C.cairo_region_create()
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

func RegionCreateRectangle(rectangle RectangleInt) *Region {
	retC := C.cairo_region_create_rectangle(rectangle.toC())
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}
