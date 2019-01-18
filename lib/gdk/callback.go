// This is a generated file - DO NOT EDIT

package gdk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	gboolean callback_windowchildfuncHandler(GObject *, GdkWindow*, gpointer, gpointer);

*/
import "C"

// Unsupported : callback EventFunc : unsupported parameter event : no type generator for Event (GdkEvent*) for param event

// Unsupported : callback FilterFunc : unsupported parameter xevent : no type generator for XEvent (GdkXEvent*) for param xevent

var callbackWindowchildfuncId int
var callbackWindowchildfuncMap = make(map[int]WindowchildfuncCallback)
var callbackWindowchildfuncLock sync.RWMutex

// WindowchildfuncCallback is a callback function for a 'WindowChildFunc' callback.
type WindowchildfuncCallback func(window *Window) bool

//export callback_windowchildfuncHandler
func callback_windowchildfuncHandler(_ *C.GObject, c_window *C.GdkWindow, data C.gpointer) C.gboolean {
	callbackWindowchildfuncLock.RLock()
	defer callbackWindowchildfuncLock.RUnlock()

	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(data))
	callback := callbackWindowchildfuncMap[index].callback
	retGo := callback(window, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}
