// +build glib_2.2 glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// AsciiStrtoull is a wrapper around the C function g_ascii_strtoull.
func AsciiStrtoull(nptr string, endptr string, base uint32) {}

// Unsupported : g_fprintf : unsupported parameter file : type gpointer, FILE*

// GetApplicationName is a wrapper around the C function g_get_application_name.
func GetApplicationName() {}

// Unsupported : g_printf : unsupported parameter ... : varargs

// SetApplicationName is a wrapper around the C function g_set_application_name.
func SetApplicationName(applicationName string) {}

// Unsupported : g_sprintf : unsupported parameter ... : varargs

// StrHasPrefix is a wrapper around the C function g_str_has_prefix.
func StrHasPrefix(str string, prefix string) {}

// StrHasSuffix is a wrapper around the C function g_str_has_suffix.
func StrHasSuffix(str string, suffix string) {}

// Utf8Strreverse is a wrapper around the C function g_utf8_strreverse.
func Utf8Strreverse(str string, len int64) {}

// Unsupported : g_vfprintf : unsupported parameter file : type gpointer, FILE*

// Unsupported : g_vprintf : unsupported parameter args : type va_list, va_list

// Unsupported : g_vsprintf : unsupported parameter args : type va_list, va_list
