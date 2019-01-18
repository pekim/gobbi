// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <glib-object.h>
// #include <stdlib.h>
/*

	gboolean callback_testlogfatalfuncHandler(GObject *, const gchar*, GLogLevelFlags, const gchar*, gpointer, gpointer);

*/
import "C"

var callbackTestlogfatalfuncId int
var callbackTestlogfatalfuncMap = make(map[int]TestlogfatalfuncCallback)
var callbackTestlogfatalfuncLock sync.RWMutex

// TestlogfatalfuncCallback is a callback function for a 'TestLogFatalFunc' callback.
type TestlogfatalfuncCallback func(logDomain string, logLevel LogLevelFlags, message string) bool

//export callback_testlogfatalfuncHandler
func callback_testlogfatalfuncHandler(_ *C.GObject, c_log_domain *C.gchar, c_log_level C.GLogLevelFlags, c_message *C.gchar, data C.gpointer) C.gboolean {
	callbackTestlogfatalfuncLock.RLock()
	defer callbackTestlogfatalfuncLock.RUnlock()

	logDomain := C.GoString(c_log_domain)

	logLevel := LogLevelFlags(c_log_level)

	message := C.GoString(c_message)

	index := int(uintptr(data))
	callback := callbackTestlogfatalfuncMap[index].callback
	retGo := callback(logDomain, logLevel, message, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}
