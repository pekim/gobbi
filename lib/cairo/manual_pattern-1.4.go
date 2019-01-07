// +build cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

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

// cairo_status_t 	cairo_pattern_get_color_stop_rgba ()
//cairo_status_t 	cairo_pattern_get_rgba ()
//cairo_status_t 	cairo_pattern_get_surface ()
//cairo_status_t 	cairo_pattern_get_linear_points ()

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
