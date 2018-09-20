// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

type BindingFlags C.GBindingFlags

const (
	BINDING_DEFAULT        BindingFlags = 0
	BINDING_BIDIRECTIONAL  BindingFlags = 1
	BINDING_SYNC_CREATE    BindingFlags = 2
	BINDING_INVERT_BOOLEAN BindingFlags = 4
)
