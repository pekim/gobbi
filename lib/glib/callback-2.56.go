// This is a generated file - DO NOT EDIT
// +build glib_2.56

package glib

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	void callback_clearhandlefuncHandler(GObject *, guint, gpointer);

*/
import "C"

var callbackClearhandlefuncId int
var callbackClearhandlefuncMap = make(map[int]ClearhandlefuncCallback)
var callbackClearhandlefuncLock sync.RWMutex

// ClearhandlefuncCallback is a callback function for a 'ClearHandleFunc' callback.
type ClearhandlefuncCallback func(handleId uint32)
