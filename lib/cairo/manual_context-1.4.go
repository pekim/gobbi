// +build cairo_1.4 cairo_1.6 cairo_1.10

package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-t.html.

func (ctx *Context) GetDashCount() int {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_dash_count(c_ctx)
	return int(retC)
}

func (ctx *Context) GetDash() ([]float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	dashCount := ctx.GetDashCount()
	c_dashes := make([]C.double, dashCount, dashCount)
	var c_offset C.double

	C.cairo_get_dash(c_ctx, &c_dashes[0], &c_offset)
	dashes := (*[1 << 30]float64)(unsafe.Pointer(&c_dashes[0]))[:dashCount:dashCount]

	return dashes, float64(c_offset)
}

func (ctx *Context) ClipExtents() (float64, float64, float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_x1 C.double
	var c_y1 C.double
	var c_x2 C.double
	var c_y2 C.double

	C.cairo_clip_extents(c_ctx, &c_x1, &c_y1, &c_x2, &c_y2)

	return float64(c_x1), float64(c_y1), float64(c_x2), float64(c_y2)
}

// cairo_rectangle_list_destroy is not exposed through an exported function.
// Destroying is taken care of in CopyClipRectangleList.

func (ctx *Context) CopyClipRectangleList() ([]Rectangle, Status) {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	rl := C.cairo_copy_clip_rectangle_list(c_ctx)

	status := Status(rl.status)
	numRectangles := int(rl.num_rectangles)
	rectangles := make([]Rectangle, numRectangles, numRectangles)

	rectNative := (*[1 << 30]C.cairo_rectangle_t)(unsafe.Pointer(rl.rectangles))
	for i := 0; i < numRectangles; i++ {
		rectangles[i] = rectangleNew(&rectNative[i])
	}

	C.cairo_rectangle_list_destroy(rl)

	return rectangles, status
}

func (ctx *Context) GetReferenceCount() int {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_reference_count(c_ctx)
	return int(retC)
}

//cairo_status_t 	cairo_set_user_data ()
//void * 	cairo_get_user_data ()
