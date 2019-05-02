// Code generated - DO NOT EDIT.
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ObjectInterfaceFindProperty is a wrapper around the C function g_object_interface_find_property.
func ObjectInterfaceFindProperty(gIface *TypeInterface, propertyName string) *ParamSpec {
	c_g_iface := (C.gpointer)(C.NULL)
	if gIface != nil {
		c_g_iface = (C.gpointer)(gIface.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_object_interface_find_property(c_g_iface, c_property_name)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInterfaceInstallProperty is a wrapper around the C function g_object_interface_install_property.
func ObjectInterfaceInstallProperty(gIface *TypeInterface, pspec *ParamSpec) {
	c_g_iface := (C.gpointer)(C.NULL)
	if gIface != nil {
		c_g_iface = (C.gpointer)(gIface.ToC())
	}

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_object_interface_install_property(c_g_iface, c_pspec)

	return
}

// g_object_interface_list_properties : array return type :
// GetRedirectTarget is a wrapper around the C function g_param_spec_get_redirect_target.
func (recv *ParamSpec) GetRedirectTarget() *ParamSpec {
	retC := C.g_param_spec_get_redirect_target((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecOverride is a wrapper around the C record GParamSpecOverride.
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

// Equals compares this ParamSpecOverride with another ParamSpecOverride, and returns true if they represent the same GObject.
func (recv *ParamSpecOverride) Equals(other *ParamSpecOverride) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecOverride) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

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

// Unsupported : g_signal_accumulator_true_handled : unsupported parameter dummy : no type generator for gpointer (gpointer) for param dummy

// Unsupported : g_type_add_interface_check : unsupported parameter check_data : no type generator for gpointer (gpointer) for param check_data

// TypeDefaultInterfacePeek is a wrapper around the C function g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType Type) TypeInterface {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_peek(c_g_type)
	retGo := *TypeInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeDefaultInterfaceRef is a wrapper around the C function g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType Type) TypeInterface {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_ref(c_g_type)
	retGo := *TypeInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeDefaultInterfaceUnref is a wrapper around the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface *TypeInterface) {
	c_g_iface := (C.gpointer)(C.NULL)
	if gIface != nil {
		c_g_iface = (C.gpointer)(gIface.ToC())
	}

	C.g_type_default_interface_unref(c_g_iface)

	return
}

// Unsupported : g_type_remove_interface_check : unsupported parameter check_data : no type generator for gpointer (gpointer) for param check_data

// OverrideProperty is a wrapper around the C function g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	c_property_id := (C.guint)(propertyId)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_object_class_override_property((*C.GObjectClass)(recv.native), c_property_id, c_name)

	return
}

// TypeClassPeekStatic is a wrapper around the C function g_type_class_peek_static.
func TypeClassPeekStatic(type_ Type) TypeClass {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek_static(c_type)
	retGo := *TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddPrivate is a wrapper around the C function g_type_class_add_private.
func (recv *TypeClass) AddPrivate(privateSize uint64) {
	c_private_size := (C.gsize)(privateSize)

	C.g_type_class_add_private((C.gpointer)(recv.native), c_private_size)

	return
}

// Unsupported : g_value_take_boxed : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Unsupported : g_value_take_object : unsupported parameter v_object : no type generator for gpointer (gpointer) for param v_object

// TakeParam is a wrapper around the C function g_value_take_param.
func (recv *Value) TakeParam(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_take_param((*C.GValue)(recv.native), c_param)

	return
}

// TakeString is a wrapper around the C function g_value_take_string.
func (recv *Value) TakeString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_take_string((*C.GValue)(recv.native), c_v_string)

	return
}
