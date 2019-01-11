// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Transformations.html.

func (ctx *Context) Translate(tx float64, ty float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_tx := C.double(tx)
	c_ty := C.double(ty)

	C.cairo_translate(c_ctx, c_tx, c_ty)
}

func (ctx *Context) Scale(sx float64, sy float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_sx := C.double(sx)
	c_sy := C.double(sy)

	C.cairo_scale(c_ctx, c_sx, c_sy)
}

func (ctx *Context) Rotate(angle float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_angle := C.double(angle)

	C.cairo_rotate(c_ctx, c_angle)
}

func (ctx *Context) Transform(matrix Matrix) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_matrix := matrix.toC()

	C.cairo_transform(c_ctx, c_matrix)
}

func (ctx *Context) SetMatrix(matrix Matrix) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_matrix := matrix.toC()

	C.cairo_set_matrix(c_ctx, c_matrix)
}

func (ctx *Context) GetMatrix() Matrix {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_matrix C.cairo_matrix_t

	C.cairo_get_matrix(c_ctx, &c_matrix)
	return matrixFromC(&c_matrix)
}

func (ctx *Context) IdentityMatrix() {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	C.cairo_identity_matrix(c_ctx)
}

func (ctx *Context) UserToDevice(x float64, y float64) (float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := C.double(x)
	c_y := C.double(y)

	C.cairo_user_to_device(c_ctx, &c_x, &c_y)

	return float64(c_x), float64(c_y)
}

func (ctx *Context) UserToDeviceDistance(x float64, y float64) (float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := C.double(x)
	c_y := C.double(y)

	C.cairo_user_to_device_distance(c_ctx, &c_x, &c_y)

	return float64(c_x), float64(c_y)
}

func (ctx *Context) DeviceToUser(x float64, y float64) (float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := C.double(x)
	c_y := C.double(y)

	C.cairo_device_to_user(c_ctx, &c_x, &c_y)

	return float64(c_x), float64(c_y)
}

func (ctx *Context) DeviceToUserDistance(x float64, y float64) (float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := C.double(x)
	c_y := C.double(y)

	C.cairo_device_to_user_distance(c_ctx, &c_x, &c_y)

	return float64(c_x), float64(c_y)
}
