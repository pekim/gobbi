// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// MarkBusy is a wrapper around the C function g_application_mark_busy.
func (recv *Application) MarkBusy() {
	C.g_application_mark_busy((*C.GApplication)(recv.native))

	return
}

// UnmarkBusy is a wrapper around the C function g_application_unmark_busy.
func (recv *Application) UnmarkBusy() {
	C.g_application_unmark_busy((*C.GApplication)(recv.native))

	return
}

func (recv *Application) Object() *gobject.Object {}

func (recv *ApplicationCommandLine) Object() *gobject.Object {}

func (recv *BufferedInputStream) FilterInputStream() *FilterInputStream {}

func (recv *BufferedOutputStream) FilterOutputStream() *FilterOutputStream {}

// BytesIconNew is a wrapper around the C function g_bytes_icon_new.
func BytesIconNew(bytes *glib.Bytes) *BytesIcon {
	c_bytes := (*C.GBytes)(bytes.ToC())

	retC := C.g_bytes_icon_new(c_bytes)
	retGo := BytesIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBytes is a wrapper around the C function g_bytes_icon_get_bytes.
func (recv *BytesIcon) GetBytes() *glib.Bytes {
	retC := C.g_bytes_icon_get_bytes((*C.GBytesIcon)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetPropertyInfo is a wrapper around the C function g_dbus_method_invocation_get_property_info.
func (recv *DBusMethodInvocation) GetPropertyInfo() *DBusPropertyInfo {
	retC := C.g_dbus_method_invocation_get_property_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetActionName is a wrapper around the C function g_desktop_app_info_get_action_name.
func (recv *DesktopAppInfo) GetActionName(actionName string) string {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_desktop_app_info_get_action_name((*C.GDesktopAppInfo)(recv.native), c_action_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// LaunchAction is a wrapper around the C function g_desktop_app_info_launch_action.
func (recv *DesktopAppInfo) LaunchAction(actionName string, launchContext *AppLaunchContext) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_launch_context := (*C.GAppLaunchContext)(launchContext.ToC())

	C.g_desktop_app_info_launch_action((*C.GDesktopAppInfo)(recv.native), c_action_name, c_launch_context)

	return
}

// Unsupported : g_desktop_app_info_list_actions : no return type

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

// RemoveAll is a wrapper around the C function g_menu_remove_all.
func (recv *Menu) RemoveAll() {
	C.g_menu_remove_all((*C.GMenu)(recv.native))

	return
}

func (recv *Menu) MenuModel() *MenuModel {}

func (recv *MenuAttributeIter) Object() *gobject.Object {}

// Unsupported : g_menu_item_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

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

// PropertyAction is a wrapper around the C record GPropertyAction.
type PropertyAction struct {
	native *C.GPropertyAction
}

func PropertyActionNewFromC(u unsafe.Pointer) *PropertyAction {
	c := (*C.GPropertyAction)(u)
	if c == nil {
		return nil
	}

	g := &PropertyAction{native: c}

	return g
}

func (recv *PropertyAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PropertyActionNew is a wrapper around the C function g_property_action_new.
func PropertyActionNew(name string, object uintptr, propertyName string) *PropertyAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object := (C.gpointer)(object)

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_property_action_new(c_name, c_object, c_property_name)
	retGo := PropertyActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

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
