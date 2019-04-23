// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

	gboolean dbusinterfaceskeleton_gAuthorizeMethodHandler(GObject *, GDBusMethodInvocation *, gpointer);

	static gulong DBusInterfaceSkeleton_signal_connect_g_authorize_method(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "g-authorize-method", G_CALLBACK(dbusinterfaceskeleton_gAuthorizeMethodHandler), data);
	}

*/
/*

	void dbusobjectmanagerclient_interfaceProxySignalHandler(GObject *, GDBusObjectProxy *, GDBusProxy *, gchar*, gchar*, GVariant *, gpointer);

	static gulong DBusObjectManagerClient_signal_connect_interface_proxy_signal(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "interface-proxy-signal", G_CALLBACK(dbusobjectmanagerclient_interfaceProxySignalHandler), data);
	}

*/
/*

	gboolean dbusobjectskeleton_authorizeMethodHandler(GObject *, GDBusInterfaceSkeleton *, GDBusMethodInvocation *, gpointer);

	static gulong DBusObjectSkeleton_signal_connect_authorize_method(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "authorize-method", G_CALLBACK(dbusobjectskeleton_authorizeMethodHandler), data);
	}

*/
/*

	void simpleaction_changeStateHandler(GObject *, GVariant *, gpointer);

	static gulong SimpleAction_signal_connect_change_state(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-state", G_CALLBACK(simpleaction_changeStateHandler), data);
	}

*/
import "C"

// Unsupported : g_dbus_connection_call_with_unix_fd_list : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_connection_call_with_unix_fd_list_finish

// Blacklisted : g_dbus_connection_call_with_unix_fd_list_sync

// DBusInterfaceSkeleton is a wrapper around the C record GDBusInterfaceSkeleton.
type DBusInterfaceSkeleton struct {
	native *C.GDBusInterfaceSkeleton
	// Private : parent_instance
	// Private : priv
}

func DBusInterfaceSkeletonNewFromC(u unsafe.Pointer) *DBusInterfaceSkeleton {
	c := (*C.GDBusInterfaceSkeleton)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeleton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusInterfaceSkeleton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusInterfaceSkeleton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusInterfaceSkeleton with another DBusInterfaceSkeleton, and returns true if they represent the same GObject.
func (recv *DBusInterfaceSkeleton) Equals(other *DBusInterfaceSkeleton) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusInterfaceSkeleton) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusInterfaceSkeleton.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusInterfaceSkeleton.
func CastToDBusInterfaceSkeleton(object *gobject.Object) *DBusInterfaceSkeleton {
	return DBusInterfaceSkeletonNewFromC(object.ToC())
}

type signalDBusInterfaceSkeletonGAuthorizeMethodDetail struct {
	callback  DBusInterfaceSkeletonSignalGAuthorizeMethodCallback
	handlerID C.gulong
}

var signalDBusInterfaceSkeletonGAuthorizeMethodId int
var signalDBusInterfaceSkeletonGAuthorizeMethodMap = make(map[int]signalDBusInterfaceSkeletonGAuthorizeMethodDetail)
var signalDBusInterfaceSkeletonGAuthorizeMethodLock sync.RWMutex

// DBusInterfaceSkeletonSignalGAuthorizeMethodCallback is a callback function for a 'g-authorize-method' signal emitted from a DBusInterfaceSkeleton.
type DBusInterfaceSkeletonSignalGAuthorizeMethodCallback func(invocation *DBusMethodInvocation) bool

/*
ConnectGAuthorizeMethod connects the callback to the 'g-authorize-method' signal for the DBusInterfaceSkeleton.

The returned value represents the connection, and may be passed to DisconnectGAuthorizeMethod to remove it.
*/
func (recv *DBusInterfaceSkeleton) ConnectGAuthorizeMethod(callback DBusInterfaceSkeletonSignalGAuthorizeMethodCallback) int {
	signalDBusInterfaceSkeletonGAuthorizeMethodLock.Lock()
	defer signalDBusInterfaceSkeletonGAuthorizeMethodLock.Unlock()

	signalDBusInterfaceSkeletonGAuthorizeMethodId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusInterfaceSkeleton_signal_connect_g_authorize_method(instance, C.gpointer(uintptr(signalDBusInterfaceSkeletonGAuthorizeMethodId)))

	detail := signalDBusInterfaceSkeletonGAuthorizeMethodDetail{callback, handlerID}
	signalDBusInterfaceSkeletonGAuthorizeMethodMap[signalDBusInterfaceSkeletonGAuthorizeMethodId] = detail

	return signalDBusInterfaceSkeletonGAuthorizeMethodId
}

/*
DisconnectGAuthorizeMethod disconnects a callback from the 'g-authorize-method' signal for the DBusInterfaceSkeleton.

The connectionID should be a value returned from a call to ConnectGAuthorizeMethod.
*/
func (recv *DBusInterfaceSkeleton) DisconnectGAuthorizeMethod(connectionID int) {
	signalDBusInterfaceSkeletonGAuthorizeMethodLock.Lock()
	defer signalDBusInterfaceSkeletonGAuthorizeMethodLock.Unlock()

	detail, exists := signalDBusInterfaceSkeletonGAuthorizeMethodMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusInterfaceSkeletonGAuthorizeMethodMap, connectionID)
}

//export dbusinterfaceskeleton_gAuthorizeMethodHandler
func dbusinterfaceskeleton_gAuthorizeMethodHandler(_ *C.GObject, c_invocation *C.GDBusMethodInvocation, data C.gpointer) C.gboolean {
	signalDBusInterfaceSkeletonGAuthorizeMethodLock.RLock()
	defer signalDBusInterfaceSkeletonGAuthorizeMethodLock.RUnlock()

	invocation := DBusMethodInvocationNewFromC(unsafe.Pointer(c_invocation))

	index := int(uintptr(data))
	callback := signalDBusInterfaceSkeletonGAuthorizeMethodMap[index].callback
	retGo := callback(invocation)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_dbus_interface_skeleton_export

// Flush is a wrapper around the C function g_dbus_interface_skeleton_flush.
func (recv *DBusInterfaceSkeleton) Flush() {
	C.g_dbus_interface_skeleton_flush((*C.GDBusInterfaceSkeleton)(recv.native))

	return
}

// Blacklisted : g_dbus_interface_skeleton_get_connection

// Blacklisted : g_dbus_interface_skeleton_get_flags

// Blacklisted : g_dbus_interface_skeleton_get_info

// Blacklisted : g_dbus_interface_skeleton_get_object_path

// Blacklisted : g_dbus_interface_skeleton_get_properties

// Blacklisted : g_dbus_interface_skeleton_get_vtable

// Blacklisted : g_dbus_interface_skeleton_set_flags

// Blacklisted : g_dbus_interface_skeleton_unexport

// Blacklisted : g_dbus_method_invocation_return_value_with_unix_fd_list

// Blacklisted : g_dbus_method_invocation_take_error

// DBusObjectManagerClient is a wrapper around the C record GDBusObjectManagerClient.
type DBusObjectManagerClient struct {
	native *C.GDBusObjectManagerClient
	// Private : parent_instance
	// Private : priv
}

func DBusObjectManagerClientNewFromC(u unsafe.Pointer) *DBusObjectManagerClient {
	c := (*C.GDBusObjectManagerClient)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClient{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusObjectManagerClient) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusObjectManagerClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectManagerClient with another DBusObjectManagerClient, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerClient) Equals(other *DBusObjectManagerClient) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusObjectManagerClient) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusObjectManagerClient.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectManagerClient.
func CastToDBusObjectManagerClient(object *gobject.Object) *DBusObjectManagerClient {
	return DBusObjectManagerClientNewFromC(object.ToC())
}

// Unsupported signal 'interface-proxy-properties-changed' for DBusObjectManagerClient : unsupported parameter invalidated_properties :

type signalDBusObjectManagerClientInterfaceProxySignalDetail struct {
	callback  DBusObjectManagerClientSignalInterfaceProxySignalCallback
	handlerID C.gulong
}

var signalDBusObjectManagerClientInterfaceProxySignalId int
var signalDBusObjectManagerClientInterfaceProxySignalMap = make(map[int]signalDBusObjectManagerClientInterfaceProxySignalDetail)
var signalDBusObjectManagerClientInterfaceProxySignalLock sync.RWMutex

// DBusObjectManagerClientSignalInterfaceProxySignalCallback is a callback function for a 'interface-proxy-signal' signal emitted from a DBusObjectManagerClient.
type DBusObjectManagerClientSignalInterfaceProxySignalCallback func(objectProxy *DBusObjectProxy, interfaceProxy *DBusProxy, senderName string, signalName string, parameters *glib.Variant)

/*
ConnectInterfaceProxySignal connects the callback to the 'interface-proxy-signal' signal for the DBusObjectManagerClient.

The returned value represents the connection, and may be passed to DisconnectInterfaceProxySignal to remove it.
*/
func (recv *DBusObjectManagerClient) ConnectInterfaceProxySignal(callback DBusObjectManagerClientSignalInterfaceProxySignalCallback) int {
	signalDBusObjectManagerClientInterfaceProxySignalLock.Lock()
	defer signalDBusObjectManagerClientInterfaceProxySignalLock.Unlock()

	signalDBusObjectManagerClientInterfaceProxySignalId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObjectManagerClient_signal_connect_interface_proxy_signal(instance, C.gpointer(uintptr(signalDBusObjectManagerClientInterfaceProxySignalId)))

	detail := signalDBusObjectManagerClientInterfaceProxySignalDetail{callback, handlerID}
	signalDBusObjectManagerClientInterfaceProxySignalMap[signalDBusObjectManagerClientInterfaceProxySignalId] = detail

	return signalDBusObjectManagerClientInterfaceProxySignalId
}

/*
DisconnectInterfaceProxySignal disconnects a callback from the 'interface-proxy-signal' signal for the DBusObjectManagerClient.

The connectionID should be a value returned from a call to ConnectInterfaceProxySignal.
*/
func (recv *DBusObjectManagerClient) DisconnectInterfaceProxySignal(connectionID int) {
	signalDBusObjectManagerClientInterfaceProxySignalLock.Lock()
	defer signalDBusObjectManagerClientInterfaceProxySignalLock.Unlock()

	detail, exists := signalDBusObjectManagerClientInterfaceProxySignalMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectManagerClientInterfaceProxySignalMap, connectionID)
}

//export dbusobjectmanagerclient_interfaceProxySignalHandler
func dbusobjectmanagerclient_interfaceProxySignalHandler(_ *C.GObject, c_object_proxy *C.GDBusObjectProxy, c_interface_proxy *C.GDBusProxy, c_sender_name *C.gchar, c_signal_name *C.gchar, c_parameters *C.GVariant, data C.gpointer) {
	signalDBusObjectManagerClientInterfaceProxySignalLock.RLock()
	defer signalDBusObjectManagerClientInterfaceProxySignalLock.RUnlock()

	objectProxy := DBusObjectProxyNewFromC(unsafe.Pointer(c_object_proxy))

	interfaceProxy := DBusProxyNewFromC(unsafe.Pointer(c_interface_proxy))

	senderName := C.GoString(c_sender_name)

	signalName := C.GoString(c_signal_name)

	parameters := glib.VariantNewFromC(unsafe.Pointer(c_parameters))

	index := int(uintptr(data))
	callback := signalDBusObjectManagerClientInterfaceProxySignalMap[index].callback
	callback(objectProxy, interfaceProxy, senderName, signalName, parameters)
}

// Blacklisted : g_dbus_object_manager_client_new_finish

// Blacklisted : g_dbus_object_manager_client_new_for_bus_finish

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// g_dbus_object_manager_client_new : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func
// g_dbus_object_manager_client_new_for_bus : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func
// Blacklisted : g_dbus_object_manager_client_get_connection

// Blacklisted : g_dbus_object_manager_client_get_flags

// Blacklisted : g_dbus_object_manager_client_get_name

// Blacklisted : g_dbus_object_manager_client_get_name_owner

// DBusObjectManagerServer is a wrapper around the C record GDBusObjectManagerServer.
type DBusObjectManagerServer struct {
	native *C.GDBusObjectManagerServer
	// Private : parent_instance
	// Private : priv
}

func DBusObjectManagerServerNewFromC(u unsafe.Pointer) *DBusObjectManagerServer {
	c := (*C.GDBusObjectManagerServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusObjectManagerServer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusObjectManagerServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectManagerServer with another DBusObjectManagerServer, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerServer) Equals(other *DBusObjectManagerServer) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusObjectManagerServer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusObjectManagerServer.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectManagerServer.
func CastToDBusObjectManagerServer(object *gobject.Object) *DBusObjectManagerServer {
	return DBusObjectManagerServerNewFromC(object.ToC())
}

// Blacklisted : g_dbus_object_manager_server_new

// Blacklisted : g_dbus_object_manager_server_export

// Blacklisted : g_dbus_object_manager_server_export_uniquely

// Blacklisted : g_dbus_object_manager_server_get_connection

// Blacklisted : g_dbus_object_manager_server_set_connection

// Blacklisted : g_dbus_object_manager_server_unexport

// DBusObjectProxy is a wrapper around the C record GDBusObjectProxy.
type DBusObjectProxy struct {
	native *C.GDBusObjectProxy
	// Private : parent_instance
	// Private : priv
}

func DBusObjectProxyNewFromC(u unsafe.Pointer) *DBusObjectProxy {
	c := (*C.GDBusObjectProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxy{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusObjectProxy) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusObjectProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectProxy with another DBusObjectProxy, and returns true if they represent the same GObject.
func (recv *DBusObjectProxy) Equals(other *DBusObjectProxy) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusObjectProxy) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusObjectProxy.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectProxy.
func CastToDBusObjectProxy(object *gobject.Object) *DBusObjectProxy {
	return DBusObjectProxyNewFromC(object.ToC())
}

// Blacklisted : g_dbus_object_proxy_new

// Blacklisted : g_dbus_object_proxy_get_connection

// DBusObjectSkeleton is a wrapper around the C record GDBusObjectSkeleton.
type DBusObjectSkeleton struct {
	native *C.GDBusObjectSkeleton
	// Private : parent_instance
	// Private : priv
}

func DBusObjectSkeletonNewFromC(u unsafe.Pointer) *DBusObjectSkeleton {
	c := (*C.GDBusObjectSkeleton)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeleton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusObjectSkeleton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusObjectSkeleton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectSkeleton with another DBusObjectSkeleton, and returns true if they represent the same GObject.
func (recv *DBusObjectSkeleton) Equals(other *DBusObjectSkeleton) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusObjectSkeleton) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusObjectSkeleton.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectSkeleton.
func CastToDBusObjectSkeleton(object *gobject.Object) *DBusObjectSkeleton {
	return DBusObjectSkeletonNewFromC(object.ToC())
}

type signalDBusObjectSkeletonAuthorizeMethodDetail struct {
	callback  DBusObjectSkeletonSignalAuthorizeMethodCallback
	handlerID C.gulong
}

var signalDBusObjectSkeletonAuthorizeMethodId int
var signalDBusObjectSkeletonAuthorizeMethodMap = make(map[int]signalDBusObjectSkeletonAuthorizeMethodDetail)
var signalDBusObjectSkeletonAuthorizeMethodLock sync.RWMutex

// DBusObjectSkeletonSignalAuthorizeMethodCallback is a callback function for a 'authorize-method' signal emitted from a DBusObjectSkeleton.
type DBusObjectSkeletonSignalAuthorizeMethodCallback func(interface_ *DBusInterfaceSkeleton, invocation *DBusMethodInvocation) bool

/*
ConnectAuthorizeMethod connects the callback to the 'authorize-method' signal for the DBusObjectSkeleton.

The returned value represents the connection, and may be passed to DisconnectAuthorizeMethod to remove it.
*/
func (recv *DBusObjectSkeleton) ConnectAuthorizeMethod(callback DBusObjectSkeletonSignalAuthorizeMethodCallback) int {
	signalDBusObjectSkeletonAuthorizeMethodLock.Lock()
	defer signalDBusObjectSkeletonAuthorizeMethodLock.Unlock()

	signalDBusObjectSkeletonAuthorizeMethodId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObjectSkeleton_signal_connect_authorize_method(instance, C.gpointer(uintptr(signalDBusObjectSkeletonAuthorizeMethodId)))

	detail := signalDBusObjectSkeletonAuthorizeMethodDetail{callback, handlerID}
	signalDBusObjectSkeletonAuthorizeMethodMap[signalDBusObjectSkeletonAuthorizeMethodId] = detail

	return signalDBusObjectSkeletonAuthorizeMethodId
}

/*
DisconnectAuthorizeMethod disconnects a callback from the 'authorize-method' signal for the DBusObjectSkeleton.

The connectionID should be a value returned from a call to ConnectAuthorizeMethod.
*/
func (recv *DBusObjectSkeleton) DisconnectAuthorizeMethod(connectionID int) {
	signalDBusObjectSkeletonAuthorizeMethodLock.Lock()
	defer signalDBusObjectSkeletonAuthorizeMethodLock.Unlock()

	detail, exists := signalDBusObjectSkeletonAuthorizeMethodMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectSkeletonAuthorizeMethodMap, connectionID)
}

//export dbusobjectskeleton_authorizeMethodHandler
func dbusobjectskeleton_authorizeMethodHandler(_ *C.GObject, c_interface *C.GDBusInterfaceSkeleton, c_invocation *C.GDBusMethodInvocation, data C.gpointer) C.gboolean {
	signalDBusObjectSkeletonAuthorizeMethodLock.RLock()
	defer signalDBusObjectSkeletonAuthorizeMethodLock.RUnlock()

	interface_ := DBusInterfaceSkeletonNewFromC(unsafe.Pointer(c_interface))

	invocation := DBusMethodInvocationNewFromC(unsafe.Pointer(c_invocation))

	index := int(uintptr(data))
	callback := signalDBusObjectSkeletonAuthorizeMethodMap[index].callback
	retGo := callback(interface_, invocation)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_dbus_object_skeleton_new

// Blacklisted : g_dbus_object_skeleton_add_interface

// Flush is a wrapper around the C function g_dbus_object_skeleton_flush.
func (recv *DBusObjectSkeleton) Flush() {
	C.g_dbus_object_skeleton_flush((*C.GDBusObjectSkeleton)(recv.native))

	return
}

// Blacklisted : g_dbus_object_skeleton_remove_interface

// Blacklisted : g_dbus_object_skeleton_remove_interface_by_name

// Blacklisted : g_dbus_object_skeleton_set_object_path

// Unsupported : g_dbus_proxy_call_with_unix_fd_list : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_proxy_call_with_unix_fd_list_finish

// Blacklisted : g_dbus_proxy_call_with_unix_fd_list_sync

// Blacklisted : g_data_input_stream_read_line_finish_utf8

// Blacklisted : g_data_input_stream_read_line_utf8

// Blacklisted : g_desktop_app_info_get_nodisplay

// Blacklisted : g_desktop_app_info_get_show_in

// Blacklisted : g_inet_address_equal

// Blacklisted : g_settings_get_uint

// Blacklisted : g_settings_set_uint

type signalSimpleActionChangeStateDetail struct {
	callback  SimpleActionSignalChangeStateCallback
	handlerID C.gulong
}

var signalSimpleActionChangeStateId int
var signalSimpleActionChangeStateMap = make(map[int]signalSimpleActionChangeStateDetail)
var signalSimpleActionChangeStateLock sync.RWMutex

// SimpleActionSignalChangeStateCallback is a callback function for a 'change-state' signal emitted from a SimpleAction.
type SimpleActionSignalChangeStateCallback func(value *glib.Variant)

/*
ConnectChangeState connects the callback to the 'change-state' signal for the SimpleAction.

The returned value represents the connection, and may be passed to DisconnectChangeState to remove it.
*/
func (recv *SimpleAction) ConnectChangeState(callback SimpleActionSignalChangeStateCallback) int {
	signalSimpleActionChangeStateLock.Lock()
	defer signalSimpleActionChangeStateLock.Unlock()

	signalSimpleActionChangeStateId++
	instance := C.gpointer(recv.native)
	handlerID := C.SimpleAction_signal_connect_change_state(instance, C.gpointer(uintptr(signalSimpleActionChangeStateId)))

	detail := signalSimpleActionChangeStateDetail{callback, handlerID}
	signalSimpleActionChangeStateMap[signalSimpleActionChangeStateId] = detail

	return signalSimpleActionChangeStateId
}

/*
DisconnectChangeState disconnects a callback from the 'change-state' signal for the SimpleAction.

The connectionID should be a value returned from a call to ConnectChangeState.
*/
func (recv *SimpleAction) DisconnectChangeState(connectionID int) {
	signalSimpleActionChangeStateLock.Lock()
	defer signalSimpleActionChangeStateLock.Unlock()

	detail, exists := signalSimpleActionChangeStateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSimpleActionChangeStateMap, connectionID)
}

//export simpleaction_changeStateHandler
func simpleaction_changeStateHandler(_ *C.GObject, c_value *C.GVariant, data C.gpointer) {
	signalSimpleActionChangeStateLock.RLock()
	defer signalSimpleActionChangeStateLock.RUnlock()

	value := glib.VariantNewFromC(unsafe.Pointer(c_value))

	index := int(uintptr(data))
	callback := signalSimpleActionChangeStateMap[index].callback
	callback(value)
}

// Blacklisted : g_simple_action_set_state

// Unsupported : g_simple_action_group_add_entries : unsupported parameter entries :

// Blacklisted : g_tls_connection_get_database

// Blacklisted : g_tls_connection_get_interaction

// Blacklisted : g_tls_connection_set_database

// Blacklisted : g_tls_connection_set_interaction

// TlsDatabase is a wrapper around the C record GTlsDatabase.
type TlsDatabase struct {
	native *C.GTlsDatabase
	// parent_instance : record
	// priv : record
}

func TlsDatabaseNewFromC(u unsafe.Pointer) *TlsDatabase {
	c := (*C.GTlsDatabase)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabase{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsDatabase) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsDatabase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsDatabase with another TlsDatabase, and returns true if they represent the same GObject.
func (recv *TlsDatabase) Equals(other *TlsDatabase) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *TlsDatabase) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to TlsDatabase.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsDatabase.
func CastToTlsDatabase(object *gobject.Object) *TlsDatabase {
	return TlsDatabaseNewFromC(object.ToC())
}

// Blacklisted : g_tls_database_create_certificate_handle

// Blacklisted : g_tls_database_lookup_certificate_for_handle

// Unsupported : g_tls_database_lookup_certificate_for_handle_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_tls_database_lookup_certificate_for_handle_finish

// Blacklisted : g_tls_database_lookup_certificate_issuer

// Unsupported : g_tls_database_lookup_certificate_issuer_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_tls_database_lookup_certificate_issuer_finish

// Blacklisted : g_tls_database_lookup_certificates_issued_by

// Unsupported : g_tls_database_lookup_certificates_issued_by_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_tls_database_lookup_certificates_issued_by_finish

// Blacklisted : g_tls_database_verify_chain

// Unsupported : g_tls_database_verify_chain_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_tls_database_verify_chain_finish

// TlsInteraction is a wrapper around the C record GTlsInteraction.
type TlsInteraction struct {
	native *C.GTlsInteraction
	// Private : parent_instance
	// Private : priv
}

func TlsInteractionNewFromC(u unsafe.Pointer) *TlsInteraction {
	c := (*C.GTlsInteraction)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteraction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsInteraction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsInteraction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsInteraction with another TlsInteraction, and returns true if they represent the same GObject.
func (recv *TlsInteraction) Equals(other *TlsInteraction) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *TlsInteraction) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to TlsInteraction.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsInteraction.
func CastToTlsInteraction(object *gobject.Object) *TlsInteraction {
	return TlsInteractionNewFromC(object.ToC())
}

// Blacklisted : g_tls_interaction_ask_password

// Unsupported : g_tls_interaction_ask_password_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_tls_interaction_ask_password_finish

// Blacklisted : g_tls_interaction_invoke_ask_password

// TlsPassword is a wrapper around the C record GTlsPassword.
type TlsPassword struct {
	native *C.GTlsPassword
	// parent_instance : record
	// priv : record
}

func TlsPasswordNewFromC(u unsafe.Pointer) *TlsPassword {
	c := (*C.GTlsPassword)(u)
	if c == nil {
		return nil
	}

	g := &TlsPassword{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsPassword) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsPassword) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsPassword with another TlsPassword, and returns true if they represent the same GObject.
func (recv *TlsPassword) Equals(other *TlsPassword) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *TlsPassword) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to TlsPassword.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsPassword.
func CastToTlsPassword(object *gobject.Object) *TlsPassword {
	return TlsPasswordNewFromC(object.ToC())
}

// Blacklisted : g_tls_password_new

// Blacklisted : g_tls_password_get_description

// Blacklisted : g_tls_password_get_flags

// Blacklisted : g_tls_password_get_value

// Blacklisted : g_tls_password_get_warning

// Blacklisted : g_tls_password_set_description

// Blacklisted : g_tls_password_set_flags

// Blacklisted : g_tls_password_set_value

// Unsupported : g_tls_password_set_value_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Blacklisted : g_tls_password_set_warning
