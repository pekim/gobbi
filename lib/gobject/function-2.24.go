// +build gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeAddClassPrivate is a wrapper around the C function g_type_add_class_private.
func TypeAddClassPrivate(classType int, privateSize int) {}
