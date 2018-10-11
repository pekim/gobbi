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

func (recv *AppInfoMonitor) Object() *gobject.Object {}

func (recv *AppLaunchContext) Object() *gobject.Object {}

// GetDbusConnection is a wrapper around the C function g_application_get_dbus_connection.
func (recv *Application) GetDbusConnection() *DBusConnection {
	retC := C.g_application_get_dbus_connection((*C.GApplication)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDbusObjectPath is a wrapper around the C function g_application_get_dbus_object_path.
func (recv *Application) GetDbusObjectPath() string {
	retC := C.g_application_get_dbus_object_path((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

func (recv *Application) Object() *gobject.Object {}

// GetStdin is a wrapper around the C function g_application_command_line_get_stdin.
func (recv *ApplicationCommandLine) GetStdin() *InputStream {
	retC := C.g_application_command_line_get_stdin((*C.GApplicationCommandLine)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *ApplicationCommandLine) Object() *gobject.Object {}

func (recv *BufferedInputStream) FilterInputStream() *FilterInputStream {}

func (recv *BufferedOutputStream) FilterOutputStream() *FilterOutputStream {}

func (recv *BytesIcon) Object() *gobject.Object {}

func (recv *Cancellable) Object() *gobject.Object {}

func (recv *CharsetConverter) Object() *gobject.Object {}

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

func (recv *ConverterInputStream) FilterInputStream() *FilterInputStream {}

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

func (recv *ConverterOutputStream) FilterOutputStream() *FilterOutputStream {}

func (recv *Credentials) Object() *gobject.Object {}

func (recv *DBusActionGroup) Object() *gobject.Object {}

// AllowMechanism is a wrapper around the C function g_dbus_auth_observer_allow_mechanism.
func (recv *DBusAuthObserver) AllowMechanism(mechanism string) bool {
	c_mechanism := C.CString(mechanism)
	defer C.free(unsafe.Pointer(c_mechanism))

	retC := C.g_dbus_auth_observer_allow_mechanism((*C.GDBusAuthObserver)(recv.native), c_mechanism)
	retGo := retC == C.TRUE

	return retGo
}

func (recv *DBusAuthObserver) Object() *gobject.Object {}

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// GetLastSerial is a wrapper around the C function g_dbus_connection_get_last_serial.
func (recv *DBusConnection) GetLastSerial() uint32 {
	retC := C.g_dbus_connection_get_last_serial((*C.GDBusConnection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

func (recv *DBusConnection) Object() *gobject.Object {}

func (recv *DBusInterfaceSkeleton) Object() *gobject.Object {}

func (recv *DBusMenuModel) MenuModel() *MenuModel {}

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

func (recv *DBusMessage) Object() *gobject.Object {}

func (recv *DBusMethodInvocation) Object() *gobject.Object {}

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

func (recv *DBusObjectManagerClient) Object() *gobject.Object {}

// IsExported is a wrapper around the C function g_dbus_object_manager_server_is_exported.
func (recv *DBusObjectManagerServer) IsExported(object *DBusObjectSkeleton) bool {
	c_object := (*C.GDBusObjectSkeleton)(object.ToC())

	retC := C.g_dbus_object_manager_server_is_exported((*C.GDBusObjectManagerServer)(recv.native), c_object)
	retGo := retC == C.TRUE

	return retGo
}

func (recv *DBusObjectManagerServer) Object() *gobject.Object {}

func (recv *DBusObjectProxy) Object() *gobject.Object {}

func (recv *DBusObjectSkeleton) Object() *gobject.Object {}

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

func (recv *DBusProxy) Object() *gobject.Object {}

func (recv *DBusServer) Object() *gobject.Object {}

func (recv *DataInputStream) BufferedInputStream() *BufferedInputStream {}

func (recv *DataOutputStream) FilterOutputStream() *FilterOutputStream {}

// GetStartupWmClass is a wrapper around the C function g_desktop_app_info_get_startup_wm_class.
func (recv *DesktopAppInfo) GetStartupWmClass() string {
	retC := C.g_desktop_app_info_get_startup_wm_class((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

func (recv *DesktopAppInfo) Object() *gobject.Object {}

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

func (recv *Emblem) Object() *gobject.Object {}

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

func (recv *EmblemedIcon) Object() *gobject.Object {}

func (recv *FileEnumerator) Object() *gobject.Object {}

func (recv *FileIOStream) IOStream() *IOStream {}

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

func (recv *FileIcon) Object() *gobject.Object {}

// Unsupported : g_file_info_get_symbolic_icon : no return generator

// Unsupported : g_file_info_set_symbolic_icon : unsupported parameter icon : no type generator for Icon, GIcon*

func (recv *FileInfo) Object() *gobject.Object {}

func (recv *FileInputStream) InputStream() *InputStream {}

func (recv *FileMonitor) Object() *gobject.Object {}

func (recv *FileOutputStream) OutputStream() *OutputStream {}

func (recv *FilenameCompleter) Object() *gobject.Object {}

func (recv *FilterInputStream) InputStream() *InputStream {}

func (recv *FilterOutputStream) OutputStream() *OutputStream {}

func (recv *IOModule) TypeModule() *gobject.TypeModule {}

func (recv *IOStream) Object() *gobject.Object {}

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

func (recv *InetAddress) Object() *gobject.Object {}

func (recv *InetAddressMask) Object() *gobject.Object {}

func (recv *InetSocketAddress) SocketAddress() *SocketAddress {}

// ReadBytes is a wrapper around the C function g_input_stream_read_bytes.
func (recv *InputStream) ReadBytes(count uint64, cancellable *Cancellable) (*glib.Bytes, error) {
	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_bytes((*C.GInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_read_bytes_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

func (recv *InputStream) Object() *gobject.Object {}

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

func (recv *ListStore) Object() *gobject.Object {}

// MemoryInputStreamNewFromBytes is a wrapper around the C function g_memory_input_stream_new_from_bytes.
func MemoryInputStreamNewFromBytes(bytes *glib.Bytes) *MemoryInputStream {
	c_bytes := (*C.GBytes)(bytes.ToC())

	retC := C.g_memory_input_stream_new_from_bytes(c_bytes)
	retGo := MemoryInputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// AddBytes is a wrapper around the C function g_memory_input_stream_add_bytes.
func (recv *MemoryInputStream) AddBytes(bytes *glib.Bytes) {
	c_bytes := (*C.GBytes)(bytes.ToC())

	C.g_memory_input_stream_add_bytes((*C.GMemoryInputStream)(recv.native), c_bytes)

	return
}

func (recv *MemoryInputStream) InputStream() *InputStream {}

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// StealAsBytes is a wrapper around the C function g_memory_output_stream_steal_as_bytes.
func (recv *MemoryOutputStream) StealAsBytes() *glib.Bytes {
	retC := C.g_memory_output_stream_steal_as_bytes((*C.GMemoryOutputStream)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *MemoryOutputStream) OutputStream() *OutputStream {}

func (recv *Menu) MenuModel() *MenuModel {}

func (recv *MenuAttributeIter) Object() *gobject.Object {}

// MenuItemNewFromModel is a wrapper around the C function g_menu_item_new_from_model.
func MenuItemNewFromModel(model *MenuModel, itemIndex int32) *MenuItem {
	c_model := (*C.GMenuModel)(model.ToC())

	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_item_new_from_model(c_model, c_item_index)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// GetLink is a wrapper around the C function g_menu_item_get_link.
func (recv *MenuItem) GetLink(link string) *MenuModel {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_item_get_link((*C.GMenuItem)(recv.native), c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *MenuItem) Object() *gobject.Object {}

func (recv *MenuLinkIter) Object() *gobject.Object {}

func (recv *MenuModel) Object() *gobject.Object {}

func (recv *MountOperation) Object() *gobject.Object {}

func (recv *NativeVolumeMonitor) VolumeMonitor() *VolumeMonitor {}

func (recv *NetworkAddress) Object() *gobject.Object {}

func (recv *NetworkService) Object() *gobject.Object {}

func (recv *Notification) Object() *gobject.Object {}

func (recv *OutputStream) Object() *gobject.Object {}

func (recv *Permission) Object() *gobject.Object {}

func (recv *PropertyAction) Object() *gobject.Object {}

// GetDestinationProtocol is a wrapper around the C function g_proxy_address_get_destination_protocol.
func (recv *ProxyAddress) GetDestinationProtocol() string {
	retC := C.g_proxy_address_get_destination_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function g_proxy_address_get_uri.
func (recv *ProxyAddress) GetUri() string {
	retC := C.g_proxy_address_get_uri((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

func (recv *ProxyAddress) InetSocketAddress() *InetSocketAddress {}

func (recv *ProxyAddressEnumerator) SocketAddressEnumerator() *SocketAddressEnumerator {}

// LookupRecords is a wrapper around the C function g_resolver_lookup_records.
func (recv *Resolver) LookupRecords(rrname string, recordType ResolverRecordType, cancellable *Cancellable) (*glib.List, error) {
	c_rrname := C.CString(rrname)
	defer C.free(unsafe.Pointer(c_rrname))

	c_record_type := (C.GResolverRecordType)(recordType)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_records((*C.GResolver)(recv.native), c_rrname, c_record_type, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_records_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_records_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

func (recv *Resolver) Object() *gobject.Object {}

func (recv *Settings) Object() *gobject.Object {}

func (recv *SettingsBackend) Object() *gobject.Object {}

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

func (recv *SimpleAction) Object() *gobject.Object {}

func (recv *SimpleActionGroup) Object() *gobject.Object {}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

func (recv *SimpleAsyncResult) Object() *gobject.Object {}

func (recv *SimpleIOStream) IOStream() *IOStream {}

func (recv *SimplePermission) Permission() *Permission {}

func (recv *SimpleProxyResolver) Object() *gobject.Object {}

func (recv *Socket) Object() *gobject.Object {}

func (recv *SocketAddress) Object() *gobject.Object {}

func (recv *SocketAddressEnumerator) Object() *gobject.Object {}

func (recv *SocketClient) Object() *gobject.Object {}

func (recv *SocketConnection) IOStream() *IOStream {}

func (recv *SocketControlMessage) Object() *gobject.Object {}

func (recv *SocketListener) Object() *gobject.Object {}

func (recv *SocketService) SocketListener() *SocketListener {}

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

func (recv *Subprocess) Object() *gobject.Object {}

func (recv *SubprocessLauncher) Object() *gobject.Object {}

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

func (recv *Task) Object() *gobject.Object {}

func (recv *TcpConnection) SocketConnection() *SocketConnection {}

func (recv *TcpWrapperConnection) TcpConnection() *TcpConnection {}

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

// TestDBusNew is a wrapper around the C function g_test_dbus_new.
func TestDBusNew(flags TestDBusFlags) *TestDBus {
	c_flags := (C.GTestDBusFlags)(flags)

	retC := C.g_test_dbus_new(c_flags)
	retGo := TestDBusNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddServiceDir is a wrapper around the C function g_test_dbus_add_service_dir.
func (recv *TestDBus) AddServiceDir(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_test_dbus_add_service_dir((*C.GTestDBus)(recv.native), c_path)

	return
}

// Down is a wrapper around the C function g_test_dbus_down.
func (recv *TestDBus) Down() {
	C.g_test_dbus_down((*C.GTestDBus)(recv.native))

	return
}

// GetBusAddress is a wrapper around the C function g_test_dbus_get_bus_address.
func (recv *TestDBus) GetBusAddress() string {
	retC := C.g_test_dbus_get_bus_address((*C.GTestDBus)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_test_dbus_get_flags.
func (recv *TestDBus) GetFlags() TestDBusFlags {
	retC := C.g_test_dbus_get_flags((*C.GTestDBus)(recv.native))
	retGo := (TestDBusFlags)(retC)

	return retGo
}

// Stop is a wrapper around the C function g_test_dbus_stop.
func (recv *TestDBus) Stop() {
	C.g_test_dbus_stop((*C.GTestDBus)(recv.native))

	return
}

// Up is a wrapper around the C function g_test_dbus_up.
func (recv *TestDBus) Up() {
	C.g_test_dbus_up((*C.GTestDBus)(recv.native))

	return
}

func (recv *TestDBus) Object() *gobject.Object {}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

func (recv *ThemedIcon) Object() *gobject.Object {}

func (recv *ThreadedSocketService) SocketService() *SocketService {}

// IsSame is a wrapper around the C function g_tls_certificate_is_same.
func (recv *TlsCertificate) IsSame(certTwo *TlsCertificate) bool {
	c_cert_two := (*C.GTlsCertificate)(certTwo.ToC())

	retC := C.g_tls_certificate_is_same((*C.GTlsCertificate)(recv.native), c_cert_two)
	retGo := retC == C.TRUE

	return retGo
}

func (recv *TlsCertificate) Object() *gobject.Object {}

func (recv *TlsConnection) IOStream() *IOStream {}

func (recv *TlsDatabase) Object() *gobject.Object {}

func (recv *TlsInteraction) Object() *gobject.Object {}

func (recv *TlsPassword) Object() *gobject.Object {}

func (recv *UnixConnection) SocketConnection() *SocketConnection {}

func (recv *UnixCredentialsMessage) SocketControlMessage() *SocketControlMessage {}

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

func (recv *UnixFDList) Object() *gobject.Object {}

func (recv *UnixFDMessage) SocketControlMessage() *SocketControlMessage {}

func (recv *UnixInputStream) InputStream() *InputStream {}

func (recv *UnixMountMonitor) Object() *gobject.Object {}

func (recv *UnixOutputStream) OutputStream() *OutputStream {}

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

func (recv *UnixSocketAddress) SocketAddress() *SocketAddress {}

func (recv *Vfs) Object() *gobject.Object {}

func (recv *VolumeMonitor) Object() *gobject.Object {}

func (recv *ZlibCompressor) Object() *gobject.Object {}

func (recv *ZlibDecompressor) Object() *gobject.Object {}
