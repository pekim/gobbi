// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
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
import "C"

// Credentials is a wrapper around the C record GCredentials.
type Credentials struct {
	native *C.GCredentials
}

func CredentialsNewFromC(u unsafe.Pointer) *Credentials {
	c := (*C.GCredentials)(u)
	if c == nil {
		return nil
	}

	g := &Credentials{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Credentials) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Credentials) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CredentialsNew is a wrapper around the C function g_credentials_new.
func CredentialsNew() *Credentials {
	retC := C.g_credentials_new()
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusAuthObserver is a wrapper around the C record GDBusAuthObserver.
type DBusAuthObserver struct {
	native *C.GDBusAuthObserver
}

func DBusAuthObserverNewFromC(u unsafe.Pointer) *DBusAuthObserver {
	c := (*C.GDBusAuthObserver)(u)
	if c == nil {
		return nil
	}

	g := &DBusAuthObserver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusAuthObserver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusAuthObserver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusAuthObserverNew is a wrapper around the C function g_dbus_auth_observer_new.
func DBusAuthObserverNew() *DBusAuthObserver {
	retC := C.g_dbus_auth_observer_new()
	retGo := DBusAuthObserverNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusConnection is a wrapper around the C record GDBusConnection.
type DBusConnection struct {
	native *C.GDBusConnection
}

func DBusConnectionNewFromC(u unsafe.Pointer) *DBusConnection {
	c := (*C.GDBusConnection)(u)
	if c == nil {
		return nil
	}

	g := &DBusConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusConnectionNewFinish is a wrapper around the C function g_dbus_connection_new_finish.
func DBusConnectionNewFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusConnectionNewForAddressFinish is a wrapper around the C function g_dbus_connection_new_for_address_finish.
func DBusConnectionNewForAddressFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusConnectionNewForAddressSync is a wrapper around the C function g_dbus_connection_new_for_address_sync.
func DBusConnectionNewForAddressSync(address string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_sync(c_address, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusConnectionNewSync is a wrapper around the C function g_dbus_connection_new_sync.
func DBusConnectionNewSync(stream *IOStream, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_sync(c_stream, c_guid, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusMessage is a wrapper around the C record GDBusMessage.
type DBusMessage struct {
	native *C.GDBusMessage
}

func DBusMessageNewFromC(u unsafe.Pointer) *DBusMessage {
	c := (*C.GDBusMessage)(u)
	if c == nil {
		return nil
	}

	g := &DBusMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMessageNew is a wrapper around the C function g_dbus_message_new.
func DBusMessageNew() *DBusMessage {
	retC := C.g_dbus_message_new()
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusMessageNewFromBlob is a wrapper around the C function g_dbus_message_new_from_blob.
func DBusMessageNewFromBlob(blob []uint8, capabilities DBusCapabilityFlags) (*DBusMessage, error) {
	c_blob_array := make([]C.guint8, len(blob)+1, len(blob)+1)
	for i, item := range blob {
		c := (C.guint8)(item)
		c_blob_array[i] = c
	}
	c_blob_array[len(blob)] = 0
	c_blob_arrayPtr := &c_blob_array[0]
	c_blob := (*C.guchar)(unsafe.Pointer(c_blob_arrayPtr))

	c_blob_len := (C.gsize)(len(blob))

	c_capabilities := (C.GDBusCapabilityFlags)(capabilities)

	var cThrowableError *C.GError

	retC := C.g_dbus_message_new_from_blob(c_blob, c_blob_len, c_capabilities, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusMessageNewMethodCall is a wrapper around the C function g_dbus_message_new_method_call.
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) *DBusMessage {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	retC := C.g_dbus_message_new_method_call(c_name, c_path, c_interface_, c_method)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusMessageNewSignal is a wrapper around the C function g_dbus_message_new_signal.
func DBusMessageNewSignal(path string, interface_ string, signal string) *DBusMessage {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_signal := C.CString(signal)
	defer C.free(unsafe.Pointer(c_signal))

	retC := C.g_dbus_message_new_signal(c_path, c_interface_, c_signal)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusMethodInvocation is a wrapper around the C record GDBusMethodInvocation.
type DBusMethodInvocation struct {
	native *C.GDBusMethodInvocation
}

func DBusMethodInvocationNewFromC(u unsafe.Pointer) *DBusMethodInvocation {
	c := (*C.GDBusMethodInvocation)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInvocation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMethodInvocation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusProxy is a wrapper around the C record GDBusProxy.
type DBusProxy struct {
	native *C.GDBusProxy
	// Private : parent_instance
	// Private : priv
}

func DBusProxyNewFromC(u unsafe.Pointer) *DBusProxy {
	c := (*C.GDBusProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxy{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusProxy) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusProxyNewFinish is a wrapper around the C function g_dbus_proxy_new_finish.
func DBusProxyNewFinish(res *AsyncResult) (*DBusProxy, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_finish(c_res, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusProxyNewForBusFinish is a wrapper around the C function g_dbus_proxy_new_for_bus_finish.
func DBusProxyNewForBusFinish(res *AsyncResult) (*DBusProxy, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_finish(c_res, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusProxyNewForBusSync is a wrapper around the C function g_dbus_proxy_new_for_bus_sync.
func DBusProxyNewForBusSync(busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_bus_type := (C.GBusType)(busType)

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_sync(c_bus_type, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusProxyNewSync is a wrapper around the C function g_dbus_proxy_new_sync.
func DBusProxyNewSync(connection *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_sync(c_connection, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusServer is a wrapper around the C record GDBusServer.
type DBusServer struct {
	native *C.GDBusServer
}

func DBusServerNewFromC(u unsafe.Pointer) *DBusServer {
	c := (*C.GDBusServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusServer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusServer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusServerNewSync is a wrapper around the C function g_dbus_server_new_sync.
func DBusServerNewSync(address string, flags DBusServerFlags, guid string, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusServer, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusServerFlags)(flags)

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_server_new_sync(c_address, c_flags, c_guid, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusServerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native *C.GProxyAddress
	// parent_instance : record
	// Private : priv
}

func ProxyAddressNewFromC(u unsafe.Pointer) *ProxyAddress {
	c := (*C.GProxyAddress)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProxyAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProxyAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressNew is a wrapper around the C function g_proxy_address_new.
func ProxyAddressNew(inetaddr *InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) *ProxyAddress {
	c_inetaddr := (*C.GInetAddress)(C.NULL)
	if inetaddr != nil {
		c_inetaddr = (*C.GInetAddress)(inetaddr.ToC())
	}

	c_port := (C.guint16)(port)

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_dest_hostname := C.CString(destHostname)
	defer C.free(unsafe.Pointer(c_dest_hostname))

	c_dest_port := (C.guint16)(destPort)

	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	retC := C.g_proxy_address_new(c_inetaddr, c_port, c_protocol, c_dest_hostname, c_dest_port, c_username, c_password)
	retGo := ProxyAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNew is a wrapper around the C function g_settings_new.
func SettingsNew(schemaId string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	retC := C.g_settings_new(c_schema_id)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNewWithBackend is a wrapper around the C function g_settings_new_with_backend.
func SettingsNewWithBackend(schemaId string, backend *SettingsBackend) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	retC := C.g_settings_new_with_backend(c_schema_id, c_backend)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNewWithBackendAndPath is a wrapper around the C function g_settings_new_with_backend_and_path.
func SettingsNewWithBackendAndPath(schemaId string, backend *SettingsBackend, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_backend_and_path(c_schema_id, c_backend, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNewWithPath is a wrapper around the C function g_settings_new_with_path.
func SettingsNewWithPath(schemaId string, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_path(c_schema_id, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SimplePermissionNew is a wrapper around the C function g_simple_permission_new.
func SimplePermissionNew(allowed bool) *SimplePermission {
	c_allowed :=
		boolToGboolean(allowed)

	retC := C.g_simple_permission_new(c_allowed)
	retGo := SimplePermissionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native *C.GUnixCredentialsMessage
	// parent_instance : record
	// priv : record
}

func UnixCredentialsMessageNewFromC(u unsafe.Pointer) *UnixCredentialsMessage {
	c := (*C.GUnixCredentialsMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixCredentialsMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessageNew is a wrapper around the C function g_unix_credentials_message_new.
func UnixCredentialsMessageNew() *UnixCredentialsMessage {
	retC := C.g_unix_credentials_message_new()
	retGo := UnixCredentialsMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixCredentialsMessageNewWithCredentials is a wrapper around the C function g_unix_credentials_message_new_with_credentials.
func UnixCredentialsMessageNewWithCredentials(credentials *Credentials) *UnixCredentialsMessage {
	c_credentials := (*C.GCredentials)(C.NULL)
	if credentials != nil {
		c_credentials = (*C.GCredentials)(credentials.ToC())
	}

	retC := C.g_unix_credentials_message_new_with_credentials(c_credentials)
	retGo := UnixCredentialsMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixSocketAddressNewWithType is a wrapper around the C function g_unix_socket_address_new_with_type.
func UnixSocketAddressNewWithType(path []rune, type_ UnixSocketAddressType) *UnixSocketAddress {
	c_path_array := make([]C.gchar, len(path)+1, len(path)+1)
	for i, item := range path {
		c := (C.gchar)(item)
		c_path_array[i] = c
	}
	c_path_array[len(path)] = 0
	c_path_arrayPtr := &c_path_array[0]
	c_path := (*C.gchar)(unsafe.Pointer(c_path_arrayPtr))

	c_path_len := (C.gint)(len(path))

	c_type := (C.GUnixSocketAddressType)(type_)

	retC := C.g_unix_socket_address_new_with_type(c_path, c_path_len, c_type)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_bus_get : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback (GBusAcquiredCallback) for param bus_acquired_handler

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback (GBusNameAcquiredCallback) for param name_acquired_handler

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations :

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter entries :

// Proxy is a wrapper around the C record GProxy.
type Proxy struct {
	native *C.GProxy
}

func ProxyNewFromC(u unsafe.Pointer) *Proxy {
	c := (*C.GProxy)(u)
	if c == nil {
		return nil
	}

	g := &Proxy{native: c}

	return g
}

func (recv *Proxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyResolver is a wrapper around the C record GProxyResolver.
type ProxyResolver struct {
	native *C.GProxyResolver
}

func ProxyResolverNewFromC(u unsafe.Pointer) *ProxyResolver {
	c := (*C.GProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolver{native: c}

	return g
}

func (recv *ProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CredentialsClass is a wrapper around the C record GCredentialsClass.
type CredentialsClass struct {
	native *C.GCredentialsClass
}

func CredentialsClassNewFromC(u unsafe.Pointer) *CredentialsClass {
	c := (*C.GCredentialsClass)(u)
	if c == nil {
		return nil
	}

	g := &CredentialsClass{native: c}

	return g
}

func (recv *CredentialsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusAnnotationInfo is a wrapper around the C record GDBusAnnotationInfo.
type DBusAnnotationInfo struct {
	native   *C.GDBusAnnotationInfo
	RefCount int32
	Key      string
	Value    string
	// no type for annotations
}

func DBusAnnotationInfoNewFromC(u unsafe.Pointer) *DBusAnnotationInfo {
	c := (*C.GDBusAnnotationInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusAnnotationInfo{
		Key:      C.GoString(c.key),
		RefCount: (int32)(c.ref_count),
		Value:    C.GoString(c.value),
		native:   c,
	}

	return g
}

func (recv *DBusAnnotationInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// DBusArgInfo is a wrapper around the C record GDBusArgInfo.
type DBusArgInfo struct {
	native    *C.GDBusArgInfo
	RefCount  int32
	Name      string
	Signature string
	// no type for annotations
}

func DBusArgInfoNewFromC(u unsafe.Pointer) *DBusArgInfo {
	c := (*C.GDBusArgInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusArgInfo{
		Name:      C.GoString(c.name),
		RefCount:  (int32)(c.ref_count),
		Signature: C.GoString(c.signature),
		native:    c,
	}

	return g
}

func (recv *DBusArgInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.signature =
		C.CString(recv.Signature)

	return (unsafe.Pointer)(recv.native)
}

// DBusErrorEntry is a wrapper around the C record GDBusErrorEntry.
type DBusErrorEntry struct {
	native        *C.GDBusErrorEntry
	ErrorCode     int32
	DbusErrorName string
}

func DBusErrorEntryNewFromC(u unsafe.Pointer) *DBusErrorEntry {
	c := (*C.GDBusErrorEntry)(u)
	if c == nil {
		return nil
	}

	g := &DBusErrorEntry{
		DbusErrorName: C.GoString(c.dbus_error_name),
		ErrorCode:     (int32)(c.error_code),
		native:        c,
	}

	return g
}

func (recv *DBusErrorEntry) ToC() unsafe.Pointer {
	recv.native.error_code =
		(C.gint)(recv.ErrorCode)
	recv.native.dbus_error_name =
		C.CString(recv.DbusErrorName)

	return (unsafe.Pointer)(recv.native)
}

// DBusInterfaceInfo is a wrapper around the C record GDBusInterfaceInfo.
type DBusInterfaceInfo struct {
	native   *C.GDBusInterfaceInfo
	RefCount int32
	Name     string
	// no type for methods
	// no type for signals
	// no type for properties
	// no type for annotations
}

func DBusInterfaceInfoNewFromC(u unsafe.Pointer) *DBusInterfaceInfo {
	c := (*C.GDBusInterfaceInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceInfo{
		Name:     C.GoString(c.name),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusInterfaceInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// DBusInterfaceVTable is a wrapper around the C record GDBusInterfaceVTable.
type DBusInterfaceVTable struct {
	native *C.GDBusInterfaceVTable
	// method_call : no type generator for DBusInterfaceMethodCallFunc, GDBusInterfaceMethodCallFunc
	// get_property : no type generator for DBusInterfaceGetPropertyFunc, GDBusInterfaceGetPropertyFunc
	// set_property : no type generator for DBusInterfaceSetPropertyFunc, GDBusInterfaceSetPropertyFunc
	// Private : padding
}

func DBusInterfaceVTableNewFromC(u unsafe.Pointer) *DBusInterfaceVTable {
	c := (*C.GDBusInterfaceVTable)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceVTable{native: c}

	return g
}

func (recv *DBusInterfaceVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMethodInfo is a wrapper around the C record GDBusMethodInfo.
type DBusMethodInfo struct {
	native   *C.GDBusMethodInfo
	RefCount int32
	Name     string
	// no type for in_args
	// no type for out_args
	// no type for annotations
}

func DBusMethodInfoNewFromC(u unsafe.Pointer) *DBusMethodInfo {
	c := (*C.GDBusMethodInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInfo{
		Name:     C.GoString(c.name),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusMethodInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// DBusNodeInfo is a wrapper around the C record GDBusNodeInfo.
type DBusNodeInfo struct {
	native   *C.GDBusNodeInfo
	RefCount int32
	Path     string
	// no type for interfaces
	// no type for nodes
	// no type for annotations
}

func DBusNodeInfoNewFromC(u unsafe.Pointer) *DBusNodeInfo {
	c := (*C.GDBusNodeInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusNodeInfo{
		Path:     C.GoString(c.path),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusNodeInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.path =
		C.CString(recv.Path)

	return (unsafe.Pointer)(recv.native)
}

// DBusNodeInfoNewForXml is a wrapper around the C function g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) (*DBusNodeInfo, error) {
	c_xml_data := C.CString(xmlData)
	defer C.free(unsafe.Pointer(c_xml_data))

	var cThrowableError *C.GError

	retC := C.g_dbus_node_info_new_for_xml(c_xml_data, &cThrowableError)
	retGo := DBusNodeInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusPropertyInfo is a wrapper around the C record GDBusPropertyInfo.
type DBusPropertyInfo struct {
	native    *C.GDBusPropertyInfo
	RefCount  int32
	Name      string
	Signature string
	Flags     DBusPropertyInfoFlags
	// no type for annotations
}

func DBusPropertyInfoNewFromC(u unsafe.Pointer) *DBusPropertyInfo {
	c := (*C.GDBusPropertyInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusPropertyInfo{
		Flags:     (DBusPropertyInfoFlags)(c.flags),
		Name:      C.GoString(c.name),
		RefCount:  (int32)(c.ref_count),
		Signature: C.GoString(c.signature),
		native:    c,
	}

	return g
}

func (recv *DBusPropertyInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.signature =
		C.CString(recv.Signature)
	recv.native.flags =
		(C.GDBusPropertyInfoFlags)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// DBusProxyClass is a wrapper around the C record GDBusProxyClass.
type DBusProxyClass struct {
	native *C.GDBusProxyClass
	// Private : parent_class
	// no type for g_properties_changed
	// no type for g_signal
	// Private : padding
}

func DBusProxyClassNewFromC(u unsafe.Pointer) *DBusProxyClass {
	c := (*C.GDBusProxyClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxyClass{native: c}

	return g
}

func (recv *DBusProxyClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusSignalInfo is a wrapper around the C record GDBusSignalInfo.
type DBusSignalInfo struct {
	native   *C.GDBusSignalInfo
	RefCount int32
	Name     string
	// no type for args
	// no type for annotations
}

func DBusSignalInfoNewFromC(u unsafe.Pointer) *DBusSignalInfo {
	c := (*C.GDBusSignalInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusSignalInfo{
		Name:     C.GoString(c.name),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusSignalInfo) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// DBusSubtreeVTable is a wrapper around the C record GDBusSubtreeVTable.
type DBusSubtreeVTable struct {
	native *C.GDBusSubtreeVTable
	// enumerate : no type generator for DBusSubtreeEnumerateFunc, GDBusSubtreeEnumerateFunc
	// introspect : no type generator for DBusSubtreeIntrospectFunc, GDBusSubtreeIntrospectFunc
	// dispatch : no type generator for DBusSubtreeDispatchFunc, GDBusSubtreeDispatchFunc
	// Private : padding
}

func DBusSubtreeVTableNewFromC(u unsafe.Pointer) *DBusSubtreeVTable {
	c := (*C.GDBusSubtreeVTable)(u)
	if c == nil {
		return nil
	}

	g := &DBusSubtreeVTable{native: c}

	return g
}

func (recv *DBusSubtreeVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressClass is a wrapper around the C record GProxyAddressClass.
type ProxyAddressClass struct {
	native *C.GProxyAddressClass
	// parent_class : record
}

func ProxyAddressClassNewFromC(u unsafe.Pointer) *ProxyAddressClass {
	c := (*C.GProxyAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressClass{native: c}

	return g
}

func (recv *ProxyAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyInterface is a wrapper around the C record GProxyInterface.
type ProxyInterface struct {
	native *C.GProxyInterface
	// g_iface : record
	// no type for connect
	// no type for connect_async
	// no type for connect_finish
	// no type for supports_hostname
}

func ProxyInterfaceNewFromC(u unsafe.Pointer) *ProxyInterface {
	c := (*C.GProxyInterface)(u)
	if c == nil {
		return nil
	}

	g := &ProxyInterface{native: c}

	return g
}

func (recv *ProxyInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsClientConnectionInterface is a wrapper around the C record GTlsClientConnectionInterface.
type TlsClientConnectionInterface struct {
	native *C.GTlsClientConnectionInterface
	// g_iface : record
	// no type for copy_session_state
}

func TlsClientConnectionInterfaceNewFromC(u unsafe.Pointer) *TlsClientConnectionInterface {
	c := (*C.GTlsClientConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsClientConnectionInterface{native: c}

	return g
}

func (recv *TlsClientConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsServerConnectionInterface is a wrapper around the C record GTlsServerConnectionInterface.
type TlsServerConnectionInterface struct {
	native *C.GTlsServerConnectionInterface
	// g_iface : record
}

func TlsServerConnectionInterfaceNewFromC(u unsafe.Pointer) *TlsServerConnectionInterface {
	c := (*C.GTlsServerConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsServerConnectionInterface{native: c}

	return g
}

func (recv *TlsServerConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessageClass is a wrapper around the C record GUnixCredentialsMessageClass.
type UnixCredentialsMessageClass struct {
	native *C.GUnixCredentialsMessageClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func UnixCredentialsMessageClassNewFromC(u unsafe.Pointer) *UnixCredentialsMessageClass {
	c := (*C.GUnixCredentialsMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessageClass{native: c}

	return g
}

func (recv *UnixCredentialsMessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
