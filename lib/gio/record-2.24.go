// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// ConverterIface is a wrapper around the C record GConverterIface.
type ConverterIface struct {
	native *C.GConverterIface
	// g_iface : record
	// no type for convert
	// no type for reset
}

func ConverterIfaceNewFromC(u unsafe.Pointer) *ConverterIface {
	c := (*C.GConverterIface)(u)
	if c == nil {
		return nil
	}

	g := &ConverterIface{native: c}

	return g
}

func (recv *ConverterIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_io_extension_get_type : no return generator

// Unsupported : g_io_extension_point_get_required_type : no return generator

// Unsupported : g_io_extension_point_set_required_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_io_scheduler_job_send_to_mainloop : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_io_scheduler_job_send_to_mainloop_async : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_resource_enumerate_children : no return type

// Unsupported : g_resource_get_info : unsupported parameter size : no type generator for gsize, gsize*

// Unsupported : g_settings_schema_list_children : no return type

// Unsupported : g_settings_schema_list_keys : no return type

// Unsupported : g_settings_schema_key_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_value_type : return type : Blacklisted record : GVariantType

// Unsupported : g_settings_schema_key_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_schema_source_list_schemas : unsupported parameter non_relocatable : no param type

// Unsupported : g_unix_mount_point_guess_icon : no return generator

// Unsupported : g_unix_mount_point_guess_symbolic_icon : no return generator
