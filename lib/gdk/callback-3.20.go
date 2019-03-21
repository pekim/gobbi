// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

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

	void callback_seatgrabpreparefuncHandler(GObject *, GdkSeat*, GdkWindow*, gpointer);

*/
import "C"

var callbackSeatgrabpreparefuncId int
var callbackSeatgrabpreparefuncMap = make(map[int]SeatgrabpreparefuncCallback)
var callbackSeatgrabpreparefuncLock sync.RWMutex

// SeatgrabpreparefuncCallback is a callback function for a 'SeatGrabPrepareFunc' callback.
type SeatgrabpreparefuncCallback func(seat *Seat, window *Window)

//export callback_seatgrabpreparefuncHandler
func callback_seatgrabpreparefuncHandler(_ *C.GObject, c_seat *C.GdkSeat, c_window *C.GdkWindow, c_user_data C.gpointer) {
	callbackSeatgrabpreparefuncLock.RLock()
	defer callbackSeatgrabpreparefuncLock.RUnlock()

	seat := SeatNewFromC(unsafe.Pointer(c_seat))

	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(c_user_data))
	callback := callbackSeatgrabpreparefuncMap[index]
	callback(seat, window)
}
