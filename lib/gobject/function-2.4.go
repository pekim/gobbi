// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecOverride_ is a wrapper around the C function g_param_spec_override.
func ParamSpecOverride_(name string, overridden *ParamSpec) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_overridden := (*C.GParamSpec)(C.NULL)
	if overridden != nil {
		c_overridden = (*C.GParamSpec)(overridden.ToC())
	}

	retC := C.g_param_spec_override(c_name, c_overridden)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SignalAccumulatorTrueHandled is a wrapper around the C function g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy uintptr) bool {
	c_ihint := (*C.GSignalInvocationHint)(C.NULL)
	if ihint != nil {
		c_ihint = (*C.GSignalInvocationHint)(ihint.ToC())
	}

	c_return_accu := (*C.GValue)(C.NULL)
	if returnAccu != nil {
		c_return_accu = (*C.GValue)(returnAccu.ToC())
	}

	c_handler_return := (*C.GValue)(C.NULL)
	if handlerReturn != nil {
		c_handler_return = (*C.GValue)(handlerReturn.ToC())
	}

	c_dummy := (C.gpointer)(dummy)

	retC := C.g_signal_accumulator_true_handled(c_ihint, c_return_accu, c_handler_return, c_dummy)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_add_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func

// TypeDefaultInterfacePeek is a wrapper around the C function g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType Type) uintptr {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_peek(c_g_type)
	retGo := (uintptr)(retC)

	return retGo
}

// TypeDefaultInterfaceRef is a wrapper around the C function g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType Type) uintptr {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_ref(c_g_type)
	retGo := (uintptr)(retC)

	return retGo
}

// TypeDefaultInterfaceUnref is a wrapper around the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface uintptr) {
	c_g_iface := (C.gpointer)(gIface)

	C.g_type_default_interface_unref(c_g_iface)

	return
}

// Unsupported : g_type_remove_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func
