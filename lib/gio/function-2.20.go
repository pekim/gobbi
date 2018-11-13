// This is a generated file - DO NOT EDIT
// +build gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// Removes all changes to the type associations done by
// g_app_info_set_as_default_for_type(),
// g_app_info_set_as_default_for_extension(),
// g_app_info_add_supports_type() or
// g_app_info_remove_supports_type().
/*

C function

g_app_info_reset_type_associations
*/
func AppInfoResetTypeAssociations(contentType string) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	C.g_app_info_reset_type_associations(c_content_type)

	return
}

// Generate a #GIcon instance from @str. This function can fail if
// @str is not valid - see g_icon_to_string() for discussion.
//
// If your application or library provides one or more #GIcon
// implementations you need to ensure that each #GType is registered
// with the type system prior to calling g_icon_new_for_string().
/*

C function

g_icon_new_for_string
*/
func IconNewForString(str string) (*Icon, error) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var cThrowableError *C.GError

	retC := C.g_icon_new_for_string(c_str, &cThrowableError)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
