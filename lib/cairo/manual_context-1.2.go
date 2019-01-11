// +build cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-t.html.

func (ctx *Context) PushGroup() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_push_group(c_ctx)
}

func (ctx *Context) PushGroupWithContent(content Content) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_content := (C.cairo_content_t)(content)
	C.cairo_push_group_with_content(c_ctx, c_content)
}

func (ctx *Context) PopGroup() *Pattern {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_pop_group(c_ctx)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Pattern) {
		o.destroy()
	})

	return retGo
}

func (ctx *Context) PopGroupToSource() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_pop_group_to_source(c_ctx)
}

func (ctx *Context) GetGroupTarget() *Surface {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_group_target(c_ctx)
	retGo := SurfaceNewFromC(unsafe.Pointer(retC))

	retGo.reference()
	runtime.SetFinalizer(retGo, func(o *Surface) {
		o.destroy()
	})

	return retGo
}
