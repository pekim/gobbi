// This is a generated file - DO NOT EDIT
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecGType is a wrapper around the C record GParamSpecGType.
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : no type generator for ParamSpec, GParamSpec
	// is_a_type : no type generator for GType, GType
}

func paramSpecGTypeNewFromC(c *C.GParamSpecGType) *ParamSpecGType {
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{native: c}

	return g
}

func (recv *ParamSpecGType) toC() *C.GParamSpecGType {

	return recv.native
}
