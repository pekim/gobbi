// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Dgettext is a wrapper around the C function g_dgettext.
func Dgettext(domain string, msgid string) {}

// Dngettext is a wrapper around the C function g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) {}

// Dpgettext2 is a wrapper around the C function g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) {}

// Unsupported : g_set_error_literal : unsupported parameter err : type Error, GError**
