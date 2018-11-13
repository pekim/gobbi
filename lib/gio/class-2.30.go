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

// Object upcasts to *Object
func (recv *DBusInterfaceSkeleton) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusInterfaceSkeleton.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusInterfaceSkeleton.
func CastToDBusInterfaceSkeleton(object *gobject.Object) *DBusInterfaceSkeleton {
	return DBusInterfaceSkeletonNewFromC(object.ToC())
}

// Exports @interface_ at @object_path on @connection.
//
// This can be called multiple times to export the same @interface_
// onto multiple connections however the @object_path provided must be
// the same for all connections.
//
// Use g_dbus_interface_skeleton_unexport() to unexport the object.
/*

C function : g_dbus_interface_skeleton_export
*/
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

// If @interface_ has outstanding changes, request for these changes to be
// emitted immediately.
//
// For example, an exported D-Bus interface may queue up property
// changes and emit the
// `org.freedesktop.DBus.Properties.PropertiesChanged`
// signal later (e.g. in an idle handler). This technique is useful
// for collapsing multiple property changes into one.
/*

C function : g_dbus_interface_skeleton_flush
*/
func (recv *DBusInterfaceSkeleton) Flush() {
	C.g_dbus_interface_skeleton_flush((*C.GDBusInterfaceSkeleton)(recv.native))

	return
}

// Gets the first connection that @interface_ is exported on, if any.
/*

C function : g_dbus_interface_skeleton_get_connection
*/
func (recv *DBusInterfaceSkeleton) GetConnection() *DBusConnection {
	retC := C.g_dbus_interface_skeleton_get_connection((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GDBusInterfaceSkeletonFlags that describes what the behavior
// of @interface_
/*

C function : g_dbus_interface_skeleton_get_flags
*/
func (recv *DBusInterfaceSkeleton) GetFlags() DBusInterfaceSkeletonFlags {
	retC := C.g_dbus_interface_skeleton_get_flags((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := (DBusInterfaceSkeletonFlags)(retC)

	return retGo
}

// Gets D-Bus introspection information for the D-Bus interface
// implemented by @interface_.
/*

C function : g_dbus_interface_skeleton_get_info
*/
func (recv *DBusInterfaceSkeleton) GetInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_skeleton_get_info((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the object path that @interface_ is exported on, if any.
/*

C function : g_dbus_interface_skeleton_get_object_path
*/
func (recv *DBusInterfaceSkeleton) GetObjectPath() string {
	retC := C.g_dbus_interface_skeleton_get_object_path((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_interface_skeleton_get_properties : return type : Blacklisted record : GVariant

// Gets the interface vtable for the D-Bus interface implemented by
// @interface_. The returned function pointers should expect @interface_
// itself to be passed as @user_data.
/*

C function : g_dbus_interface_skeleton_get_vtable
*/
func (recv *DBusInterfaceSkeleton) GetVtable() *DBusInterfaceVTable {
	retC := C.g_dbus_interface_skeleton_get_vtable((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusInterfaceVTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets flags describing what the behavior of @skeleton should be.
/*

C function : g_dbus_interface_skeleton_set_flags
*/
func (recv *DBusInterfaceSkeleton) SetFlags(flags DBusInterfaceSkeletonFlags) {
	c_flags := (C.GDBusInterfaceSkeletonFlags)(flags)

	C.g_dbus_interface_skeleton_set_flags((*C.GDBusInterfaceSkeleton)(recv.native), c_flags)

	return
}

// Stops exporting @interface_ on all connections it is exported on.
//
// To unexport @interface_ from only a single connection, use
// g_dbus_interface_skeleton_unexport_from_connection()
/*

C function : g_dbus_interface_skeleton_unexport
*/
func (recv *DBusInterfaceSkeleton) Unexport() {
	C.g_dbus_interface_skeleton_unexport((*C.GDBusInterfaceSkeleton)(recv.native))

	return
}

// Unsupported : g_dbus_method_invocation_return_value_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Like g_dbus_method_invocation_return_gerror() but takes ownership
// of @error so the caller does not need to free it.
//
// This method will take ownership of @invocation. See
// #GDBusInterfaceVTable for more information about the ownership of
// @invocation.
/*

C function : g_dbus_method_invocation_take_error
*/
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

// Object upcasts to *Object
func (recv *DBusObjectManagerClient) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusObjectManagerClient.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectManagerClient.
func CastToDBusObjectManagerClient(object *gobject.Object) *DBusObjectManagerClient {
	return DBusObjectManagerClientNewFromC(object.ToC())
}

// Finishes an operation started with g_dbus_object_manager_client_new().
/*

C function : g_dbus_object_manager_client_new_finish
*/
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

// Finishes an operation started with g_dbus_object_manager_client_new_for_bus().
/*

C function : g_dbus_object_manager_client_new_for_bus_finish
*/
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

// Gets the #GDBusConnection used by @manager.
/*

C function : g_dbus_object_manager_client_get_connection
*/
func (recv *DBusObjectManagerClient) GetConnection() *DBusConnection {
	retC := C.g_dbus_object_manager_client_get_connection((*C.GDBusObjectManagerClient)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the flags that @manager was constructed with.
/*

C function : g_dbus_object_manager_client_get_flags
*/
func (recv *DBusObjectManagerClient) GetFlags() DBusObjectManagerClientFlags {
	retC := C.g_dbus_object_manager_client_get_flags((*C.GDBusObjectManagerClient)(recv.native))
	retGo := (DBusObjectManagerClientFlags)(retC)

	return retGo
}

// Gets the name that @manager is for, or %NULL if not a message bus
// connection.
/*

C function : g_dbus_object_manager_client_get_name
*/
func (recv *DBusObjectManagerClient) GetName() string {
	retC := C.g_dbus_object_manager_client_get_name((*C.GDBusObjectManagerClient)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// The unique name that owns the name that @manager is for or %NULL if
// no-one currently owns that name. You can connect to the
// #GObject::notify signal to track changes to the
// #GDBusObjectManagerClient:name-owner property.
/*

C function : g_dbus_object_manager_client_get_name_owner
*/
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

// Object upcasts to *Object
func (recv *DBusObjectManagerServer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusObjectManagerServer.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectManagerServer.
func CastToDBusObjectManagerServer(object *gobject.Object) *DBusObjectManagerServer {
	return DBusObjectManagerServerNewFromC(object.ToC())
}

// Creates a new #GDBusObjectManagerServer object.
//
// The returned server isn't yet exported on any connection. To do so,
// use g_dbus_object_manager_server_set_connection(). Normally you
// want to export all of your objects before doing so to avoid
// [InterfacesAdded](http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// signals being emitted.
/*

C function : g_dbus_object_manager_server_new
*/
func DBusObjectManagerServerNew(objectPath string) *DBusObjectManagerServer {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_manager_server_new(c_object_path)
	retGo := DBusObjectManagerServerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Exports @object on @manager.
//
// If there is already a #GDBusObject exported at the object path,
// then the old object is removed.
//
// The object path for @object must be in the hierarchy rooted by the
// object path for @manager.
//
// Note that @manager will take a reference on @object for as long as
// it is exported.
/*

C function : g_dbus_object_manager_server_export
*/
func (recv *DBusObjectManagerServer) Export(object *DBusObjectSkeleton) {
	c_object := (*C.GDBusObjectSkeleton)(C.NULL)
	if object != nil {
		c_object = (*C.GDBusObjectSkeleton)(object.ToC())
	}

	C.g_dbus_object_manager_server_export((*C.GDBusObjectManagerServer)(recv.native), c_object)

	return
}

// Like g_dbus_object_manager_server_export() but appends a string of
// the form _N (with N being a natural number) to @object's object path
// if an object with the given path already exists. As such, the
// #GDBusObjectProxy:g-object-path property of @object may be modified.
/*

C function : g_dbus_object_manager_server_export_uniquely
*/
func (recv *DBusObjectManagerServer) ExportUniquely(object *DBusObjectSkeleton) {
	c_object := (*C.GDBusObjectSkeleton)(C.NULL)
	if object != nil {
		c_object = (*C.GDBusObjectSkeleton)(object.ToC())
	}

	C.g_dbus_object_manager_server_export_uniquely((*C.GDBusObjectManagerServer)(recv.native), c_object)

	return
}

// Gets the #GDBusConnection used by @manager.
/*

C function : g_dbus_object_manager_server_get_connection
*/
func (recv *DBusObjectManagerServer) GetConnection() *DBusConnection {
	retC := C.g_dbus_object_manager_server_get_connection((*C.GDBusObjectManagerServer)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Exports all objects managed by @manager on @connection. If
// @connection is %NULL, stops exporting objects.
/*

C function : g_dbus_object_manager_server_set_connection
*/
func (recv *DBusObjectManagerServer) SetConnection(connection *DBusConnection) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	C.g_dbus_object_manager_server_set_connection((*C.GDBusObjectManagerServer)(recv.native), c_connection)

	return
}

// If @manager has an object at @path, removes the object. Otherwise
// does nothing.
//
// Note that @object_path must be in the hierarchy rooted by the
// object path for @manager.
/*

C function : g_dbus_object_manager_server_unexport
*/
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

// Object upcasts to *Object
func (recv *DBusObjectProxy) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusObjectProxy.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectProxy.
func CastToDBusObjectProxy(object *gobject.Object) *DBusObjectProxy {
	return DBusObjectProxyNewFromC(object.ToC())
}

// Creates a new #GDBusObjectProxy for the given connection and
// object path.
/*

C function : g_dbus_object_proxy_new
*/
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

// Gets the connection that @proxy is for.
/*

C function : g_dbus_object_proxy_get_connection
*/
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

// Object upcasts to *Object
func (recv *DBusObjectSkeleton) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusObjectSkeleton.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusObjectSkeleton.
func CastToDBusObjectSkeleton(object *gobject.Object) *DBusObjectSkeleton {
	return DBusObjectSkeletonNewFromC(object.ToC())
}

// Creates a new #GDBusObjectSkeleton.
/*

C function : g_dbus_object_skeleton_new
*/
func DBusObjectSkeletonNew(objectPath string) *DBusObjectSkeleton {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_skeleton_new(c_object_path)
	retGo := DBusObjectSkeletonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds @interface_ to @object.
//
// If @object already contains a #GDBusInterfaceSkeleton with the same
// interface name, it is removed before @interface_ is added.
//
// Note that @object takes its own reference on @interface_ and holds
// it until removed.
/*

C function : g_dbus_object_skeleton_add_interface
*/
func (recv *DBusObjectSkeleton) AddInterface(interface_ *DBusInterfaceSkeleton) {
	c_interface_ := (*C.GDBusInterfaceSkeleton)(C.NULL)
	if interface_ != nil {
		c_interface_ = (*C.GDBusInterfaceSkeleton)(interface_.ToC())
	}

	C.g_dbus_object_skeleton_add_interface((*C.GDBusObjectSkeleton)(recv.native), c_interface_)

	return
}

// This method simply calls g_dbus_interface_skeleton_flush() on all
// interfaces belonging to @object. See that method for when flushing
// is useful.
/*

C function : g_dbus_object_skeleton_flush
*/
func (recv *DBusObjectSkeleton) Flush() {
	C.g_dbus_object_skeleton_flush((*C.GDBusObjectSkeleton)(recv.native))

	return
}

// Removes @interface_ from @object.
/*

C function : g_dbus_object_skeleton_remove_interface
*/
func (recv *DBusObjectSkeleton) RemoveInterface(interface_ *DBusInterfaceSkeleton) {
	c_interface_ := (*C.GDBusInterfaceSkeleton)(C.NULL)
	if interface_ != nil {
		c_interface_ = (*C.GDBusInterfaceSkeleton)(interface_.ToC())
	}

	C.g_dbus_object_skeleton_remove_interface((*C.GDBusObjectSkeleton)(recv.native), c_interface_)

	return
}

// Removes the #GDBusInterface with @interface_name from @object.
//
// If no D-Bus interface of the given interface exists, this function
// does nothing.
/*

C function : g_dbus_object_skeleton_remove_interface_by_name
*/
func (recv *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	C.g_dbus_object_skeleton_remove_interface_by_name((*C.GDBusObjectSkeleton)(recv.native), c_interface_name)

	return
}

// Sets the object path for @object.
/*

C function : g_dbus_object_skeleton_set_object_path
*/
func (recv *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	C.g_dbus_object_skeleton_set_object_path((*C.GDBusObjectSkeleton)(recv.native), c_object_path)

	return
}

// Unsupported : g_dbus_proxy_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_finish : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Finish an asynchronous call started by
// g_data_input_stream_read_line_async().
/*

C function : g_data_input_stream_read_line_finish_utf8
*/
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

// Reads a UTF-8 encoded line from the data input stream.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_data_input_stream_read_line_utf8
*/
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

// Gets the value of the NoDisplay key, which helps determine if the
// application info should be shown in menus. See
// #G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
/*

C function : g_desktop_app_info_get_nodisplay
*/
func (recv *DesktopAppInfo) GetNodisplay() bool {
	retC := C.g_desktop_app_info_get_nodisplay((*C.GDesktopAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the application info should be shown in menus that list available
// applications for a specific name of the desktop, based on the
// `OnlyShowIn` and `NotShowIn` keys.
//
// @desktop_env should typically be given as %NULL, in which case the
// `XDG_CURRENT_DESKTOP` environment variable is consulted.  If you want
// to override the default mechanism then you may specify @desktop_env,
// but this is not recommended.
//
// Note that g_app_info_should_show() for @info will include this check (with
// %NULL for @desktop_env) as well as additional checks.
/*

C function : g_desktop_app_info_get_show_in
*/
func (recv *DesktopAppInfo) GetShowIn(desktopEnv string) bool {
	c_desktop_env := C.CString(desktopEnv)
	defer C.free(unsafe.Pointer(c_desktop_env))

	retC := C.g_desktop_app_info_get_show_in((*C.GDesktopAppInfo)(recv.native), c_desktop_env)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if two #GInetAddress instances are equal, e.g. the same address.
/*

C function : g_inet_address_equal
*/
func (recv *InetAddress) Equal(otherAddress *InetAddress) bool {
	c_other_address := (*C.GInetAddress)(C.NULL)
	if otherAddress != nil {
		c_other_address = (*C.GInetAddress)(otherAddress.ToC())
	}

	retC := C.g_inet_address_equal((*C.GInetAddress)(recv.native), c_other_address)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for 32-bit unsigned
// integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a uint32 type in the schema for @settings.
/*

C function : g_settings_get_uint
*/
func (recv *Settings) GetUint(key string) uint32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_uint((*C.GSettings)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for 32-bit unsigned
// integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a uint32 type in the schema for @settings.
/*

C function : g_settings_set_uint
*/
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

// Gets the certificate database that @conn uses to verify
// peer certificates. See g_tls_connection_set_database().
/*

C function : g_tls_connection_get_database
*/
func (recv *TlsConnection) GetDatabase() *TlsDatabase {
	retC := C.g_tls_connection_get_database((*C.GTlsConnection)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the object that will be used to interact with the user. It will be used
// for things like prompting the user for passwords. If %NULL is returned, then
// no user interaction will occur for this connection.
/*

C function : g_tls_connection_get_interaction
*/
func (recv *TlsConnection) GetInteraction() *TlsInteraction {
	retC := C.g_tls_connection_get_interaction((*C.GTlsConnection)(recv.native))
	retGo := TlsInteractionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the certificate database that is used to verify peer certificates.
// This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to %NULL, then
// peer certificate validation will always set the
// %G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// #GTlsConnection::accept-certificate will always be emitted on
// client-side connections, unless that bit is not set in
// #GTlsClientConnection:validation-flags).
/*

C function : g_tls_connection_set_database
*/
func (recv *TlsConnection) SetDatabase(database *TlsDatabase) {
	c_database := (*C.GTlsDatabase)(C.NULL)
	if database != nil {
		c_database = (*C.GTlsDatabase)(database.ToC())
	}

	C.g_tls_connection_set_database((*C.GTlsConnection)(recv.native), c_database)

	return
}

// Set the object that will be used to interact with the user. It will be used
// for things like prompting the user for passwords.
//
// The @interaction argument will normally be a derived subclass of
// #GTlsInteraction. %NULL can also be provided if no user interaction
// should occur for this connection.
/*

C function : g_tls_connection_set_interaction
*/
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

// Object upcasts to *Object
func (recv *TlsDatabase) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TlsDatabase.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsDatabase.
func CastToTlsDatabase(object *gobject.Object) *TlsDatabase {
	return TlsDatabaseNewFromC(object.ToC())
}

// Create a handle string for the certificate. The database will only be able
// to create a handle for certificates that originate from the database. In
// cases where the database cannot create a handle for a certificate, %NULL
// will be returned.
//
// This handle should be stable across various instances of the application,
// and between applications. If a certificate is modified in the database,
// then it is not guaranteed that this handle will continue to point to it.
/*

C function : g_tls_database_create_certificate_handle
*/
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

// Lookup a certificate by its handle.
//
// The handle should have been created by calling
// g_tls_database_create_certificate_handle() on a #GTlsDatabase object of
// the same TLS backend. The handle is designed to remain valid across
// instantiations of the database.
//
// If the handle is no longer valid, or does not point to a certificate in
// this database, then %NULL will be returned.
//
// This function can block, use g_tls_database_lookup_certificate_for_handle_async() to perform
// the lookup operation asynchronously.
/*

C function : g_tls_database_lookup_certificate_for_handle
*/
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

// Finish an asynchronous lookup of a certificate by its handle. See
// g_tls_database_lookup_certificate_by_handle() for more information.
//
// If the handle is no longer valid, or does not point to a certificate in
// this database, then %NULL will be returned.
/*

C function : g_tls_database_lookup_certificate_for_handle_finish
*/
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

// Lookup the issuer of @certificate in the database.
//
// The %issuer property
// of @certificate is not modified, and the two certificates are not hooked
// into a chain.
//
// This function can block, use g_tls_database_lookup_certificate_issuer_async() to perform
// the lookup operation asynchronously.
/*

C function : g_tls_database_lookup_certificate_issuer
*/
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

// Finish an asynchronous lookup issuer operation. See
// g_tls_database_lookup_certificate_issuer() for more information.
/*

C function : g_tls_database_lookup_certificate_issuer_finish
*/
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

// Lookup certificates issued by this issuer in the database.
//
// This function can block, use g_tls_database_lookup_certificates_issued_by_async() to perform
// the lookup operation asynchronously.
/*

C function : g_tls_database_lookup_certificates_issued_by
*/
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

// Finish an asynchronous lookup of certificates. See
// g_tls_database_lookup_certificates_issued_by() for more information.
/*

C function : g_tls_database_lookup_certificates_issued_by_finish
*/
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

// Determines the validity of a certificate chain after looking up and
// adding any missing certificates to the chain.
//
// @chain is a chain of #GTlsCertificate objects each pointing to the next
// certificate in the chain by its #GTlsCertificate:issuer property. The chain may initially
// consist of one or more certificates. After the verification process is
// complete, @chain may be modified by adding missing certificates, or removing
// extra certificates. If a certificate anchor was found, then it is added to
// the @chain.
//
// @purpose describes the purpose (or usage) for which the certificate
// is being used. Typically @purpose will be set to #G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER
// which means that the certificate is being used to authenticate a server
// (and we are acting as the client).
//
// The @identity is used to check for pinned certificates (trust exceptions)
// in the database. These will override the normal verification process on a
// host by host basis.
//
// Currently there are no @flags, and %G_TLS_DATABASE_VERIFY_NONE should be
// used.
//
// If @chain is found to be valid, then the return value will be 0. If
// @chain is found to be invalid, then the return value will indicate
// the problems found. If the function is unable to determine whether
// @chain is valid or not (eg, because @cancellable is triggered
// before it completes) then the return value will be
// %G_TLS_CERTIFICATE_GENERIC_ERROR and @error will be set
// accordingly. @error is not set when @chain is successfully analyzed
// but found to be invalid.
//
// This function can block, use g_tls_database_verify_chain_async() to perform
// the verification operation asynchronously.
/*

C function : g_tls_database_verify_chain
*/
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

// Finish an asynchronous verify chain operation. See
// g_tls_database_verify_chain() for more information.
//
// If @chain is found to be valid, then the return value will be 0. If
// @chain is found to be invalid, then the return value will indicate
// the problems found. If the function is unable to determine whether
// @chain is valid or not (eg, because @cancellable is triggered
// before it completes) then the return value will be
// %G_TLS_CERTIFICATE_GENERIC_ERROR and @error will be set
// accordingly. @error is not set when @chain is successfully analyzed
// but found to be invalid.
/*

C function : g_tls_database_verify_chain_finish
*/
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

// Object upcasts to *Object
func (recv *TlsInteraction) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TlsInteraction.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsInteraction.
func CastToTlsInteraction(object *gobject.Object) *TlsInteraction {
	return TlsInteractionNewFromC(object.ToC())
}

// Run synchronous interaction to ask the user for a password. In general,
// g_tls_interaction_invoke_ask_password() should be used instead of this
// function.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The @password value will
// be filled in and then @callback will be called. Alternatively the user may
// abort this password request, which will usually abort the TLS connection.
//
// If the interaction is cancelled by the cancellation object, or by the
// user then %G_TLS_INTERACTION_FAILED will be returned with an error that
// contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
// not support immediate cancellation.
/*

C function : g_tls_interaction_ask_password
*/
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

// Complete an ask password user interaction request. This should be once
// the g_tls_interaction_ask_password_async() completion callback is called.
//
// If %G_TLS_INTERACTION_HANDLED is returned, then the #GTlsPassword passed
// to g_tls_interaction_ask_password() will have its password filled in.
//
// If the interaction is cancelled by the cancellation object, or by the
// user then %G_TLS_INTERACTION_FAILED will be returned with an error that
// contains a %G_IO_ERROR_CANCELLED error code.
/*

C function : g_tls_interaction_ask_password_finish
*/
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

// Invoke the interaction to ask the user for a password. It invokes this
// interaction in the main loop, specifically the #GMainContext returned by
// g_main_context_get_thread_default() when the interaction is created. This
// is called by called by #GTlsConnection or #GTlsDatabase to ask the user
// for a password.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The @password value will
// be filled in and then @callback will be called. Alternatively the user may
// abort this password request, which will usually abort the TLS connection.
//
// The implementation can either be a synchronous (eg: modal dialog) or an
// asynchronous one (eg: modeless dialog). This function will take care of
// calling which ever one correctly.
//
// If the interaction is cancelled by the cancellation object, or by the
// user then %G_TLS_INTERACTION_FAILED will be returned with an error that
// contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
// not support immediate cancellation.
/*

C function : g_tls_interaction_invoke_ask_password
*/
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

// Object upcasts to *Object
func (recv *TlsPassword) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TlsPassword.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsPassword.
func CastToTlsPassword(object *gobject.Object) *TlsPassword {
	return TlsPasswordNewFromC(object.ToC())
}

// Create a new #GTlsPassword object.
/*

C function : g_tls_password_new
*/
func TlsPasswordNew(flags TlsPasswordFlags, description string) *TlsPassword {
	c_flags := (C.GTlsPasswordFlags)(flags)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	retC := C.g_tls_password_new(c_flags, c_description)
	retGo := TlsPasswordNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get a description string about what the password will be used for.
/*

C function : g_tls_password_get_description
*/
func (recv *TlsPassword) GetDescription() string {
	retC := C.g_tls_password_get_description((*C.GTlsPassword)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Get flags about the password.
/*

C function : g_tls_password_get_flags
*/
func (recv *TlsPassword) GetFlags() TlsPasswordFlags {
	retC := C.g_tls_password_get_flags((*C.GTlsPassword)(recv.native))
	retGo := (TlsPasswordFlags)(retC)

	return retGo
}

// Blacklisted : g_tls_password_get_value

// Get a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
/*

C function : g_tls_password_get_warning
*/
func (recv *TlsPassword) GetWarning() string {
	retC := C.g_tls_password_get_warning((*C.GTlsPassword)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Set a description string about what the password will be used for.
/*

C function : g_tls_password_set_description
*/
func (recv *TlsPassword) SetDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_tls_password_set_description((*C.GTlsPassword)(recv.native), c_description)

	return
}

// Set flags about the password.
/*

C function : g_tls_password_set_flags
*/
func (recv *TlsPassword) SetFlags(flags TlsPasswordFlags) {
	c_flags := (C.GTlsPasswordFlags)(flags)

	C.g_tls_password_set_flags((*C.GTlsPassword)(recv.native), c_flags)

	return
}

// Set the value for this password. The @value will be copied by the password
// object.
//
// Specify the @length, for a non-nul-terminated password. Pass -1 as
// @length if using a nul-terminated password, and @length will be
// calculated automatically. (Note that the terminating nul is not
// considered part of the password in this case.)
/*

C function : g_tls_password_set_value
*/
func (recv *TlsPassword) SetValue(value []uint8) {
	c_value := &value[0]

	c_length := (C.gssize)(len(value))

	C.g_tls_password_set_value((*C.GTlsPassword)(recv.native), (*C.guchar)(unsafe.Pointer(c_value)), c_length)

	return
}

// Unsupported : g_tls_password_set_value_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Set a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
/*

C function : g_tls_password_set_warning
*/
func (recv *TlsPassword) SetWarning(warning string) {
	c_warning := C.CString(warning)
	defer C.free(unsafe.Pointer(c_warning))

	C.g_tls_password_set_warning((*C.GTlsPassword)(recv.native), c_warning)

	return
}
