// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Flags to be passed to g_object_bind_property() or
// g_object_bind_property_full().
//
// This enumeration can be extended at later date.
type BindingFlags C.GBindingFlags

const (
	// The default binding; if the source property
	// changes, the target property is updated with its value.
	BINDING_DEFAULT BindingFlags = 0

	// Bidirectional binding; if either the
	// property of the source or the property of the target changes,
	// the other is updated.
	BINDING_BIDIRECTIONAL BindingFlags = 1

	// Synchronize the values of the source and
	// target properties when creating the binding; the direction of
	// the synchronization is always from the source to the target.
	BINDING_SYNC_CREATE BindingFlags = 2

	// If the two properties being bound are
	// booleans, setting one to %TRUE will result in the other being
	// set to %FALSE and vice versa. This flag will only work for
	// boolean properties, and cannot be used when passing custom
	// transformation functions to g_object_bind_property_full().
	BINDING_INVERT_BOOLEAN BindingFlags = 4
)
