// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

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
