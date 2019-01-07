// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"
import "unsafe"

func (p *Pattern) AddColorStopRGB(offset float64, red float64, green float64, blue float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_offset := (C.double)(offset)
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)

	C.cairo_pattern_add_color_stop_rgb(c_p, c_offset, c_red, c_green, c_blue)
}

func (p *Pattern) AddColorStopRGBA(offset float64, red float64, green float64, blue float64, alpha float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_offset := (C.double)(offset)
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)
	c_alpha := (C.double)(alpha)

	C.cairo_pattern_add_color_stop_rgba(c_p, c_offset, c_red, c_green, c_blue, c_alpha)
}

func CreateRGB(red float64, green float64, blue float64) *Pattern {
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)

	retC := C.cairo_pattern_create_rgb(c_red, c_green, c_blue)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

func CreateRGBA(red float64, green float64, blue float64, alpha float64) *Pattern {
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)
	c_alpha := (C.double)(alpha)

	retC := C.cairo_pattern_create_rgba(c_red, c_green, c_blue, c_alpha)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (s *Surface) CreatePattern() *Pattern {
	c_s := (*C.cairo_surface_t)(s.ToC())

	retC := C.cairo_pattern_create_for_surface(c_s)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

func CreateLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Pattern {
	c_x0 := (C.double)(x0)
	c_y0 := (C.double)(y0)
	c_x1 := (C.double)(x1)
	c_y1 := (C.double)(y1)

	retC := C.cairo_pattern_create_linear(c_x0, c_y0, c_x1, c_y1)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

func CreateRadial(
	cx0 float64, cy0 float64, radius0 float64,
	cx1 float64, cy1 float64, radius1 float64,
) *Pattern {
	c_cx0 := (C.double)(cx0)
	c_cy0 := (C.double)(cy0)
	c_radius0 := (C.double)(radius0)
	c_cx1 := (C.double)(cx1)
	c_cy1 := (C.double)(cy1)
	c_radius1 := (C.double)(radius1)

	retC := C.cairo_pattern_create_radial(c_cx0, c_cy0, c_radius0, c_cx1, c_cy1, c_radius1)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (p *Pattern) Reference() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_pattern_reference(c_p)
}

func (p *Pattern) Destroy() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_pattern_destroy(c_p)
}

func (p *Pattern) Status() Status {
	c_p := (*C.cairo_pattern_t)(p.ToC())

	retC := C.cairo_pattern_status(c_p)
	return Status(retC)
}

func (p *Pattern) SetExtend(extend Extend) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_extend := C.cairo_extend_t(extend)

	C.cairo_pattern_set_extend(c_p, c_extend)
}

func (p *Pattern) GetExtend() Extend {
	c_p := (*C.cairo_pattern_t)(p.ToC())

	retC := C.cairo_pattern_get_extend(c_p)
	return Extend(retC)
}

func (p *Pattern) SetFilter(filter Filter) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_filter := C.cairo_filter_t(filter)

	C.cairo_pattern_set_filter(c_p, c_filter)
}

func (p *Pattern) GetFilter() Filter {
	c_p := (*C.cairo_pattern_t)(p.ToC())

	retC := C.cairo_pattern_get_filter(c_p)
	return Filter(retC)
}

func (p *Pattern) SetMatrix(matrix Matrix) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_matrix := matrix.toC()

	C.cairo_pattern_set_matrix(c_p, c_matrix)
}

func (p *Pattern) GetMatrix() Matrix {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var c_matrix C.cairo_matrix_t

	C.cairo_pattern_get_matrix(c_p, &c_matrix)
	return matrixFromC(&c_matrix)
}
