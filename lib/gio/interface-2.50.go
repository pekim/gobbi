// This is a generated file - DO NOT EDIT
// +build gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// g_app_info_launch_default_for_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// AppInfoLaunchDefaultForUriFinish is a wrapper around the C function g_app_info_launch_default_for_uri_finish.
func AppInfoLaunchDefaultForUriFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_default_for_uri_finish(c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IsRemovable is a wrapper around the C function g_drive_is_removable.
func (recv *Drive) IsRemovable() bool {
	retC := C.g_drive_is_removable((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
