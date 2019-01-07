package cairo

// #include <cairo/cairo.h>
import "C"

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

func (p *Pattern) Reference() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_pattern_reference(c_p)
}

func (p *Pattern) Destroy() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_pattern_destroy(c_p)
}
