// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10

package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Paths.html.

func (ctx *Context) CopyPath() *Path {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_copy_path(c_ctx)
	retGo := PathNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Path) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) CopyPathFlat() *Path {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_copy_path_flat(c_ctx)
	retGo := PathNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Path) {
		o.Destroy()
	})

	return retGo
}

func (path *Path) Destroy() {
	c_path := (*C.cairo_path_t)(path.ToC())
	C.cairo_path_destroy(c_path)
}

func (ctx *Context) AppendPath(path *Path) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_path := (*C.cairo_path_t)(path.ToC())
	C.cairo_append_path(c_ctx, c_path)
}

func (ctx *Context) GetCurrentPoint() (float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_x C.double
	var c_y C.double

	C.cairo_get_current_point(c_ctx, &c_x, &c_y)

	return float64(c_x), float64(c_y)
}

func (ctx *Context) NewPath() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_new_path(c_ctx)
}

func (ctx *Context) ClosePath() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_close_path(c_ctx)
}

func (ctx *Context) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_xc := (C.double)(xc)
	c_yc := (C.double)(yc)
	c_radius := (C.double)(radius)
	c_angle1 := (C.double)(angle1)
	c_angle2 := (C.double)(angle2)

	C.cairo_arc(c_ctx, c_xc, c_yc, c_radius, c_angle1, c_angle2)
}

func (ctx *Context) ArcNegative(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_xc := (C.double)(xc)
	c_yc := (C.double)(yc)
	c_radius := (C.double)(radius)
	c_angle1 := (C.double)(angle1)
	c_angle2 := (C.double)(angle2)

	C.cairo_arc_negative(c_ctx, c_xc, c_yc, c_radius, c_angle1, c_angle2)
}
