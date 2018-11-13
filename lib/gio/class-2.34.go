// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Gets the #GDBusConnection being used by the application, or %NULL.
//
// If #GApplication is using its D-Bus backend then this function will
// return the #GDBusConnection being used for uniqueness and
// communication with the desktop environment and other instances of the
// application.
//
// If #GApplication is not using D-Bus then this function will return
// %NULL.  This includes the situation where the D-Bus backend would
// normally be in use but we were unable to connect to the bus.
//
// This function must not be called before the application has been
// registered.  See g_application_get_is_registered().
/*

C function : g_application_get_dbus_connection
*/
func (recv *Application) GetDbusConnection() *DBusConnection {
	retC := C.g_application_get_dbus_connection((*C.GApplication)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the D-Bus object path being used by the application, or %NULL.
//
// If #GApplication is using its D-Bus backend then this function will
// return the D-Bus object path that #GApplication is using.  If the
// application is the primary instance then there is an object published
// at this path.  If the application is not the primary instance then
// the result of this function is undefined.
//
// If #GApplication is not using D-Bus then this function will return
// %NULL.  This includes the situation where the D-Bus backend would
// normally be in use but we were unable to connect to the bus.
//
// This function must not be called before the application has been
// registered.  See g_application_get_is_registered().
/*

C function : g_application_get_dbus_object_path
*/
func (recv *Application) GetDbusObjectPath() string {
	retC := C.g_application_get_dbus_object_path((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the stdin of the invoking process.
//
// The #GInputStream can be used to read data passed to the standard
// input of the invoking process.
// This doesn't work on all platforms.  Presently, it is only available
// on UNIX when using a DBus daemon capable of passing file descriptors.
// If stdin is not available then %NULL will be returned.  In the
// future, support may be expanded to other platforms.
//
// You must only call this function once per commandline invocation.
/*

C function : g_application_command_line_get_stdin
*/
func (recv *ApplicationCommandLine) GetStdin() *InputStream {
	retC := C.g_application_command_line_get_stdin((*C.GApplicationCommandLine)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Emits the #GDBusAuthObserver::allow-mechanism signal on @observer.
/*

C function : g_dbus_auth_observer_allow_mechanism
*/
func (recv *DBusAuthObserver) AllowMechanism(mechanism string) bool {
	c_mechanism := C.CString(mechanism)
	defer C.free(unsafe.Pointer(c_mechanism))

	retC := C.g_dbus_auth_observer_allow_mechanism((*C.GDBusAuthObserver)(recv.native), c_mechanism)
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the last serial number assigned to a #GDBusMessage on
// the current thread. This includes messages sent via both low-level
// API such as g_dbus_connection_send_message() as well as
// high-level API such as g_dbus_connection_emit_signal(),
// g_dbus_connection_call() or g_dbus_proxy_call().
/*

C function : g_dbus_connection_get_last_serial
*/
func (recv *DBusConnection) GetLastSerial() uint32 {
	retC := C.g_dbus_connection_get_last_serial((*C.GDBusConnection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns whether @object is currently exported on @manager.
/*

C function : g_dbus_object_manager_server_is_exported
*/
func (recv *DBusObjectManagerServer) IsExported(object *DBusObjectSkeleton) bool {
	c_object := (*C.GDBusObjectSkeleton)(C.NULL)
	if object != nil {
		c_object = (*C.GDBusObjectSkeleton)(object.ToC())
	}

	retC := C.g_dbus_object_manager_server_is_exported((*C.GDBusObjectManagerServer)(recv.native), c_object)
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the StartupWMClass field from @info. This represents the
// WM_CLASS property of the main window of the application, if launched
// through @info.
/*

C function : g_desktop_app_info_get_startup_wm_class
*/
func (recv *DesktopAppInfo) GetStartupWmClass() string {
	retC := C.g_desktop_app_info_get_startup_wm_class((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the symbolic icon for a file.
/*

C function : g_file_info_get_symbolic_icon
*/
func (recv *FileInfo) GetSymbolicIcon() *Icon {
	retC := C.g_file_info_get_symbolic_icon((*C.GFileInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the symbolic icon for a given #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
/*

C function : g_file_info_set_symbolic_icon
*/
func (recv *FileInfo) SetSymbolicIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_file_info_set_symbolic_icon((*C.GFileInfo)(recv.native), c_icon)

	return
}

// Like g_input_stream_read(), this tries to read @count bytes from
// the stream in a blocking fashion. However, rather than reading into
// a user-supplied buffer, this will create a new #GBytes containing
// the data that was read. This may be easier to use from language
// bindings.
//
// If count is zero, returns a zero-length #GBytes and does nothing. A
// value of @count larger than %G_MAXSSIZE will cause a
// %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, a new #GBytes is returned. It is not an error if the
// size of this object is not the same as the requested size, as it
// can happen e.g. near the end of a file. A zero-length #GBytes is
// returned on end of file (or if @count is zero), but never
// otherwise.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error %NULL is returned and @error is set accordingly.
/*

C function : g_input_stream_read_bytes
*/
func (recv *InputStream) ReadBytes(count uint64, cancellable *Cancellable) (*glib.Bytes, error) {
	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_bytes((*C.GInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous stream read-into-#GBytes operation.
/*

C function : g_input_stream_read_bytes_finish
*/
func (recv *InputStream) ReadBytesFinish(result *AsyncResult) (*glib.Bytes, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_bytes_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new #GMemoryInputStream with data from the given @bytes.
/*

C function : g_memory_input_stream_new_from_bytes
*/
func MemoryInputStreamNewFromBytes(bytes *glib.Bytes) *MemoryInputStream {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	retC := C.g_memory_input_stream_new_from_bytes(c_bytes)
	retGo := MemoryInputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends @bytes to data that can be read from the input stream.
/*

C function : g_memory_input_stream_add_bytes
*/
func (recv *MemoryInputStream) AddBytes(bytes *glib.Bytes) {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	C.g_memory_input_stream_add_bytes((*C.GMemoryInputStream)(recv.native), c_bytes)

	return
}

// Returns data from the @ostream as a #GBytes. @ostream must be
// closed before calling this function.
/*

C function : g_memory_output_stream_steal_as_bytes
*/
func (recv *MemoryOutputStream) StealAsBytes() *glib.Bytes {
	retC := C.g_memory_output_stream_steal_as_bytes((*C.GMemoryOutputStream)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GMenuItem as an exact copy of an existing menu item in a
// #GMenuModel.
//
// @item_index must be valid (ie: be sure to call
// g_menu_model_get_n_items() first).
/*

C function : g_menu_item_new_from_model
*/
func MenuItemNewFromModel(model *MenuModel, itemIndex int32) *MenuItem {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_item_new_from_model(c_model, c_item_index)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Queries the named @link on @menu_item.
/*

C function : g_menu_item_get_link
*/
func (recv *MenuItem) GetLink(link string) *MenuModel {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_item_get_link((*C.GMenuItem)(recv.native), c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'show-unmount-progress' for MountOperation : unsupported parameter message : type utf8 :

// Gets the protocol that is being spoken to the destination
// server; eg, "http" or "ftp".
/*

C function : g_proxy_address_get_destination_protocol
*/
func (recv *ProxyAddress) GetDestinationProtocol() string {
	retC := C.g_proxy_address_get_destination_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the proxy URI that @proxy was constructed from.
/*

C function : g_proxy_address_get_uri
*/
func (recv *ProxyAddress) GetUri() string {
	retC := C.g_proxy_address_get_uri((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Synchronously performs a DNS record lookup for the given @rrname and returns
// a list of records as #GVariant tuples. See #GResolverRecordType for
// information on what the records contain for each @record_type.
//
// If the DNS resolution fails, @error (if non-%NULL) will be set to
// a value from #GResolverError and %NULL will be returned.
//
// If @cancellable is non-%NULL, it can be used to cancel the
// operation, in which case @error (if non-%NULL) will be set to
// %G_IO_ERROR_CANCELLED.
/*

C function : g_resolver_lookup_records
*/
func (recv *Resolver) LookupRecords(rrname string, recordType ResolverRecordType, cancellable *Cancellable) (*glib.List, error) {
	c_rrname := C.CString(rrname)
	defer C.free(unsafe.Pointer(c_rrname))

	c_record_type := (C.GResolverRecordType)(recordType)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_records((*C.GResolver)(recv.native), c_rrname, c_record_type, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_records_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Retrieves the result of a previous call to
// g_resolver_lookup_records_async(). Returns a non-empty list of records as
// #GVariant tuples. See #GResolverRecordType for information on what the
// records contain.
//
// If the DNS resolution failed, @error (if non-%NULL) will be set to
// a value from #GResolverError. If the operation was cancelled,
// @error will be set to %G_IO_ERROR_CANCELLED.
/*

C function : g_resolver_lookup_records_finish
*/
func (recv *Resolver) LookupRecordsFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_records_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

	return g
}

func (recv *TestDBus) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TestDBus) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TestDBus.
// Exercise care, as this is a potentially dangerous function if the Object is not a TestDBus.
func CastToTestDBus(object *gobject.Object) *TestDBus {
	return TestDBusNewFromC(object.ToC())
}

// Create a new #GTestDBus object.
/*

C function : g_test_dbus_new
*/
func TestDBusNew(flags TestDBusFlags) *TestDBus {
	c_flags := (C.GTestDBusFlags)(flags)

	retC := C.g_test_dbus_new(c_flags)
	retGo := TestDBusNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add a path where dbus-daemon will look up .service files. This can't be
// called after g_test_dbus_up().
/*

C function : g_test_dbus_add_service_dir
*/
func (recv *TestDBus) AddServiceDir(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_test_dbus_add_service_dir((*C.GTestDBus)(recv.native), c_path)

	return
}

// Stop the session bus started by g_test_dbus_up().
//
// This will wait for the singleton returned by g_bus_get() or g_bus_get_sync()
// is destroyed. This is done to ensure that the next unit test won't get a
// leaked singleton from this test.
/*

C function : g_test_dbus_down
*/
func (recv *TestDBus) Down() {
	C.g_test_dbus_down((*C.GTestDBus)(recv.native))

	return
}

// Get the address on which dbus-daemon is running. If g_test_dbus_up() has not
// been called yet, %NULL is returned. This can be used with
// g_dbus_connection_new_for_address().
/*

C function : g_test_dbus_get_bus_address
*/
func (recv *TestDBus) GetBusAddress() string {
	retC := C.g_test_dbus_get_bus_address((*C.GTestDBus)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Get the flags of the #GTestDBus object.
/*

C function : g_test_dbus_get_flags
*/
func (recv *TestDBus) GetFlags() TestDBusFlags {
	retC := C.g_test_dbus_get_flags((*C.GTestDBus)(recv.native))
	retGo := (TestDBusFlags)(retC)

	return retGo
}

// Stop the session bus started by g_test_dbus_up().
//
// Unlike g_test_dbus_down(), this won't verify the #GDBusConnection
// singleton returned by g_bus_get() or g_bus_get_sync() is destroyed. Unit
// tests wanting to verify behaviour after the session bus has been stopped
// can use this function but should still call g_test_dbus_down() when done.
/*

C function : g_test_dbus_stop
*/
func (recv *TestDBus) Stop() {
	C.g_test_dbus_stop((*C.GTestDBus)(recv.native))

	return
}

// Start a dbus-daemon instance and set DBUS_SESSION_BUS_ADDRESS. After this
// call, it is safe for unit tests to start sending messages on the session bus.
//
// If this function is called from setup callback of g_test_add(),
// g_test_dbus_down() must be called in its teardown callback.
//
// If this function is called from unit test's main(), then g_test_dbus_down()
// must be called after g_test_run().
/*

C function : g_test_dbus_up
*/
func (recv *TestDBus) Up() {
	C.g_test_dbus_up((*C.GTestDBus)(recv.native))

	return
}

// Check if two #GTlsCertificate objects represent the same certificate.
// The raw DER byte data of the two certificates are checked for equality.
// This has the effect that two certificates may compare equal even if
// their #GTlsCertificate:issuer, #GTlsCertificate:private-key, or
// #GTlsCertificate:private-key-pem properties differ.
/*

C function : g_tls_certificate_is_same
*/
func (recv *TlsCertificate) IsSame(certTwo *TlsCertificate) bool {
	c_cert_two := (*C.GTlsCertificate)(C.NULL)
	if certTwo != nil {
		c_cert_two = (*C.GTlsCertificate)(certTwo.ToC())
	}

	retC := C.g_tls_certificate_is_same((*C.GTlsCertificate)(recv.native), c_cert_two)
	retGo := retC == C.TRUE

	return retGo
}
