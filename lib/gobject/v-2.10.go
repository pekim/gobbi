// Code generated - DO NOT EDIT.
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ForceFloating is a wrapper around the C function g_object_force_floating.
func (recv *Object) ForceFloating() {
	C.g_object_force_floating((*C.GObject)(recv.native))

	return
}

// IsFloating is a wrapper around the C function g_object_is_floating.
func (recv *Object) IsFloating() bool {
	retC := C.g_object_is_floating((C.gpointer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RefSink is a wrapper around the C function g_object_ref_sink.
func (recv *Object) RefSink() Object {
	retC := C.g_object_ref_sink((C.gpointer)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefSink is a wrapper around the C function g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	retC := C.g_param_spec_ref_sink((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecGType is a wrapper around the C record GParamSpecGType.
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : record
	IsAType Type
}

func ParamSpecGTypeNewFromC(u unsafe.Pointer) *ParamSpecGType {
	c := (*C.GParamSpecGType)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{
		IsAType: (Type)(c.is_a_type),
		native:  c,
	}

	return g
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {
	recv.native.is_a_type =
		(C.GType)(recv.IsAType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecGType with another ParamSpecGType, and returns true if they represent the same GObject.
func (recv *ParamSpecGType) Equals(other *ParamSpecGType) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecGType) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecGtype is a wrapper around the C function g_param_spec_gtype.
func ParamSpecGtype(name string, nick string, blurb string, isAType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_is_a_type := (C.GType)(isAType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_gtype(c_name, c_nick, c_blurb, c_is_a_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}
