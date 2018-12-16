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

// AppInfoResetTypeAssociations is a wrapper around the C function g_app_info_reset_type_associations.
func AppInfoResetTypeAssociations(contentType string) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	C.g_app_info_reset_type_associations(c_content_type)

	return
}

// CanDelete is a wrapper around the C function g_app_info_can_delete.
func (recv *AppInfo) CanDelete() bool {
	retC := C.g_app_info_can_delete((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Delete is a wrapper around the C function g_app_info_delete.
func (recv *AppInfo) Delete() bool {
	retC := C.g_app_info_delete((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCommandline is a wrapper around the C function g_app_info_get_commandline.
func (recv *AppInfo) GetCommandline() string {
	retC := C.g_app_info_get_commandline((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IconNewForString is a wrapper around the C function g_icon_new_for_string.
func IconNewForString(str string) (*Icon, error) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var cThrowableError *C.GError

	retC := C.g_icon_new_for_string(c_str, &cThrowableError)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ToString is a wrapper around the C function g_icon_to_string.
func (recv *Icon) ToString() string {
	retC := C.g_icon_to_string((*C.GIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsShadowed is a wrapper around the C function g_mount_is_shadowed.
func (recv *Mount) IsShadowed() bool {
	retC := C.g_mount_is_shadowed((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Shadow is a wrapper around the C function g_mount_shadow.
func (recv *Mount) Shadow() {
	C.g_mount_shadow((*C.GMount)(recv.native))

	return
}

// Unshadow is a wrapper around the C function g_mount_unshadow.
func (recv *Mount) Unshadow() {
	C.g_mount_unshadow((*C.GMount)(recv.native))

	return
}
