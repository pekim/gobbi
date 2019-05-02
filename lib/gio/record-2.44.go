// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "C"

// Equals compares this ListModelInterface with another ListModelInterface, and returns true if they represent the same GObject.
func (recv *ListModelInterface) Equals(other *ListModelInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this OutputMessage with another OutputMessage, and returns true if they represent the same GObject.
func (recv *OutputMessage) Equals(other *OutputMessage) bool {
	return other.ToC() == recv.ToC()
}

// ListChildren is a wrapper around the C function g_settings_schema_list_children.
func (recv *SettingsSchema) ListChildren() []string {
	retC := C.g_settings_schema_list_children((*C.GSettingsSchema)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetName is a wrapper around the C function g_settings_schema_key_get_name.
func (recv *SettingsSchemaKey) GetName() string {
	retC := C.g_settings_schema_key_get_name((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
