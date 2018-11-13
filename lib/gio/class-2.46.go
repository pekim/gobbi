// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Version of g_dbus_connection_register_object() using closures instead of a
// #GDBusInterfaceVTable for easier binding in other languages.
/*

C function : g_dbus_connection_register_object_with_closures
*/
func (recv *DBusConnection) RegisterObjectWithClosures(objectPath string, interfaceInfo *DBusInterfaceInfo, methodCallClosure *gobject.Closure, getPropertyClosure *gobject.Closure, setPropertyClosure *gobject.Closure) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if interfaceInfo != nil {
		c_interface_info = (*C.GDBusInterfaceInfo)(interfaceInfo.ToC())
	}

	c_method_call_closure := (*C.GClosure)(C.NULL)
	if methodCallClosure != nil {
		c_method_call_closure = (*C.GClosure)(methodCallClosure.ToC())
	}

	c_get_property_closure := (*C.GClosure)(C.NULL)
	if getPropertyClosure != nil {
		c_get_property_closure = (*C.GClosure)(getPropertyClosure.ToC())
	}

	c_set_property_closure := (*C.GClosure)(C.NULL)
	if setPropertyClosure != nil {
		c_set_property_closure = (*C.GClosure)(setPropertyClosure.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_register_object_with_closures((*C.GDBusConnection)(recv.native), c_object_path, c_interface_info, c_method_call_closure, c_get_property_closure, c_set_property_closure, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_list_store_sort : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func
