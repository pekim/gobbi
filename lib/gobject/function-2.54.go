// +build gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// EnumToString is a wrapper around the C function g_enum_to_string.
func EnumToString(gEnumType int, value int) {}

// FlagsToString is a wrapper around the C function g_flags_to_string.
func FlagsToString(flagsType int, value int) {}
