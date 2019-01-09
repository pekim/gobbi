// +build cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Regions.html.

func RegionCreate() *Region {
	retC := C.cairo_region_create()
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Region) {
		o.destroy()
	})

	return retGo
}

func RegionCreateRectangle(rectangle RectangleInt) *Region {
	c_rectangle := rectangle.toC()

	retC := C.cairo_region_create_rectangle(&c_rectangle)
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Region) {
		o.destroy()
	})

	return retGo
}

func RegionCreateRectangles(rectangles []RectangleInt) *Region {
	count := len(rectangles)
	c_count := C.int(count)

	c_rectangles := make([]C.cairo_rectangle_int_t, count, count)
	for i, rectangle := range rectangles {
		c_rectangles[i] = rectangle.toC()
	}

	retC := C.cairo_region_create_rectangles(&c_rectangles[0], c_count)
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Region) {
		o.destroy()
	})

	return retGo
}

func (r *Region) Copy() *Region {
	c_r := (*C.cairo_region_t)(r.ToC())

	retC := C.cairo_region_copy(c_r)
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Region) {
		o.destroy()
	})

	return retGo
}

func (r *Region) destroy() {
	c_r := (*C.cairo_region_t)(r.ToC())
	C.cairo_region_destroy(c_r)
}

func (r *Region) reference() {
	c_r := (*C.cairo_region_t)(r.ToC())
	C.cairo_region_reference(c_r)
}

func (r *Region) Status() Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	retC := C.cairo_region_status(c_r)

	return Status(retC)
}
