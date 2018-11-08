// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// GetKey is a wrapper around the C function g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) *SettingsSchemaKey {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_schema_get_key((*C.GSettingsSchema)(recv.native), c_name)
	retGo := SettingsSchemaKeyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasKey is a wrapper around the C function g_settings_schema_has_key.
func (recv *SettingsSchema) HasKey(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_schema_has_key((*C.GSettingsSchema)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_settings_schema_key_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_value_type : return type : Blacklisted record : GVariantType

// Unsupported : g_settings_schema_key_range_check : unsupported parameter value : Blacklisted record : GVariant

// Ref is a wrapper around the C function g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	retC := C.g_settings_schema_key_ref((*C.GSettingsSchemaKey)(recv.native))
	retGo := SettingsSchemaKeyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() {
	C.g_settings_schema_key_unref((*C.GSettingsSchemaKey)(recv.native))

	return
}

// Unsupported : g_settings_schema_source_list_schemas : unsupported parameter non_relocatable : no type generator for utf8 (gchar**) for array param non_relocatable
