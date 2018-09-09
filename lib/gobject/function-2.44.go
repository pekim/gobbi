// +build gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeGetInstanceCount is a wrapper around the C function g_type_get_instance_count.
func TypeGetInstanceCount(type_ int) {}
