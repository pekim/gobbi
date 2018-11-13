// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_param_spec_override : unsupported parameter overridden : Blacklisted record : GParamSpec

// A predefined #GSignalAccumulator for signals that return a
// boolean values. The behavior that this accumulator gives is
// that a return of %TRUE stops the signal emission: no further
// callbacks will be invoked, while a return of %FALSE allows
// the emission to continue. The idea here is that a %TRUE return
// indicates that the callback handled the signal, and no further
// handling is needed.
/*

C function

g_signal_accumulator_true_handled
*/
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

// A more efficient version of g_type_class_peek() which works only for
// static types.
/*

C function

g_type_class_peek_static
*/
func TypeClassPeekStatic(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek_static(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

// If the interface type @g_type is currently in use, returns its
// default interface vtable.
/*

C function

g_type_default_interface_peek
*/
func TypeDefaultInterfacePeek(gType Type) uintptr {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_peek(c_g_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Increments the reference count for the interface type @g_type,
// and returns the default interface vtable for the type.
//
// If the type is not currently in use, then the default vtable
// for the type will be created and initalized by calling
// the base interface init and default vtable init functions for
// the type (the @base_init and @class_init members of #GTypeInfo).
// Calling g_type_default_interface_ref() is useful when you
// want to make sure that signals and properties for an interface
// have been installed.
/*

C function

g_type_default_interface_ref
*/
func TypeDefaultInterfaceRef(gType Type) uintptr {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_ref(c_g_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Decrements the reference count for the type corresponding to the
// interface default vtable @g_iface. If the type is dynamic, then
// when no one is using the interface and all references have
// been released, the finalize function for the interface's default
// vtable (the @class_finalize member of #GTypeInfo) will be called.
/*

C function

g_type_default_interface_unref
*/
func TypeDefaultInterfaceUnref(gIface uintptr) {
	c_g_iface := (C.gpointer)(gIface)

	C.g_type_default_interface_unref(c_g_iface)

	return
}

// Unsupported : g_type_remove_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func
