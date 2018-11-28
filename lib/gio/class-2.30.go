// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_dbus_connection_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_with_unix_fd_list_finish : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

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

// CastToWidget down casts any arbitary Object to DBusInterfaceSkeleton.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusInterfaceSkeleton.
func CastToDBusInterfaceSkeleton(object *gobject.Object) *DBusInterfaceSkeleton {
	return DBusInterfaceSkeletonNewFromC(object.ToC())
}

// Export is a wrapper around the C function g_dbus_interface_skeleton_export.
func (recv *DBusInterfaceSkeleton) Export(connection *DBusConnection, objectPath string) (bool, error) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	var cThrowableError *C.GError

	retC := C.g_dbus_interface_skeleton_export((*C.GDBusInterfaceSkeleton)(recv.native), c_connection, c_object_path, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Flush is a wrapper around the C function g_dbus_interface_skeleton_flush.
func (recv *DBusInterfaceSkeleton) Flush() {
	C.g_dbus_interface_skeleton_flush((*C.GDBusInterfaceSkeleton)(recv.native))

	return
}

// GetConnection is a wrapper around the C function g_dbus_interface_skeleton_get_connection.
func (recv *DBusInterfaceSkeleton) GetConnection() *DBusConnection {
	retC := C.g_dbus_interface_skeleton_get_connection((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_interface_skeleton_get_flags.
func (recv *DBusInterfaceSkeleton) GetFlags() DBusInterfaceSkeletonFlags {
	retC := C.g_dbus_interface_skeleton_get_flags((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := (DBusInterfaceSkeletonFlags)(retC)

	return retGo
}

// GetInfo is a wrapper around the C function g_dbus_interface_skeleton_get_info.
func (recv *DBusInterfaceSkeleton) GetInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_skeleton_get_info((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_interface_skeleton_get_object_path.
func (recv *DBusInterfaceSkeleton) GetObjectPath() string {
	retC := C.g_dbus_interface_skeleton_get_object_path((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_interface_skeleton_get_properties : return type : Blacklisted record : GVariant

// GetVtable is a wrapper around the C function g_dbus_interface_skeleton_get_vtable.
func (recv *DBusInterfaceSkeleton) GetVtable() *DBusInterfaceVTable {
	retC := C.g_dbus_interface_skeleton_get_vtable((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusInterfaceVTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFlags is a wrapper around the C function g_dbus_interface_skeleton_set_flags.
func (recv *DBusInterfaceSkeleton) SetFlags(flags DBusInterfaceSkeletonFlags) {
	c_flags := (C.GDBusInterfaceSkeletonFlags)(flags)

	C.g_dbus_interface_skeleton_set_flags((*C.GDBusInterfaceSkeleton)(recv.native), c_flags)

	return
}

// Unexport is a wrapper around the C function g_dbus_interface_skeleton_unexport.
func (recv *DBusInterfaceSkeleton) Unexport() {
	C.g_dbus_interface_skeleton_unexport((*C.GDBusInterfaceSkeleton)(recv.native))

	return
}

// Unsupported : g_dbus_method_invocation_return_value_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// TakeError is a wrapper around the C function g_dbus_method_invocation_take_error.
func (recv *DBusMethodInvocation) TakeError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_dbus_method_invocation_take_error((*C.GDBusMethodInvocation)(recv.native), c_error)

	return
}

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

// CastToWidget down casts any arbitary Object to DBusObjectManagerClient.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectManagerClient.
func CastToDBusObjectManagerClient(object *gobject.Object) *DBusObjectManagerClient {
	return DBusObjectManagerClientNewFromC(object.ToC())
}

// DBusObjectManagerClientNewFinish is a wrapper around the C function g_dbus_object_manager_client_new_finish.
func DBusObjectManagerClientNewFinish(res *AsyncResult) (*DBusObjectManagerClient, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_object_manager_client_new_finish(c_res, &cThrowableError)
	retGo := DBusObjectManagerClientNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// DBusObjectManagerClientNewForBusFinish is a wrapper around the C function g_dbus_object_manager_client_new_for_bus_finish.
func DBusObjectManagerClientNewForBusFinish(res *AsyncResult) (*DBusObjectManagerClient, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_object_manager_client_new_for_bus_finish(c_res, &cThrowableError)
	retGo := DBusObjectManagerClientNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// GetConnection is a wrapper around the C function g_dbus_object_manager_client_get_connection.
func (recv *DBusObjectManagerClient) GetConnection() *DBusConnection {
	retC := C.g_dbus_object_manager_client_get_connection((*C.GDBusObjectManagerClient)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_object_manager_client_get_flags.
func (recv *DBusObjectManagerClient) GetFlags() DBusObjectManagerClientFlags {
	retC := C.g_dbus_object_manager_client_get_flags((*C.GDBusObjectManagerClient)(recv.native))
	retGo := (DBusObjectManagerClientFlags)(retC)

	return retGo
}

// GetName is a wrapper around the C function g_dbus_object_manager_client_get_name.
func (recv *DBusObjectManagerClient) GetName() string {
	retC := C.g_dbus_object_manager_client_get_name((*C.GDBusObjectManagerClient)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNameOwner is a wrapper around the C function g_dbus_object_manager_client_get_name_owner.
func (recv *DBusObjectManagerClient) GetNameOwner() string {
	retC := C.g_dbus_object_manager_client_get_name_owner((*C.GDBusObjectManagerClient)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

// CastToWidget down casts any arbitary Object to DBusObjectManagerServer.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectManagerServer.
func CastToDBusObjectManagerServer(object *gobject.Object) *DBusObjectManagerServer {
	return DBusObjectManagerServerNewFromC(object.ToC())
}

// DBusObjectManagerServerNew is a wrapper around the C function g_dbus_object_manager_server_new.
func DBusObjectManagerServerNew(objectPath string) *DBusObjectManagerServer {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_manager_server_new(c_object_path)
	retGo := DBusObjectManagerServerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Export is a wrapper around the C function g_dbus_object_manager_server_export.
func (recv *DBusObjectManagerServer) Export(object *DBusObjectSkeleton) {
	c_object := (*C.GDBusObjectSkeleton)(C.NULL)
	if object != nil {
		c_object = (*C.GDBusObjectSkeleton)(object.ToC())
	}

	C.g_dbus_object_manager_server_export((*C.GDBusObjectManagerServer)(recv.native), c_object)

	return
}

// ExportUniquely is a wrapper around the C function g_dbus_object_manager_server_export_uniquely.
func (recv *DBusObjectManagerServer) ExportUniquely(object *DBusObjectSkeleton) {
	c_object := (*C.GDBusObjectSkeleton)(C.NULL)
	if object != nil {
		c_object = (*C.GDBusObjectSkeleton)(object.ToC())
	}

	C.g_dbus_object_manager_server_export_uniquely((*C.GDBusObjectManagerServer)(recv.native), c_object)

	return
}

// GetConnection is a wrapper around the C function g_dbus_object_manager_server_get_connection.
func (recv *DBusObjectManagerServer) GetConnection() *DBusConnection {
	retC := C.g_dbus_object_manager_server_get_connection((*C.GDBusObjectManagerServer)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetConnection is a wrapper around the C function g_dbus_object_manager_server_set_connection.
func (recv *DBusObjectManagerServer) SetConnection(connection *DBusConnection) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	C.g_dbus_object_manager_server_set_connection((*C.GDBusObjectManagerServer)(recv.native), c_connection)

	return
}

// Unexport is a wrapper around the C function g_dbus_object_manager_server_unexport.
func (recv *DBusObjectManagerServer) Unexport(objectPath string) bool {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_manager_server_unexport((*C.GDBusObjectManagerServer)(recv.native), c_object_path)
	retGo := retC == C.TRUE

	return retGo
}

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

// CastToWidget down casts any arbitary Object to DBusObjectProxy.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectProxy.
func CastToDBusObjectProxy(object *gobject.Object) *DBusObjectProxy {
	return DBusObjectProxyNewFromC(object.ToC())
}

// DBusObjectProxyNew is a wrapper around the C function g_dbus_object_proxy_new.
func DBusObjectProxyNew(connection *DBusConnection, objectPath string) *DBusObjectProxy {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_proxy_new(c_connection, c_object_path)
	retGo := DBusObjectProxyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetConnection is a wrapper around the C function g_dbus_object_proxy_get_connection.
func (recv *DBusObjectProxy) GetConnection() *DBusConnection {
	retC := C.g_dbus_object_proxy_get_connection((*C.GDBusObjectProxy)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// CastToWidget down casts any arbitary Object to DBusObjectSkeleton.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectSkeleton.
func CastToDBusObjectSkeleton(object *gobject.Object) *DBusObjectSkeleton {
	return DBusObjectSkeletonNewFromC(object.ToC())
}

// DBusObjectSkeletonNew is a wrapper around the C function g_dbus_object_skeleton_new.
func DBusObjectSkeletonNew(objectPath string) *DBusObjectSkeleton {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_skeleton_new(c_object_path)
	retGo := DBusObjectSkeletonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddInterface is a wrapper around the C function g_dbus_object_skeleton_add_interface.
func (recv *DBusObjectSkeleton) AddInterface(interface_ *DBusInterfaceSkeleton) {
	c_interface_ := (*C.GDBusInterfaceSkeleton)(C.NULL)
	if interface_ != nil {
		c_interface_ = (*C.GDBusInterfaceSkeleton)(interface_.ToC())
	}

	C.g_dbus_object_skeleton_add_interface((*C.GDBusObjectSkeleton)(recv.native), c_interface_)

	return
}

// Flush is a wrapper around the C function g_dbus_object_skeleton_flush.
func (recv *DBusObjectSkeleton) Flush() {
	C.g_dbus_object_skeleton_flush((*C.GDBusObjectSkeleton)(recv.native))

	return
}

// RemoveInterface is a wrapper around the C function g_dbus_object_skeleton_remove_interface.
func (recv *DBusObjectSkeleton) RemoveInterface(interface_ *DBusInterfaceSkeleton) {
	c_interface_ := (*C.GDBusInterfaceSkeleton)(C.NULL)
	if interface_ != nil {
		c_interface_ = (*C.GDBusInterfaceSkeleton)(interface_.ToC())
	}

	C.g_dbus_object_skeleton_remove_interface((*C.GDBusObjectSkeleton)(recv.native), c_interface_)

	return
}

// RemoveInterfaceByName is a wrapper around the C function g_dbus_object_skeleton_remove_interface_by_name.
func (recv *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	C.g_dbus_object_skeleton_remove_interface_by_name((*C.GDBusObjectSkeleton)(recv.native), c_interface_name)

	return
}

// SetObjectPath is a wrapper around the C function g_dbus_object_skeleton_set_object_path.
func (recv *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	C.g_dbus_object_skeleton_set_object_path((*C.GDBusObjectSkeleton)(recv.native), c_object_path)

	return
}

// Unsupported : g_dbus_proxy_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_finish : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

// ReadLineFinishUtf8 is a wrapper around the C function g_data_input_stream_read_line_finish_utf8.
func (recv *DataInputStream) ReadLineFinishUtf8(result *AsyncResult) (string, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_line_finish_utf8((*C.GDataInputStream)(recv.native), c_result, &c_length, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goThrowableError
}

// ReadLineUtf8 is a wrapper around the C function g_data_input_stream_read_line_utf8.
func (recv *DataInputStream) ReadLineUtf8(cancellable *Cancellable) (string, uint64, error) {
	var c_length C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_line_utf8((*C.GDataInputStream)(recv.native), &c_length, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goThrowableError
}

// GetNodisplay is a wrapper around the C function g_desktop_app_info_get_nodisplay.
func (recv *DesktopAppInfo) GetNodisplay() bool {
	retC := C.g_desktop_app_info_get_nodisplay((*C.GDesktopAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowIn is a wrapper around the C function g_desktop_app_info_get_show_in.
func (recv *DesktopAppInfo) GetShowIn(desktopEnv string) bool {
	c_desktop_env := C.CString(desktopEnv)
	defer C.free(unsafe.Pointer(c_desktop_env))

	retC := C.g_desktop_app_info_get_show_in((*C.GDesktopAppInfo)(recv.native), c_desktop_env)
	retGo := retC == C.TRUE

	return retGo
}

// Equal is a wrapper around the C function g_inet_address_equal.
func (recv *InetAddress) Equal(otherAddress *InetAddress) bool {
	c_other_address := (*C.GInetAddress)(C.NULL)
	if otherAddress != nil {
		c_other_address = (*C.GInetAddress)(otherAddress.ToC())
	}

	retC := C.g_inet_address_equal((*C.GInetAddress)(recv.native), c_other_address)
	retGo := retC == C.TRUE

	return retGo
}

// GetUint is a wrapper around the C function g_settings_get_uint.
func (recv *Settings) GetUint(key string) uint32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_uint((*C.GSettings)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

// SetUint is a wrapper around the C function g_settings_set_uint.
func (recv *Settings) SetUint(key string, value uint32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint)(value)

	retC := C.g_settings_set_uint((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported signal 'change-state' for SimpleAction : unsupported parameter value : type GLib.Variant : Blacklisted record : GVariant

// Unsupported : g_simple_action_set_state : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_group_add_entries : unsupported parameter entries :

// GetDatabase is a wrapper around the C function g_tls_connection_get_database.
func (recv *TlsConnection) GetDatabase() *TlsDatabase {
	retC := C.g_tls_connection_get_database((*C.GTlsConnection)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInteraction is a wrapper around the C function g_tls_connection_get_interaction.
func (recv *TlsConnection) GetInteraction() *TlsInteraction {
	retC := C.g_tls_connection_get_interaction((*C.GTlsConnection)(recv.native))
	retGo := TlsInteractionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDatabase is a wrapper around the C function g_tls_connection_set_database.
func (recv *TlsConnection) SetDatabase(database *TlsDatabase) {
	c_database := (*C.GTlsDatabase)(C.NULL)
	if database != nil {
		c_database = (*C.GTlsDatabase)(database.ToC())
	}

	C.g_tls_connection_set_database((*C.GTlsConnection)(recv.native), c_database)

	return
}

// SetInteraction is a wrapper around the C function g_tls_connection_set_interaction.
func (recv *TlsConnection) SetInteraction(interaction *TlsInteraction) {
	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	C.g_tls_connection_set_interaction((*C.GTlsConnection)(recv.native), c_interaction)

	return
}

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

// CastToWidget down casts any arbitary Object to TlsDatabase.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsDatabase.
func CastToTlsDatabase(object *gobject.Object) *TlsDatabase {
	return TlsDatabaseNewFromC(object.ToC())
}

// CreateCertificateHandle is a wrapper around the C function g_tls_database_create_certificate_handle.
func (recv *TlsDatabase) CreateCertificateHandle(certificate *TlsCertificate) string {
	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	retC := C.g_tls_database_create_certificate_handle((*C.GTlsDatabase)(recv.native), c_certificate)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// LookupCertificateForHandle is a wrapper around the C function g_tls_database_lookup_certificate_for_handle.
func (recv *TlsDatabase) LookupCertificateForHandle(handle string, interaction *TlsInteraction, flags TlsDatabaseLookupFlags, cancellable *Cancellable) (*TlsCertificate, error) {
	c_handle := C.CString(handle)
	defer C.free(unsafe.Pointer(c_handle))

	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	c_flags := (C.GTlsDatabaseLookupFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_database_lookup_certificate_for_handle((*C.GTlsDatabase)(recv.native), c_handle, c_interaction, c_flags, c_cancellable, &cThrowableError)
	var retGo (*TlsCertificate)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TlsCertificateNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_database_lookup_certificate_for_handle_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupCertificateForHandleFinish is a wrapper around the C function g_tls_database_lookup_certificate_for_handle_finish.
func (recv *TlsDatabase) LookupCertificateForHandleFinish(result *AsyncResult) (*TlsCertificate, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_database_lookup_certificate_for_handle_finish((*C.GTlsDatabase)(recv.native), c_result, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LookupCertificateIssuer is a wrapper around the C function g_tls_database_lookup_certificate_issuer.
func (recv *TlsDatabase) LookupCertificateIssuer(certificate *TlsCertificate, interaction *TlsInteraction, flags TlsDatabaseLookupFlags, cancellable *Cancellable) (*TlsCertificate, error) {
	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	c_flags := (C.GTlsDatabaseLookupFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_database_lookup_certificate_issuer((*C.GTlsDatabase)(recv.native), c_certificate, c_interaction, c_flags, c_cancellable, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_database_lookup_certificate_issuer_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupCertificateIssuerFinish is a wrapper around the C function g_tls_database_lookup_certificate_issuer_finish.
func (recv *TlsDatabase) LookupCertificateIssuerFinish(result *AsyncResult) (*TlsCertificate, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_database_lookup_certificate_issuer_finish((*C.GTlsDatabase)(recv.native), c_result, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LookupCertificatesIssuedBy is a wrapper around the C function g_tls_database_lookup_certificates_issued_by.
func (recv *TlsDatabase) LookupCertificatesIssuedBy(issuerRawDn []uint8, interaction *TlsInteraction, flags TlsDatabaseLookupFlags, cancellable *Cancellable) (*glib.List, error) {
	c_issuer_raw_dn := &issuerRawDn[0]

	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	c_flags := (C.GTlsDatabaseLookupFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_database_lookup_certificates_issued_by((*C.GTlsDatabase)(recv.native), (*C.GByteArray)(unsafe.Pointer(c_issuer_raw_dn)), c_interaction, c_flags, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_database_lookup_certificates_issued_by_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupCertificatesIssuedByFinish is a wrapper around the C function g_tls_database_lookup_certificates_issued_by_finish.
func (recv *TlsDatabase) LookupCertificatesIssuedByFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_database_lookup_certificates_issued_by_finish((*C.GTlsDatabase)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// VerifyChain is a wrapper around the C function g_tls_database_verify_chain.
func (recv *TlsDatabase) VerifyChain(chain *TlsCertificate, purpose string, identity *SocketConnectable, interaction *TlsInteraction, flags TlsDatabaseVerifyFlags, cancellable *Cancellable) (TlsCertificateFlags, error) {
	c_chain := (*C.GTlsCertificate)(C.NULL)
	if chain != nil {
		c_chain = (*C.GTlsCertificate)(chain.ToC())
	}

	c_purpose := C.CString(purpose)
	defer C.free(unsafe.Pointer(c_purpose))

	c_identity := (*C.GSocketConnectable)(identity.ToC())

	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	c_flags := (C.GTlsDatabaseVerifyFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_database_verify_chain((*C.GTlsDatabase)(recv.native), c_chain, c_purpose, c_identity, c_interaction, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsCertificateFlags)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_database_verify_chain_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// VerifyChainFinish is a wrapper around the C function g_tls_database_verify_chain_finish.
func (recv *TlsDatabase) VerifyChainFinish(result *AsyncResult) (TlsCertificateFlags, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_database_verify_chain_finish((*C.GTlsDatabase)(recv.native), c_result, &cThrowableError)
	retGo := (TlsCertificateFlags)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// CastToWidget down casts any arbitary Object to TlsInteraction.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsInteraction.
func CastToTlsInteraction(object *gobject.Object) *TlsInteraction {
	return TlsInteractionNewFromC(object.ToC())
}

// AskPassword is a wrapper around the C function g_tls_interaction_ask_password.
func (recv *TlsInteraction) AskPassword(password *TlsPassword, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_password := (*C.GTlsPassword)(C.NULL)
	if password != nil {
		c_password = (*C.GTlsPassword)(password.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_ask_password((*C.GTlsInteraction)(recv.native), c_password, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_interaction_ask_password_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AskPasswordFinish is a wrapper around the C function g_tls_interaction_ask_password_finish.
func (recv *TlsInteraction) AskPasswordFinish(result *AsyncResult) (TlsInteractionResult, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_ask_password_finish((*C.GTlsInteraction)(recv.native), c_result, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// InvokeAskPassword is a wrapper around the C function g_tls_interaction_invoke_ask_password.
func (recv *TlsInteraction) InvokeAskPassword(password *TlsPassword, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_password := (*C.GTlsPassword)(C.NULL)
	if password != nil {
		c_password = (*C.GTlsPassword)(password.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_invoke_ask_password((*C.GTlsInteraction)(recv.native), c_password, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// CastToWidget down casts any arbitary Object to TlsPassword.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsPassword.
func CastToTlsPassword(object *gobject.Object) *TlsPassword {
	return TlsPasswordNewFromC(object.ToC())
}

// TlsPasswordNew is a wrapper around the C function g_tls_password_new.
func TlsPasswordNew(flags TlsPasswordFlags, description string) *TlsPassword {
	c_flags := (C.GTlsPasswordFlags)(flags)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	retC := C.g_tls_password_new(c_flags, c_description)
	retGo := TlsPasswordNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDescription is a wrapper around the C function g_tls_password_get_description.
func (recv *TlsPassword) GetDescription() string {
	retC := C.g_tls_password_get_description((*C.GTlsPassword)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_tls_password_get_flags.
func (recv *TlsPassword) GetFlags() TlsPasswordFlags {
	retC := C.g_tls_password_get_flags((*C.GTlsPassword)(recv.native))
	retGo := (TlsPasswordFlags)(retC)

	return retGo
}

// Blacklisted : g_tls_password_get_value

// GetWarning is a wrapper around the C function g_tls_password_get_warning.
func (recv *TlsPassword) GetWarning() string {
	retC := C.g_tls_password_get_warning((*C.GTlsPassword)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetDescription is a wrapper around the C function g_tls_password_set_description.
func (recv *TlsPassword) SetDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_tls_password_set_description((*C.GTlsPassword)(recv.native), c_description)

	return
}

// SetFlags is a wrapper around the C function g_tls_password_set_flags.
func (recv *TlsPassword) SetFlags(flags TlsPasswordFlags) {
	c_flags := (C.GTlsPasswordFlags)(flags)

	C.g_tls_password_set_flags((*C.GTlsPassword)(recv.native), c_flags)

	return
}

// SetValue is a wrapper around the C function g_tls_password_set_value.
func (recv *TlsPassword) SetValue(value []uint8) {
	c_value := &value[0]

	c_length := (C.gssize)(len(value))

	C.g_tls_password_set_value((*C.GTlsPassword)(recv.native), (*C.guchar)(unsafe.Pointer(c_value)), c_length)

	return
}

// Unsupported : g_tls_password_set_value_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// SetWarning is a wrapper around the C function g_tls_password_set_warning.
func (recv *TlsPassword) SetWarning(warning string) {
	c_warning := C.CString(warning)
	defer C.free(unsafe.Pointer(c_warning))

	C.g_tls_password_set_warning((*C.GTlsPassword)(recv.native), c_warning)

	return
}
