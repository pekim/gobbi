// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "C"

// ListKeys is a wrapper around the C function g_settings_schema_list_keys.
func (recv *SettingsSchema) ListKeys() []string {
	retC := C.g_settings_schema_list_keys((*C.GSettingsSchema)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}
