// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecGtype is a wrapper around the C function g_param_spec_gtype.
func ParamSpecGtype(name int, nick int, blurb int, isAType int, flags int) {}
