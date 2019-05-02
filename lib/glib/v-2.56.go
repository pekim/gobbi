// This is a generated file - DO NOT EDIT
// +build glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_build_filename_valist : unsupported parameter args : no type generator for va_list (va_list*) for param args

// Unsupported : g_clear_handle_id : unsupported parameter clear_func : no type generator for ClearHandleFunc (GClearHandleFunc) for param clear_func

// Copy is a wrapper around the C function g_date_copy.
func (recv *Date) Copy() *Date {
	retC := C.g_date_copy((*C.GDate)(recv.native))
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromIso8601 is a wrapper around the C function g_date_time_new_from_iso8601.
func DateTimeNewFromIso8601(text string, defaultTz *TimeZone) *DateTime {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_default_tz := (*C.GTimeZone)(C.NULL)
	if defaultTz != nil {
		c_default_tz = (*C.GTimeZone)(defaultTz.ToC())
	}

	retC := C.g_date_time_new_from_iso8601(c_text, c_default_tz)
	var retGo (*DateTime)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DateTimeNewFromC(unsafe.Pointer(retC))
	}

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
