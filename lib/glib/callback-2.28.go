// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

	void callback_testdatafuncHandler(GObject *, gconstpointer, gpointer);

*/
/*

	void callback_testfuncHandler(GObject *, gpointer);

*/
import "C"

var callbackTestdatafuncId int
var callbackTestdatafuncMap = make(map[int]TestdatafuncCallback)
var callbackTestdatafuncLock sync.RWMutex

// TestdatafuncCallback is a callback function for a 'TestDataFunc' callback.
type TestdatafuncCallback func()

//export callback_testdatafuncHandler
func callback_testdatafuncHandler(_ *C.GObject, data C.gpointer) {
	callbackTestdatafuncLock.RLock()
	defer callbackTestdatafuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackTestdatafuncMap[index].callback
	callback(userData)
}

// Unsupported : callback TestFixtureFunc : unsupported parameter fixture : no type generator for gpointer (gpointer) for param fixture

var callbackTestfuncId int
var callbackTestfuncMap = make(map[int]TestfuncCallback)
var callbackTestfuncLock sync.RWMutex

// TestfuncCallback is a callback function for a 'TestFunc' callback.
type TestfuncCallback func()

//export callback_testfuncHandler
func callback_testfuncHandler(_ *C.GObject, data C.gpointer) {
	callbackTestfuncLock.RLock()
	defer callbackTestfuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackTestfuncMap[index].callback
	callback()
}
