// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
/*

	gboolean callback_socketsourcefuncHandler(GObject *, GSocket*, GIOCondition, gpointer);

*/
import "C"

var callbackSocketsourcefuncId int
var callbackSocketsourcefuncMap = make(map[int]SocketsourcefuncCallback)
var callbackSocketsourcefuncLock sync.RWMutex

// SocketsourcefuncCallback is a callback function for a 'SocketSourceFunc' callback.
type SocketsourcefuncCallback func(socket *Socket, condition glib.IOCondition) bool

//export callback_socketsourcefuncHandler
func callback_socketsourcefuncHandler(_ *C.GObject, c_socket *C.GSocket, c_condition C.GIOCondition, c_user_data C.gpointer) C.gboolean {
	callbackSocketsourcefuncLock.RLock()
	defer callbackSocketsourcefuncLock.RUnlock()

	socket := SocketNewFromC(unsafe.Pointer(c_socket))

	condition := glib.IOCondition(c_condition)

	index := int(uintptr(c_user_data))
	callback := callbackSocketsourcefuncMap[index]
	retGo := callback(socket, condition)
	retC :=
		boolToGboolean(retGo)
	return retC
}
