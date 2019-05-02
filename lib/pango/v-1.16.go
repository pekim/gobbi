// Code generated - DO NOT EDIT.
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetBaseGravity is a wrapper around the C function pango_context_get_base_gravity.
func (recv *Context) GetBaseGravity() Gravity {
	retC := C.pango_context_get_base_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetGravity is a wrapper around the C function pango_context_get_gravity.
func (recv *Context) GetGravity() Gravity {
	retC := C.pango_context_get_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetGravityHint is a wrapper around the C function pango_context_get_gravity_hint.
func (recv *Context) GetGravityHint() GravityHint {
	retC := C.pango_context_get_gravity_hint((*C.PangoContext)(recv.native))
	retGo := (GravityHint)(retC)

	return retGo
}

// SetBaseGravity is a wrapper around the C function pango_context_set_base_gravity.
func (recv *Context) SetBaseGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_context_set_base_gravity((*C.PangoContext)(recv.native), c_gravity)

	return
}

// SetGravityHint is a wrapper around the C function pango_context_set_gravity_hint.
func (recv *Context) SetGravityHint(hint GravityHint) {
	c_hint := (C.PangoGravityHint)(hint)

	C.pango_context_set_gravity_hint((*C.PangoContext)(recv.native), c_hint)

	return
}

// GetLineReadonly is a wrapper around the C function pango_layout_get_line_readonly.
func (recv *Layout) GetLineReadonly(line int32) *LayoutLine {
	c_line := (C.int)(line)

	retC := C.pango_layout_get_line_readonly((*C.PangoLayout)(recv.native), c_line)
	var retGo (*LayoutLine)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutLineNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLinesReadonly is a wrapper around the C function pango_layout_get_lines_readonly.
func (recv *Layout) GetLinesReadonly() *glib.SList {
	retC := C.pango_layout_get_lines_readonly((*C.PangoLayout)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUnknownGlyphsCount is a wrapper around the C function pango_layout_get_unknown_glyphs_count.
func (recv *Layout) GetUnknownGlyphsCount() int32 {
	retC := C.pango_layout_get_unknown_glyphs_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsEllipsized is a wrapper around the C function pango_layout_is_ellipsized.
func (recv *Layout) IsEllipsized() bool {
	retC := C.pango_layout_is_ellipsized((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsWrapped is a wrapper around the C function pango_layout_is_wrapped.
func (recv *Layout) IsWrapped() bool {
	retC := C.pango_layout_is_wrapped((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

const ANALYSIS_FLAG_CENTERED_BASELINE int32 = C.PANGO_ANALYSIS_FLAG_CENTERED_BASELINE

type Gravity C.PangoGravity

const (
	PANGO_GRAVITY_SOUTH Gravity = 0
	PANGO_GRAVITY_EAST  Gravity = 1
	PANGO_GRAVITY_NORTH Gravity = 2
	PANGO_GRAVITY_WEST  Gravity = 3
	PANGO_GRAVITY_AUTO  Gravity = 4
)

// GravityGetForMatrix is a wrapper around the C function pango_gravity_get_for_matrix.
func GravityGetForMatrix(matrix *Matrix) Gravity {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	retC := C.pango_gravity_get_for_matrix(c_matrix)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityGetForScript is a wrapper around the C function pango_gravity_get_for_script.
func GravityGetForScript(script Script, baseGravity Gravity, hint GravityHint) Gravity {
	c_script := (C.PangoScript)(script)

	c_base_gravity := (C.PangoGravity)(baseGravity)

	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_gravity_get_for_script(c_script, c_base_gravity, c_hint)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityToRotation is a wrapper around the C function pango_gravity_to_rotation.
func GravityToRotation(gravity Gravity) float64 {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_gravity_to_rotation(c_gravity)
	retGo := (float64)(retC)

	return retGo
}

type GravityHint C.PangoGravityHint

const (
	PANGO_GRAVITY_HINT_NATURAL GravityHint = 0
	PANGO_GRAVITY_HINT_STRONG  GravityHint = 1
	PANGO_GRAVITY_HINT_LINE    GravityHint = 2
)

// AttrGravityHintNew is a wrapper around the C function pango_attr_gravity_hint_new.
func AttrGravityHintNew(hint GravityHint) *Attribute {
	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_attr_gravity_hint_new(c_hint)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrGravityNew is a wrapper around the C function pango_attr_gravity_new.
func AttrGravityNew(gravity Gravity) *Attribute {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_attr_gravity_new(c_gravity)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ExtentsToPixels is a wrapper around the C function pango_extents_to_pixels.
func ExtentsToPixels(inclusive *Rectangle, nearest *Rectangle) {
	c_inclusive := (*C.PangoRectangle)(C.NULL)
	if inclusive != nil {
		c_inclusive = (*C.PangoRectangle)(inclusive.ToC())
	}

	c_nearest := (*C.PangoRectangle)(C.NULL)
	if nearest != nil {
		c_nearest = (*C.PangoRectangle)(nearest.ToC())
	}

	C.pango_extents_to_pixels(c_inclusive, c_nearest)

	return
}

// ParseEnum is a wrapper around the C function pango_parse_enum.
func ParseEnum(type_ gobject.Type, str string, warn bool) (bool, int32, string) {
	c_type := (C.GType)(type_)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_value C.int

	c_warn :=
		boolToGboolean(warn)

	var c_possible_values *C.char

	retC := C.pango_parse_enum(c_type, c_str, &c_value, c_warn, &c_possible_values)
	retGo := retC == C.TRUE

	value := (int32)(c_value)

	possibleValues := C.GoString(c_possible_values)
	defer C.free(unsafe.Pointer(c_possible_values))

	return retGo, value, possibleValues
}

// UnitsFromDouble is a wrapper around the C function pango_units_from_double.
func UnitsFromDouble(d float64) int32 {
	c_d := (C.double)(d)

	retC := C.pango_units_from_double(c_d)
	retGo := (int32)(retC)

	return retGo
}

// UnitsToDouble is a wrapper around the C function pango_units_to_double.
func UnitsToDouble(i int32) float64 {
	c_i := (C.int)(i)

	retC := C.pango_units_to_double(c_i)
	retGo := (float64)(retC)

	return retGo
}

// Version is a wrapper around the C function pango_version.
func Version() int32 {
	retC := C.pango_version()
	retGo := (int32)(retC)

	return retGo
}

// VersionCheck is a wrapper around the C function pango_version_check.
func VersionCheck(requiredMajor int32, requiredMinor int32, requiredMicro int32) string {
	c_required_major := (C.int)(requiredMajor)

	c_required_minor := (C.int)(requiredMinor)

	c_required_micro := (C.int)(requiredMicro)

	retC := C.pango_version_check(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// VersionString is a wrapper around the C function pango_version_string.
func VersionString() string {
	retC := C.pango_version_string()
	retGo := C.GoString(retC)

	return retGo
}

// ToString is a wrapper around the C function pango_color_to_string.
func (recv *Color) ToString() string {
	retC := C.pango_color_to_string((*C.PangoColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetGravity is a wrapper around the C function pango_font_description_get_gravity.
func (recv *FontDescription) GetGravity() Gravity {
	retC := C.pango_font_description_get_gravity((*C.PangoFontDescription)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// SetGravity is a wrapper around the C function pango_font_description_set_gravity.
func (recv *FontDescription) SetGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_font_description_set_gravity((*C.PangoFontDescription)(recv.native), c_gravity)

	return
}

// LanguageGetDefault is a wrapper around the C function pango_language_get_default.
func LanguageGetDefault() *Language {
	retC := C.pango_language_get_default()
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TransformDistance is a wrapper around the C function pango_matrix_transform_distance.
func (recv *Matrix) TransformDistance(dx float64, dy float64) {
	c_dx := (C.double)(dx)

	c_dy := (C.double)(dy)

	C.pango_matrix_transform_distance((*C.PangoMatrix)(recv.native), &c_dx, &c_dy)

	return
}

// TransformPixelRectangle is a wrapper around the C function pango_matrix_transform_pixel_rectangle.
func (recv *Matrix) TransformPixelRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_pixel_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

// TransformPoint is a wrapper around the C function pango_matrix_transform_point.
func (recv *Matrix) TransformPoint(x float64, y float64) {
	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_matrix_transform_point((*C.PangoMatrix)(recv.native), &c_x, &c_y)

	return
}

// TransformRectangle is a wrapper around the C function pango_matrix_transform_rectangle.
func (recv *Matrix) TransformRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}
