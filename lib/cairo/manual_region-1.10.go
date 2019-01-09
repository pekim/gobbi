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

func (r *Region) GetExtents() RectangleInt {
	c_r := (*C.cairo_region_t)(r.ToC())
	var c_extents C.cairo_rectangle_int_t

	C.cairo_region_get_extents(c_r, &c_extents)

	return rectangleIntFromC(&c_extents)
}

func (r *Region) NumRectangles() int {
	c_r := (*C.cairo_region_t)(r.ToC())
	retC := C.cairo_region_num_rectangles(c_r)

	return int(retC)
}

func (r *Region) GetRectangle(nth int) RectangleInt {
	c_r := (*C.cairo_region_t)(r.ToC())
	var c_rectangle C.cairo_rectangle_int_t

	C.cairo_region_get_extents(c_r, &c_rectangle)
	return rectangleIntFromC(&c_rectangle)
}

func (r *Region) IsEmpty() bool {
	c_r := (*C.cairo_region_t)(r.ToC())

	retC := C.cairo_region_is_empty(c_r)
	return retC == 1
}

func (r *Region) ContainsPoint(x int, y int) bool {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_x := C.int(x)
	c_y := C.int(y)

	retC := C.cairo_region_contains_point(c_r, c_x, c_y)
	return retC == 1
}

func (r *Region) ContainsRectangle(rectangle RectangleInt) RegionOverlap {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_rectangle := rectangle.toC()

	retC := C.cairo_region_contains_rectangle(c_r, &c_rectangle)
	return RegionOverlap(retC)
}

func (r *Region) Equal(other *Region) bool {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_other := (*C.cairo_region_t)(other.ToC())

	retC := C.cairo_region_equal(c_r, c_other)
	return retC == 1
}

func (r *Region) Translate(dx int, dy int) {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_dx := C.int(dx)
	c_dy := C.int(dy)

	C.cairo_region_translate(c_r, c_dx, c_dy)
}

// Intersect updates the Region r with the intersection of
// r and other.
func (r *Region) Intersect(other *Region) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_other := (*C.cairo_region_t)(other.ToC())

	retC := C.cairo_region_intersect(c_r, c_other)
	return Status(retC)
}

// IntersectRectangle updates the Region r with the intersection of
// r and rectangle.
func (r *Region) IntersectRectangle(rectangle RectangleInt) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_rectangle := rectangle.toC()

	retC := C.cairo_region_intersect_rectangle(c_r, &c_rectangle)
	return Status(retC)
}

// Subtract updates the Region r with the result of subtracting
// other from r.
func (r *Region) Subtract(other *Region) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_other := (*C.cairo_region_t)(other.ToC())

	retC := C.cairo_region_subtract(c_r, c_other)
	return Status(retC)
}

// SubtractRectangle updates the Region r with the result of subtracting
// rectangle from r.
func (r *Region) SubtractRectangle(rectangle RectangleInt) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_rectangle := rectangle.toC()

	retC := C.cairo_region_subtract_rectangle(c_r, &c_rectangle)
	return Status(retC)
}

// Union updates the Region r with the union of
// r and other.
func (r *Region) Union(other *Region) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_other := (*C.cairo_region_t)(other.ToC())

	retC := C.cairo_region_union(c_r, c_other)
	return Status(retC)
}

// UnionRectangle updates the Region r with the union of
// r and rectangle.
func (r *Region) UnionRectangle(rectangle RectangleInt) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_rectangle := rectangle.toC()

	retC := C.cairo_region_union_rectangle(c_r, &c_rectangle)
	return Status(retC)
}

// Xor updates the Region r with the exclusive difference of
// r and other.
func (r *Region) Xor(other *Region) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_other := (*C.cairo_region_t)(other.ToC())

	retC := C.cairo_region_xor(c_r, c_other)
	return Status(retC)
}

// XorRectangle updates the Region r with the exclusive difference of
// r and rectangle.
func (r *Region) XorRectangle(rectangle RectangleInt) Status {
	c_r := (*C.cairo_region_t)(r.ToC())
	c_rectangle := rectangle.toC()

	retC := C.cairo_region_xor_rectangle(c_r, &c_rectangle)
	return Status(retC)
}
