// This is a generated file - DO NOT EDIT
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

func (recv *Binding) Object() *Object {}

func (recv *InitiallyUnowned) Object() *Object {}

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

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
func (recv *Object) RefSink() uintptr {
	retC := C.g_object_ref_sink((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// RefSink is a wrapper around the C function g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	retC := C.g_param_spec_ref_sink((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *ParamSpecBoolean) ParamSpec() *ParamSpec {}

func (recv *ParamSpecBoxed) ParamSpec() *ParamSpec {}

func (recv *ParamSpecChar) ParamSpec() *ParamSpec {}

func (recv *ParamSpecDouble) ParamSpec() *ParamSpec {}

func (recv *ParamSpecEnum) ParamSpec() *ParamSpec {}

func (recv *ParamSpecFlags) ParamSpec() *ParamSpec {}

func (recv *ParamSpecFloat) ParamSpec() *ParamSpec {}

// ParamSpecGType is a wrapper around the C record GParamSpecGType.
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : record
	// is_a_type : no type generator for GType, GType
}

func ParamSpecGTypeNewFromC(u unsafe.Pointer) *ParamSpecGType {
	c := (*C.GParamSpecGType)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{native: c}

	return g
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

func (recv *ParamSpecGType) ParamSpec() *ParamSpec {}

func (recv *ParamSpecInt) ParamSpec() *ParamSpec {}

func (recv *ParamSpecInt64) ParamSpec() *ParamSpec {}

func (recv *ParamSpecLong) ParamSpec() *ParamSpec {}

func (recv *ParamSpecObject) ParamSpec() *ParamSpec {}

func (recv *ParamSpecOverride) ParamSpec() *ParamSpec {}

func (recv *ParamSpecParam) ParamSpec() *ParamSpec {}

func (recv *ParamSpecPointer) ParamSpec() *ParamSpec {}

func (recv *ParamSpecString) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUChar) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUInt) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUInt64) ParamSpec() *ParamSpec {}

func (recv *ParamSpecULong) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUnichar) ParamSpec() *ParamSpec {}

func (recv *ParamSpecValueArray) ParamSpec() *ParamSpec {}

func (recv *ParamSpecVariant) ParamSpec() *ParamSpec {}

func (recv *TypeModule) Object() *Object {}
