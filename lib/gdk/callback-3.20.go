// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void callback_seatgrabpreparefuncHandler(GObject *, GdkSeat*, GdkWindow*, gpointer, gpointer);

*/
import "C"

var callbackSeatgrabpreparefuncId int
var callbackSeatgrabpreparefuncMap = make(map[int]SeatgrabpreparefuncCallback)
var callbackSeatgrabpreparefuncLock sync.RWMutex

// SeatgrabpreparefuncCallback is a callback function for a 'SeatGrabPrepareFunc' callback.
type SeatgrabpreparefuncCallback func(seat *Seat, window *Window)
