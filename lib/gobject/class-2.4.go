// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecOverride is a wrapper around the C record GParamSpecOverride.
type ParamSpecOverride struct {
	native *C.GParamSpecOverride
	// parent_instance : no type generator for ParamSpec, GParamSpec
	// overridden : no type generator for ParamSpec, GParamSpec*
}

func paramSpecOverrideNewFromC(c *C.GParamSpecOverride) *ParamSpecOverride {
	if c == nil {
		return nil
	}

	g := &ParamSpecOverride{native: c}

	return g
}
