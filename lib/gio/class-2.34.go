// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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

	gboolean dbusauthobserver_allowMechanismHandler(GObject *, gchar*, gpointer);

	static gulong DBusAuthObserver_signal_connect_allow_mechanism(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "allow-mechanism", G_CALLBACK(dbusauthobserver_allowMechanismHandler), data);
	}

*/
/*

	void mountoperation_showUnmountProgressHandler(GObject *, gchar*, gint64, gint64, gpointer);

	static gulong MountOperation_signal_connect_show_unmount_progress(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-unmount-progress", G_CALLBACK(mountoperation_showUnmountProgressHandler), data);
	}

*/
import "C"

// Blacklisted : g_application_get_dbus_connection

// Blacklisted : g_application_get_dbus_object_path

// Blacklisted : g_application_command_line_get_stdin

type signalDBusAuthObserverAllowMechanismDetail struct {
	callback  DBusAuthObserverSignalAllowMechanismCallback
	handlerID C.gulong
}

var signalDBusAuthObserverAllowMechanismId int
var signalDBusAuthObserverAllowMechanismMap = make(map[int]signalDBusAuthObserverAllowMechanismDetail)
var signalDBusAuthObserverAllowMechanismLock sync.RWMutex

// DBusAuthObserverSignalAllowMechanismCallback is a callback function for a 'allow-mechanism' signal emitted from a DBusAuthObserver.
type DBusAuthObserverSignalAllowMechanismCallback func(mechanism string) bool

/*
ConnectAllowMechanism connects the callback to the 'allow-mechanism' signal for the DBusAuthObserver.

The returned value represents the connection, and may be passed to DisconnectAllowMechanism to remove it.
*/
func (recv *DBusAuthObserver) ConnectAllowMechanism(callback DBusAuthObserverSignalAllowMechanismCallback) int {
	signalDBusAuthObserverAllowMechanismLock.Lock()
	defer signalDBusAuthObserverAllowMechanismLock.Unlock()

	signalDBusAuthObserverAllowMechanismId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusAuthObserver_signal_connect_allow_mechanism(instance, C.gpointer(uintptr(signalDBusAuthObserverAllowMechanismId)))

	detail := signalDBusAuthObserverAllowMechanismDetail{callback, handlerID}
	signalDBusAuthObserverAllowMechanismMap[signalDBusAuthObserverAllowMechanismId] = detail

	return signalDBusAuthObserverAllowMechanismId
}

/*
DisconnectAllowMechanism disconnects a callback from the 'allow-mechanism' signal for the DBusAuthObserver.

The connectionID should be a value returned from a call to ConnectAllowMechanism.
*/
func (recv *DBusAuthObserver) DisconnectAllowMechanism(connectionID int) {
	signalDBusAuthObserverAllowMechanismLock.Lock()
	defer signalDBusAuthObserverAllowMechanismLock.Unlock()

	detail, exists := signalDBusAuthObserverAllowMechanismMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusAuthObserverAllowMechanismMap, connectionID)
}

//export dbusauthobserver_allowMechanismHandler
func dbusauthobserver_allowMechanismHandler(_ *C.GObject, c_mechanism *C.gchar, data C.gpointer) C.gboolean {
	signalDBusAuthObserverAllowMechanismLock.RLock()
	defer signalDBusAuthObserverAllowMechanismLock.RUnlock()

	mechanism := C.GoString(c_mechanism)

	index := int(uintptr(data))
	callback := signalDBusAuthObserverAllowMechanismMap[index].callback
	retGo := callback(mechanism)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_dbus_auth_observer_allow_mechanism

// Blacklisted : g_dbus_connection_get_last_serial

// Blacklisted : g_dbus_object_manager_server_is_exported

// Blacklisted : g_desktop_app_info_get_startup_wm_class

// Blacklisted : g_file_info_get_symbolic_icon

// Blacklisted : g_file_info_set_symbolic_icon

// Blacklisted : g_input_stream_read_bytes

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_input_stream_read_bytes_finish

// Blacklisted : g_memory_input_stream_new_from_bytes

// Blacklisted : g_memory_input_stream_add_bytes

// Blacklisted : g_memory_output_stream_steal_as_bytes

// Blacklisted : g_menu_item_new_from_model

// Blacklisted : g_menu_item_get_attribute

// Blacklisted : g_menu_item_get_attribute_value

// Blacklisted : g_menu_item_get_link

type signalMountOperationShowUnmountProgressDetail struct {
	callback  MountOperationSignalShowUnmountProgressCallback
	handlerID C.gulong
}

var signalMountOperationShowUnmountProgressId int
var signalMountOperationShowUnmountProgressMap = make(map[int]signalMountOperationShowUnmountProgressDetail)
var signalMountOperationShowUnmountProgressLock sync.RWMutex

// MountOperationSignalShowUnmountProgressCallback is a callback function for a 'show-unmount-progress' signal emitted from a MountOperation.
type MountOperationSignalShowUnmountProgressCallback func(message string, timeLeft int64, bytesLeft int64)

/*
ConnectShowUnmountProgress connects the callback to the 'show-unmount-progress' signal for the MountOperation.

The returned value represents the connection, and may be passed to DisconnectShowUnmountProgress to remove it.
*/
func (recv *MountOperation) ConnectShowUnmountProgress(callback MountOperationSignalShowUnmountProgressCallback) int {
	signalMountOperationShowUnmountProgressLock.Lock()
	defer signalMountOperationShowUnmountProgressLock.Unlock()

	signalMountOperationShowUnmountProgressId++
	instance := C.gpointer(recv.native)
	handlerID := C.MountOperation_signal_connect_show_unmount_progress(instance, C.gpointer(uintptr(signalMountOperationShowUnmountProgressId)))

	detail := signalMountOperationShowUnmountProgressDetail{callback, handlerID}
	signalMountOperationShowUnmountProgressMap[signalMountOperationShowUnmountProgressId] = detail

	return signalMountOperationShowUnmountProgressId
}

/*
DisconnectShowUnmountProgress disconnects a callback from the 'show-unmount-progress' signal for the MountOperation.

The connectionID should be a value returned from a call to ConnectShowUnmountProgress.
*/
func (recv *MountOperation) DisconnectShowUnmountProgress(connectionID int) {
	signalMountOperationShowUnmountProgressLock.Lock()
	defer signalMountOperationShowUnmountProgressLock.Unlock()

	detail, exists := signalMountOperationShowUnmountProgressMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountOperationShowUnmountProgressMap, connectionID)
}

//export mountoperation_showUnmountProgressHandler
func mountoperation_showUnmountProgressHandler(_ *C.GObject, c_message *C.gchar, c_time_left C.gint64, c_bytes_left C.gint64, data C.gpointer) {
	signalMountOperationShowUnmountProgressLock.RLock()
	defer signalMountOperationShowUnmountProgressLock.RUnlock()

	message := C.GoString(c_message)

	timeLeft := int64(c_time_left)

	bytesLeft := int64(c_bytes_left)

	index := int(uintptr(data))
	callback := signalMountOperationShowUnmountProgressMap[index].callback
	callback(message, timeLeft, bytesLeft)
}

// Blacklisted : g_proxy_address_get_destination_protocol

// Blacklisted : g_proxy_address_get_uri

// Blacklisted : g_resolver_lookup_records

// Unsupported : g_resolver_lookup_records_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_resolver_lookup_records_finish

// TestDBus is a wrapper around the C record GTestDBus.
type TestDBus struct {
	native *C.GTestDBus
}

func TestDBusNewFromC(u unsafe.Pointer) *TestDBus {
	c := (*C.GTestDBus)(u)
	if c == nil {
		return nil
	}

	g := &TestDBus{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TestDBus) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TestDBus) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestDBus with another TestDBus, and returns true if they represent the same GObject.
func (recv *TestDBus) Equals(other *TestDBus) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *TestDBus) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to TestDBus.
// Exercise care, as this is a potentially dangerous function if the Object is not a TestDBus.
func CastToTestDBus(object *gobject.Object) *TestDBus {
	return TestDBusNewFromC(object.ToC())
}

// Blacklisted : g_test_dbus_new

// Blacklisted : g_test_dbus_add_service_dir

// Blacklisted : g_test_dbus_down

// Blacklisted : g_test_dbus_get_bus_address

// Blacklisted : g_test_dbus_get_flags

// Blacklisted : g_test_dbus_stop

// Blacklisted : g_test_dbus_up

// Blacklisted : g_tls_certificate_is_same
