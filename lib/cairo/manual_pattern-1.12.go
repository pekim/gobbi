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

func (p *Pattern) MeshSetCornerColorRGB(cornerNum uint, red float64, green float64, blue float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_cornerNum := C.uint(cornerNum)
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)

	C.cairo_mesh_pattern_set_corner_color_rgb(c_p, c_cornerNum, c_red, c_green, c_blue)
}

func (p *Pattern) MeshSetCornerColorRGBA(cornerNum uint, red float64, green float64, blue float64, alpha float64) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_cornerNum := C.uint(cornerNum)
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)
	c_alpha := (C.double)(alpha)

	C.cairo_mesh_pattern_set_corner_color_rgba(c_p, c_cornerNum, c_red, c_green, c_blue, c_alpha)
}

func (p *Pattern) MeshGetPatchCount() (uint, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	var count C.uint

	retC := C.cairo_mesh_pattern_get_patch_count(c_p, &count)
	return uint(count), Status(retC)
}

func (p *Pattern) MeshGetPath(patchNum uint) *Path {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_patchNum := C.uint(patchNum)

	retC := C.cairo_mesh_pattern_get_path(c_p, c_patchNum)
	retGo := PathNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (p *Pattern) MeshGetControlPoint(patchNum uint, pointNum uint) (float64, float64, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_patchNum := C.uint(patchNum)
	c_pointNum := C.uint(pointNum)
	var c_x C.double
	var c_y C.double

	retC := C.cairo_mesh_pattern_get_control_point(c_p, c_patchNum, c_pointNum, &c_x, &c_y)

	return float64(c_x), float64(c_y), Status(retC)
}

func (p *Pattern) MeshGetCornerColorRGBA(patchNum uint, pointNum uint) (float64, float64, float64, float64, Status) {
	c_p := (*C.cairo_pattern_t)(p.ToC())
	c_patchNum := C.uint(patchNum)
	c_pointNum := C.uint(pointNum)
	var c_red C.double
	var c_green C.double
	var c_blue C.double
	var c_alpha C.double

	retC := C.cairo_mesh_pattern_get_corner_color_rgba(c_p, c_patchNum, c_pointNum,
		&c_red, &c_green, &c_blue, &c_alpha)

	return float64(c_red), float64(c_green), float64(c_blue), float64(c_alpha), Status(retC)
}
