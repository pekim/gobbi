// This is a generated file - DO NOT EDIT
// +build glib_2.56

package glib

import (
	"C"
	"unsafe"
)

// Copy is a wrapper around the C function g_date_copy.
func (recv *Date) Copy() *Date {
	retC := C.g_date_copy((*C.GDate)(recv.native))
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLocaleForKey is a wrapper around the C function g_key_file_get_locale_for_key.
func (recv *KeyFile) GetLocaleForKey(groupName string, key string, locale string) string {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	retC := C.g_key_file_get_locale_for_key((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
