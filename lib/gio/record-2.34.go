// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	"unsafe"
)

// GetDescription is a wrapper around the C function g_settings_schema_key_get_description.
func (recv *SettingsSchemaKey) GetDescription() string {
	retC := C.g_settings_schema_key_get_description((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSummary is a wrapper around the C function g_settings_schema_key_get_summary.
func (recv *SettingsSchemaKey) GetSummary() string {
	retC := C.g_settings_schema_key_get_summary((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GuessSymbolicIcon is a wrapper around the C function g_unix_mount_point_guess_symbolic_icon.
func (recv *UnixMountPoint) GuessSymbolicIcon() *Icon {
	retC := C.g_unix_mount_point_guess_symbolic_icon((*C.GUnixMountPoint)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}
