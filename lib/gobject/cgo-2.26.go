// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Binding is a wrapper around the C record GBinding.
type Binding struct {
	native *C.GBinding
}

func BindingNewFromC(u unsafe.Pointer) *Binding {
	c := (*C.GBinding)(u)
	if c == nil {
		return nil
	}

	g := &Binding{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Binding) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Binding) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecVariant is a wrapper around the C record GParamSpecVariant.
type ParamSpecVariant struct {
	native *C.GParamSpecVariant
	// parent_instance : record
	// _type : record
	// default_value : record
	// Private : padding
}

func ParamSpecVariantNewFromC(u unsafe.Pointer) *ParamSpecVariant {
	c := (*C.GParamSpecVariant)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecVariant{native: c}

	return g
}

func (recv *ParamSpecVariant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_param_spec_variant : return type :
