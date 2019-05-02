// This is a generated file - DO NOT EDIT
// +build gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	"C"
	"unsafe"
)

// Unbind is a wrapper around the C function g_binding_unbind.
func (recv *Binding) Unbind() {
	C.g_binding_unbind((*C.GBinding)(recv.native))

	return
}

// GetDefaultValue is a wrapper around the C function g_param_spec_get_default_value.
func (recv *ParamSpec) GetDefaultValue() *Value {
	retC := C.g_param_spec_get_default_value((*C.GParamSpec)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}
