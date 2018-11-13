// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Create a new gravity hint attribute.
/*

C function : pango_attr_gravity_hint_new
*/
func AttrGravityHintNew(hint GravityHint) *Attribute {
	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_attr_gravity_hint_new(c_hint)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new gravity attribute.
/*

C function : pango_attr_gravity_new
*/
func AttrGravityNew(gravity Gravity) *Attribute {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_attr_gravity_new(c_gravity)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts extents from Pango units to device units, dividing by the
// %PANGO_SCALE factor and performing rounding.
//
// The @inclusive rectangle is converted by flooring the x/y coordinates and extending
// width/height, such that the final rectangle completely includes the original
// rectangle.
//
// The @nearest rectangle is converted by rounding the coordinates
// of the rectangle to the nearest device unit (pixel).
//
// The rule to which argument to use is: if you want the resulting device-space
// rectangle to completely contain the original rectangle, pass it in as @inclusive.
// If you want two touching-but-not-overlapping rectangles stay
// touching-but-not-overlapping after rounding to device units, pass them in
// as @nearest.
/*

C function : pango_extents_to_pixels
*/
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

// Finds the gravity that best matches the rotation component
// in a #PangoMatrix.
/*

C function : pango_gravity_get_for_matrix
*/
func GravityGetForMatrix(matrix *Matrix) Gravity {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	retC := C.pango_gravity_get_for_matrix(c_matrix)
	retGo := (Gravity)(retC)

	return retGo
}

// Based on the script, base gravity, and hint, returns actual gravity
// to use in laying out a single #PangoItem.
//
// If @base_gravity is %PANGO_GRAVITY_AUTO, it is first replaced with the
// preferred gravity of @script.  To get the preferred gravity of a script,
// pass %PANGO_GRAVITY_AUTO and %PANGO_GRAVITY_HINT_STRONG in.
/*

C function : pango_gravity_get_for_script
*/
func GravityGetForScript(script Script, baseGravity Gravity, hint GravityHint) Gravity {
	c_script := (C.PangoScript)(script)

	c_base_gravity := (C.PangoGravity)(baseGravity)

	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_gravity_get_for_script(c_script, c_base_gravity, c_hint)
	retGo := (Gravity)(retC)

	return retGo
}

// Converts a #PangoGravity value to its natural rotation in radians.
// @gravity should not be %PANGO_GRAVITY_AUTO.
//
// Note that pango_matrix_rotate() takes angle in degrees, not radians.
// So, to call pango_matrix_rotate() with the output of this function
// you should multiply it by (180. / G_PI).
/*

C function : pango_gravity_to_rotation
*/
func GravityToRotation(gravity Gravity) float64 {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_gravity_to_rotation(c_gravity)
	retGo := (float64)(retC)

	return retGo
}

// Returns the #PangoLanguage for the current locale of the process.
// Note that this can change over the life of an application.
//
// On Unix systems, this is the return value is derived from
// <literal>setlocale(LC_CTYPE, NULL)</literal>, and the user can
// affect this through the environment variables LC_ALL, LC_CTYPE or
// LANG (checked in that order). The locale string typically is in
// the form lang_COUNTRY, where lang is an ISO-639 language code, and
// COUNTRY is an ISO-3166 country code. For instance, sv_FI for
// Swedish as written in Finland or pt_BR for Portuguese as written in
// Brazil.
//
// On Windows, the C library does not use any such environment
// variables, and setting them won't affect the behavior of functions
// like ctime(). The user sets the locale through the Regional Options
// in the Control Panel. The C library (in the setlocale() function)
// does not use country and language codes, but country and language
// names spelled out in English.
// However, this function does check the above environment
// variables, and does return a Unix-style locale string based on
// either said environment variables or the thread's current locale.
//
// Your application should call <literal>setlocale(LC_ALL, "");</literal>
// for the user settings to take effect.  Gtk+ does this in its initialization
// functions automatically (by calling gtk_set_locale()).
// See <literal>man setlocale</literal> for more details.
/*

C function : pango_language_get_default
*/
func LanguageGetDefault() *Language {
	retC := C.pango_language_get_default()
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parses an enum type and stores the result in @value.
//
// If @str does not match the nick name of any of the possible values for the
// enum and is not an integer, %FALSE is returned, a warning is issued
// if @warn is %TRUE, and a
// string representing the list of possible values is stored in
// @possible_values.  The list is slash-separated, eg.
// "none/start/middle/end".  If failed and @possible_values is not %NULL,
// returned string should be freed using g_free().
/*

C function : pango_parse_enum
*/
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

// Converts a floating-point number to Pango units: multiplies
// it by %PANGO_SCALE and rounds to nearest integer.
/*

C function : pango_units_from_double
*/
func UnitsFromDouble(d float64) int32 {
	c_d := (C.double)(d)

	retC := C.pango_units_from_double(c_d)
	retGo := (int32)(retC)

	return retGo
}

// Converts a number in Pango units to floating-point: divides
// it by %PANGO_SCALE.
/*

C function : pango_units_to_double
*/
func UnitsToDouble(i int32) float64 {
	c_i := (C.int)(i)

	retC := C.pango_units_to_double(c_i)
	retGo := (float64)(retC)

	return retGo
}

// This is similar to the macro %PANGO_VERSION except that
// it returns the encoded version of Pango available at run-time,
// as opposed to the version available at compile-time.
//
// A version number can be encoded into an integer using
// PANGO_VERSION_ENCODE().
/*

C function : pango_version
*/
func Version() int32 {
	retC := C.pango_version()
	retGo := (int32)(retC)

	return retGo
}

// Checks that the Pango library in use is compatible with the
// given version. Generally you would pass in the constants
// %PANGO_VERSION_MAJOR, %PANGO_VERSION_MINOR, %PANGO_VERSION_MICRO
// as the three arguments to this function; that produces
// a check that the library in use at run-time is compatible with
// the version of Pango the application or module was compiled against.
//
// Compatibility is defined by two things: first the version
// of the running library is newer than the version
// @required_major.required_minor.@required_micro. Second
// the running library must be binary compatible with the
// version @required_major.required_minor.@required_micro
// (same major version.)
//
// For compile-time version checking use PANGO_VERSION_CHECK().
/*

C function : pango_version_check
*/
func VersionCheck(requiredMajor int32, requiredMinor int32, requiredMicro int32) string {
	c_required_major := (C.int)(requiredMajor)

	c_required_minor := (C.int)(requiredMinor)

	c_required_micro := (C.int)(requiredMicro)

	retC := C.pango_version_check(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// This is similar to the macro %PANGO_VERSION_STRING except that
// it returns the version of Pango available at run-time, as opposed to
// the version available at compile-time.
/*

C function : pango_version_string
*/
func VersionString() string {
	retC := C.pango_version_string()
	retGo := C.GoString(retC)

	return retGo
}
