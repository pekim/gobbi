// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
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

	GType callback_dbusproxytypefuncHandler(GObject *, GDBusObjectManagerClient*, gchar*, gchar*, gpointer);

*/
import "C"

var callbackDbusproxytypefuncId int
var callbackDbusproxytypefuncMap = make(map[int]DbusproxytypefuncCallback)
var callbackDbusproxytypefuncLock sync.RWMutex

// DbusproxytypefuncCallback is a callback function for a 'DBusProxyTypeFunc' callback.
type DbusproxytypefuncCallback func(manager *DBusObjectManagerClient, objectPath string, interfaceName string) gobject.Type

//export callback_dbusproxytypefuncHandler
func callback_dbusproxytypefuncHandler(_ *C.GObject, c_manager *C.GDBusObjectManagerClient, c_object_path *C.gchar, c_interface_name *C.gchar, c_user_data C.gpointer) C.GType {
	callbackDbusproxytypefuncLock.RLock()
	defer callbackDbusproxytypefuncLock.RUnlock()

	manager := DBusObjectManagerClientNewFromC(unsafe.Pointer(c_manager))

	objectPath := C.GoString(c_object_path)

	interfaceName := C.GoString(c_interface_name)

	index := int(uintptr(c_user_data))
	callback := callbackDbusproxytypefuncMap[index]
	retGo := callback(manager, objectPath, interfaceName)
	retC :=
		(C.GType)(retGo)
	return retC
}
