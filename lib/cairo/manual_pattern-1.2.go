// +build cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

func (p *Pattern) GetType() PatternType {
	c_p := (*C.cairo_pattern_t)(p.ToC())

	retC := C.cairo_pattern_get_type(c_p)
	return PatternType(retC)
}
