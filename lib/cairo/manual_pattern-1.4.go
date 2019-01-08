// +build cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"
import "unsafe"

func (p *Pattern) GetReferenceCount() int {
	c_p := (*C.cairo_pattern_t)(p.ToC())

	retC := C.cairo_pattern_get_reference_count(c_p)
	return int(retC)
}

func (p *Pattern) GetColorStopCount() (int, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var count C.int

	retC := C.cairo_pattern_get_color_stop_count(c_p, &count)
	return int(count), Status(retC)
}

func (p *Pattern) GetColorStopRGBA(index int) (float64, float64, float64, float64, float64, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_index := (C.int)(index)
	var c_offset C.double
	var c_red C.double
	var c_green C.double
	var c_blue C.double
	var c_alpha C.double

	retC := C.cairo_pattern_get_color_stop_rgba(c_p, c_index, &c_offset, &c_red, &c_green, &c_blue, &c_alpha)

	return float64(c_offset), float64(c_red), float64(c_green), float64(c_blue), float64(c_alpha),
		Status(retC)
}

func (p *Pattern) GetRGBA() (float64, float64, float64, float64, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var c_red C.double
	var c_green C.double
	var c_blue C.double
	var c_alpha C.double

	retC := C.cairo_pattern_get_rgba(c_p, &c_red, &c_green, &c_blue, &c_alpha)

	return float64(c_red), float64(c_green), float64(c_blue), float64(c_alpha),
		Status(retC)
}

func (p *Pattern) GetSurface() (*Surface, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var c_surface *C.cairo_surface_t

	retC := C.cairo_pattern_get_surface(c_p, &c_surface)

	return SurfaceNewFromC(unsafe.Pointer(c_surface)), Status(retC)
}

//cairo_status_t 	cairo_pattern_get_linear_points ()

func (p *Pattern) GetLinearPoints() (float64, float64, float64, float64, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var c_x0 C.double
	var c_y0 C.double
	var c_x1 C.double
	var c_y1 C.double

	retC := C.cairo_pattern_get_linear_points(c_p, &c_x0, &c_y0, &c_x1, &c_y1)

	return float64(c_x0), float64(c_y0),
		float64(c_x1), float64(c_y1),
		Status(retC)
}

func (p *Pattern) GetRadialCircles() (float64, float64, float64, float64, float64, float64, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var c_x0 C.double
	var c_y0 C.double
	var c_r0 C.double
	var c_x1 C.double
	var c_y1 C.double
	var c_r1 C.double

	retC := C.cairo_pattern_get_radial_circles(c_p, &c_x0, &c_y0, &c_r0, &c_x1, &c_y1, &c_r0)

	return float64(c_x0), float64(c_y0), float64(c_r0),
		float64(c_x1), float64(c_y1), float64(c_r1),
		Status(retC)
}

//cairo_status_t 	cairo_pattern_set_user_data ()
//void * 	cairo_pattern_get_user_data ()
