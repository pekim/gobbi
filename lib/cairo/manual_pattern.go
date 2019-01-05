package cairo

// #include <cairo/cairo.h>
import "C"

func (p *Pattern) GetReferenceCount() int {
	c_p := (*C.cairo_pattern_t)(p.ToC())

	retC := C.cairo_pattern_get_reference_count(c_p)
	return int(retC)
}

func (p *Pattern) Reference() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_pattern_reference(c_p)
}

func (p *Pattern) Destroy() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_pattern_destroy(c_p)
}
