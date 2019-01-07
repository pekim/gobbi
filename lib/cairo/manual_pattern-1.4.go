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
