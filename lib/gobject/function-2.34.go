// +build gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeEnsure is a wrapper around the C function g_type_ensure.
func TypeEnsure(type_ int) {}
