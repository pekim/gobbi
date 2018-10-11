// This is a generated file - DO NOT EDIT
// +build gio_2.56

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

// SetOptionContextDescription is a wrapper around the C function g_application_set_option_context_description.
func (recv *Application) SetOptionContextDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_application_set_option_context_description((*C.GApplication)(recv.native), c_description)

	return
}

// SetOptionContextParameterString is a wrapper around the C function g_application_set_option_context_parameter_string.
func (recv *Application) SetOptionContextParameterString(parameterString string) {
	c_parameter_string := C.CString(parameterString)
	defer C.free(unsafe.Pointer(c_parameter_string))

	C.g_application_set_option_context_parameter_string((*C.GApplication)(recv.native), c_parameter_string)

	return
}

// SetOptionContextSummary is a wrapper around the C function g_application_set_option_context_summary.
func (recv *Application) SetOptionContextSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_application_set_option_context_summary((*C.GApplication)(recv.native), c_summary)

	return
}

func (recv *Application) Object() *gobject.Object {}

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

func (recv *DBusAuthObserver) Object() *gobject.Object {}

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

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

func (recv *DBusObjectManagerServer) Object() *gobject.Object {}

func (recv *DBusObjectProxy) Object() *gobject.Object {}

func (recv *DBusObjectSkeleton) Object() *gobject.Object {}

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

func (recv *DBusProxy) Object() *gobject.Object {}

func (recv *DBusServer) Object() *gobject.Object {}

func (recv *DataInputStream) BufferedInputStream() *BufferedInputStream {}

func (recv *DataOutputStream) FilterOutputStream() *FilterOutputStream {}

// GetLocaleString is a wrapper around the C function g_desktop_app_info_get_locale_string.
func (recv *DesktopAppInfo) GetLocaleString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_locale_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

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

func (recv *InputStream) Object() *gobject.Object {}

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

func (recv *ListStore) Object() *gobject.Object {}

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

func (recv *MemoryInputStream) InputStream() *InputStream {}

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

func (recv *MemoryOutputStream) OutputStream() *OutputStream {}

func (recv *Menu) MenuModel() *MenuModel {}

func (recv *MenuAttributeIter) Object() *gobject.Object {}

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

func (recv *ProxyAddress) InetSocketAddress() *InetSocketAddress {}

func (recv *ProxyAddressEnumerator) SocketAddressEnumerator() *SocketAddressEnumerator {}

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

// JoinMulticastGroupSsm is a wrapper around the C function g_socket_join_multicast_group_ssm.
func (recv *Socket) JoinMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(group.ToC())

	c_source_specific := (*C.GInetAddress)(sourceSpecific.ToC())

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LeaveMulticastGroupSsm is a wrapper around the C function g_socket_leave_multicast_group_ssm.
func (recv *Socket) LeaveMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(group.ToC())

	c_source_specific := (*C.GInetAddress)(sourceSpecific.ToC())

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

func (recv *TestDBus) Object() *gobject.Object {}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

func (recv *ThemedIcon) Object() *gobject.Object {}

func (recv *ThreadedSocketService) SocketService() *SocketService {}

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
