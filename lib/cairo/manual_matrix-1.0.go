// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-matrix-t.html.

func (m *Matrix) Init(xx, yx, xy, yy, x0, y0 float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_xx := C.double(xx)
	c_yx := C.double(yx)
	c_xy := C.double(xy)
	c_yy := C.double(yy)
	c_x0 := C.double(x0)
	c_y0 := C.double(y0)

	C.cairo_matrix_init(c_m, c_xx, c_yx, c_xy, c_yy, c_x0, c_y0)
}

func (m *Matrix) InitIdentity() {
	c_m := (*C.cairo_matrix_t)(m.toC())
	C.cairo_matrix_init_identity(c_m)
}

func (m *Matrix) InitTranslate(tx, ty float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_tx := C.double(tx)
	c_ty := C.double(ty)

	C.cairo_matrix_init_translate(c_m, c_tx, c_ty)
}

func (m *Matrix) InitScale(sx, sy float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_sx := C.double(sx)
	c_sy := C.double(sy)

	C.cairo_matrix_init_scale(c_m, c_sx, c_sy)
}

func (m *Matrix) InitRotate(radians float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_radians := C.double(radians)

	C.cairo_matrix_init_rotate(c_m, c_radians)
}

func (m *Matrix) Translate(tx, ty float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_tx := C.double(tx)
	c_ty := C.double(ty)

	C.cairo_matrix_translate(c_m, c_tx, c_ty)
}

func (m *Matrix) Scale(sx, sy float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_sx := C.double(sx)
	c_sy := C.double(sy)

	C.cairo_matrix_scale(c_m, c_sx, c_sy)
}

func (m *Matrix) Rotate(radians float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_radians := C.double(radians)

	C.cairo_matrix_rotate(c_m, c_radians)
}

func (m *Matrix) Invert() Status {
	c_m := (*C.cairo_matrix_t)(m.toC())

	retC := C.cairo_matrix_invert(c_m)
	return Status(retC)
}

// Multiply multiplies the affine transformations in a and b together
// and stores the result in this matrix (m).
func (m *Matrix) Multiply(a, b *Matrix) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_a := (*C.cairo_matrix_t)(a.toC())
	c_b := (*C.cairo_matrix_t)(b.toC())

	C.cairo_matrix_multiply(c_m, c_a, c_b)
}

func (m *Matrix) TransformDistance(dx, dy float64) (float64, float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_dx := C.double(dx)
	c_dy := C.double(dy)

	C.cairo_matrix_transform_distance(c_m, &c_dx, &c_dy)
	return float64(dx), float64(dy)
}

func (m *Matrix) TransformPoint(x, y float64) (float64, float64) {
	c_m := (*C.cairo_matrix_t)(m.toC())
	c_x := C.double(x)
	c_y := C.double(y)

	C.cairo_matrix_transform_point(c_m, &c_x, &c_y)
	return float64(x), float64(y)
}
