// +build glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// StrIsAscii is a wrapper around the C function g_str_is_ascii.
func StrIsAscii(str string) {}

// Unsupported function: g_str_match_string : unsupported parameter accept_alternates : type gboolean, gboolean

// StrToAscii is a wrapper around the C function g_str_to_ascii.
func StrToAscii(str string, fromLocale string) {}

// Unsupported function: g_str_tokenize_and_fold : unsupported parameter ascii_alternates : no type

// Unsupported function: g_variant_parse_error_print_context : unsupported parameter error : type Error, GError*
