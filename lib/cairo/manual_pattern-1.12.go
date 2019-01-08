// +build cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"
import "unsafe"

func PatternCreateMesh() *Pattern {
	retC := C.cairo_pattern_create_mesh()
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (p *Pattern) MeshBeginPatch() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_mesh_pattern_begin_patch(c_p)
}

func (p *Pattern) MeshEndPatch() {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	C.cairo_mesh_pattern_end_patch(c_p)
}

func (p *Pattern) MeshMoveTo(x float64, y float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_mesh_pattern_move_to(c_p, c_x, c_y)
}

func (p *Pattern) MeshLineTo(x float64, y float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_mesh_pattern_line_to(c_p, c_x, c_y)
}

func (p *Pattern) MeshCurveTo(
	x1 float64, y1 float64,
	x2 float64, y2 float64,
	x3 float64, y3 float64,
) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_x1 := (C.double)(x1)
	c_y1 := (C.double)(y1)
	c_x2 := (C.double)(x2)
	c_y2 := (C.double)(y2)
	c_x3 := (C.double)(x3)
	c_y3 := (C.double)(y3)

	C.cairo_mesh_pattern_curve_to(c_p, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3)
}

func (p *Pattern) MeshSetControlPoint(pointNum uint, x float64, y float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_pointNum := C.uint(pointNum)
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_mesh_pattern_set_control_point(c_p, c_pointNum, c_x, c_y)
}

//void 	cairo_mesh_pattern_set_corner_color_rgb ()
//void 	cairo_mesh_pattern_set_corner_color_rgba ()
//cairo_status_t 	cairo_mesh_pattern_get_patch_count ()
//cairo_path_t * 	cairo_mesh_pattern_get_path ()
//cairo_status_t 	cairo_mesh_pattern_get_control_point ()
//cairo_status_t 	cairo_mesh_pattern_get_corner_color_rgba ()
