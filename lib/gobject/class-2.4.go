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
