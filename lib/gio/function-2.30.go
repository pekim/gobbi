// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// DbusGvalueToGvariant is a wrapper around the C function g_dbus_gvalue_to_gvariant.
func DbusGvalueToGvariant(gvalue *gobject.Value, type_ *glib.VariantType) *glib.Variant {
	c_gvalue := (*C.GValue)(C.NULL)
	if gvalue != nil {
		c_gvalue = (*C.GValue)(gvalue.ToC())
	}

	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	retC := C.g_dbus_gvalue_to_gvariant(c_gvalue, c_type)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DbusGvariantToGvalue is a wrapper around the C function g_dbus_gvariant_to_gvalue.
func DbusGvariantToGvalue(value *glib.Variant) *gobject.Value {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	var c_out_gvalue C.GValue

	C.g_dbus_gvariant_to_gvalue(c_value, &c_out_gvalue)

	outGvalue := gobject.ValueNewFromC(unsafe.Pointer(&c_out_gvalue))

	return outGvalue
}

// IoModulesLoadAllInDirectoryWithScope is a wrapper around the C function g_io_modules_load_all_in_directory_with_scope.
func IoModulesLoadAllInDirectoryWithScope(dirname string, scope *IOModuleScope) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	c_scope := (*C.GIOModuleScope)(C.NULL)
	if scope != nil {
		c_scope = (*C.GIOModuleScope)(scope.ToC())
	}

	retC := C.g_io_modules_load_all_in_directory_with_scope(c_dirname, c_scope)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoModulesScanAllInDirectoryWithScope is a wrapper around the C function g_io_modules_scan_all_in_directory_with_scope.
func IoModulesScanAllInDirectoryWithScope(dirname string, scope *IOModuleScope) {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	c_scope := (*C.GIOModuleScope)(C.NULL)
	if scope != nil {
		c_scope = (*C.GIOModuleScope)(scope.ToC())
	}

	C.g_io_modules_scan_all_in_directory_with_scope(c_dirname, c_scope)

	return
}
