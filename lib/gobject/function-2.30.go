// +build gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// CclosureMarshalGeneric is a wrapper around the C function g_cclosure_marshal_generic.
func CclosureMarshalGeneric(closure int, returnGvalue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}
