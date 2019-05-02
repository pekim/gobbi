// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	void socketlistener_eventHandler(GObject *, GSocketListenerEvent, GSocket *, gpointer);

	static gulong SocketListener_signal_connect_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "event", G_CALLBACK(socketlistener_eventHandler), data);
	}

*/
import "C"

// RegisterObjectWithClosures is a wrapper around the C function g_dbus_connection_register_object_with_closures.
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_list_store_sort : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

type signalSocketListenerEventDetail struct {
	callback  SocketListenerSignalEventCallback
	handlerID C.gulong
}

var signalSocketListenerEventId int
var signalSocketListenerEventMap = make(map[int]signalSocketListenerEventDetail)
var signalSocketListenerEventLock sync.RWMutex

// SocketListenerSignalEventCallback is a callback function for a 'event' signal emitted from a SocketListener.
type SocketListenerSignalEventCallback func(event SocketListenerEvent, socket *Socket)

/*
ConnectEvent connects the callback to the 'event' signal for the SocketListener.

The returned value represents the connection, and may be passed to DisconnectEvent to remove it.
*/
func (recv *SocketListener) ConnectEvent(callback SocketListenerSignalEventCallback) int {
	signalSocketListenerEventLock.Lock()
	defer signalSocketListenerEventLock.Unlock()

	signalSocketListenerEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.SocketListener_signal_connect_event(instance, C.gpointer(uintptr(signalSocketListenerEventId)))

	detail := signalSocketListenerEventDetail{callback, handlerID}
	signalSocketListenerEventMap[signalSocketListenerEventId] = detail

	return signalSocketListenerEventId
}

/*
DisconnectEvent disconnects a callback from the 'event' signal for the SocketListener.

The connectionID should be a value returned from a call to ConnectEvent.
*/
func (recv *SocketListener) DisconnectEvent(connectionID int) {
	signalSocketListenerEventLock.Lock()
	defer signalSocketListenerEventLock.Unlock()

	detail, exists := signalSocketListenerEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketListenerEventMap, connectionID)
}

//export socketlistener_eventHandler
func socketlistener_eventHandler(_ *C.GObject, c_event C.GSocketListenerEvent, c_socket *C.GSocket, data C.gpointer) {
	signalSocketListenerEventLock.RLock()
	defer signalSocketListenerEventLock.RUnlock()

	event := SocketListenerEvent(c_event)

	socket := SocketNewFromC(unsafe.Pointer(c_socket))

	index := int(uintptr(data))
	callback := signalSocketListenerEventMap[index].callback
	callback(event, socket)
}
