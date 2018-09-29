// This is a generated file - DO NOT EDIT
// +build gobject_2.46 gobject_2.54

package gobject

import glib "github.com/pekim/gobbi/lib/glib"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

// GetNameQuark is a wrapper around the C function g_param_spec_get_name_quark.
func (recv *ParamSpec) GetNameQuark() glib.Quark {
	retC := C.g_param_spec_get_name_quark((*C.GParamSpec)(recv.native))
	retGo := (glib.Quark)(retC)

	return retGo
}
