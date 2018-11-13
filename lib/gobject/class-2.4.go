// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// This is a type of #GParamSpec type that simply redirects operations to
// another paramspec.  All operations other than getting or
// setting the value are redirected, including accessing the nick and
// blurb, validating a value, and so forth. See
// g_param_spec_get_redirect_target() for retrieving the overidden
// property. #GParamSpecOverride is used in implementing
// g_object_class_override_property(), and will not be directly useful
// unless you are implementing a new base type similar to GObject.
/*

C type

GParamSpecOverride
*/
type ParamSpecOverride struct {
	native *C.GParamSpecOverride
	// Private : parent_instance
	// Private : overridden
}

func ParamSpecOverrideNewFromC(u unsafe.Pointer) *ParamSpecOverride {
	c := (*C.GParamSpecOverride)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecOverride{native: c}

	return g
}

func (recv *ParamSpecOverride) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
