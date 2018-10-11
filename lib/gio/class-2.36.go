// This is a generated file - DO NOT EDIT
// +build gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

func (recv *Application) Object() *gobject.Object {}

// Unsupported : g_application_command_line_create_file_for_arg : no return generator

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

// Unsupported : g_credentials_get_unix_pid : no return generator

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

// GetBoolean is a wrapper around the C function g_desktop_app_info_get_boolean.
func (recv *DesktopAppInfo) GetBoolean(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_boolean((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// GetString is a wrapper around the C function g_desktop_app_info_get_string.
func (recv *DesktopAppInfo) GetString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HasKey is a wrapper around the C function g_desktop_app_info_has_key.
func (recv *DesktopAppInfo) HasKey(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_has_key((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

func (recv *DesktopAppInfo) Object() *gobject.Object {}

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

func (recv *Emblem) Object() *gobject.Object {}

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

func (recv *EmblemedIcon) Object() *gobject.Object {}

// Unsupported : g_file_enumerator_get_child : no return generator

func (recv *FileEnumerator) Object() *gobject.Object {}

func (recv *FileIOStream) IOStream() *IOStream {}

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

func (recv *FileIcon) Object() *gobject.Object {}

// GetDeletionDate is a wrapper around the C function g_file_info_get_deletion_date.
func (recv *FileInfo) GetDeletionDate() *glib.DateTime {
	retC := C.g_file_info_get_deletion_date((*C.GFileInfo)(recv.native))
	retGo := glib.DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// MemoryOutputStreamNewResizable is a wrapper around the C function g_memory_output_stream_new_resizable.
func MemoryOutputStreamNewResizable() *MemoryOutputStream {
	retC := C.g_memory_output_stream_new_resizable()
	retGo := MemoryOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// SetDefaultProxy is a wrapper around the C function g_simple_proxy_resolver_set_default_proxy.
func (recv *SimpleProxyResolver) SetDefaultProxy(defaultProxy string) {
	c_default_proxy := C.CString(defaultProxy)
	defer C.free(unsafe.Pointer(c_default_proxy))

	C.g_simple_proxy_resolver_set_default_proxy((*C.GSimpleProxyResolver)(recv.native), c_default_proxy)

	return
}

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// SetUriProxy is a wrapper around the C function g_simple_proxy_resolver_set_uri_proxy.
func (recv *SimpleProxyResolver) SetUriProxy(uriScheme string, proxy string) {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	c_proxy := C.CString(proxy)
	defer C.free(unsafe.Pointer(c_proxy))

	C.g_simple_proxy_resolver_set_uri_proxy((*C.GSimpleProxyResolver)(recv.native), c_uri_scheme, c_proxy)

	return
}

func (recv *SimpleProxyResolver) Object() *gobject.Object {}

// GetOption is a wrapper around the C function g_socket_get_option.
func (recv *Socket) GetOption(level int32, optname int32) (bool, *int32, error) {
	c_level := (C.gint)(level)

	c_optname := (C.gint)(optname)

	var c_value C.gint

	var cThrowableError *C.GError

	retC := C.g_socket_get_option((*C.GSocket)(recv.native), c_level, c_optname, &c_value, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	value := (*int32)(&c_value)

	return retGo, value, goThrowableError
}

// SetOption is a wrapper around the C function g_socket_set_option.
func (recv *Socket) SetOption(level int32, optname int32, value int32) (bool, error) {
	c_level := (C.gint)(level)

	c_optname := (C.gint)(optname)

	c_value := (C.gint)(value)

	var cThrowableError *C.GError

	retC := C.g_socket_set_option((*C.GSocket)(recv.native), c_level, c_optname, c_value, &cThrowableError)
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

// Unsupported : g_socket_client_get_proxy_resolver : no return generator

// Unsupported : g_socket_client_set_proxy_resolver : unsupported parameter proxy_resolver : no type generator for ProxyResolver, GProxyResolver*

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

// Unsupported : g_task_attach_source : unsupported parameter callback : no type generator for GLib.SourceFunc, GSourceFunc

// GetCancellable is a wrapper around the C function g_task_get_cancellable.
func (recv *Task) GetCancellable() *Cancellable {
	retC := C.g_task_get_cancellable((*C.GTask)(recv.native))
	retGo := CancellableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCheckCancellable is a wrapper around the C function g_task_get_check_cancellable.
func (recv *Task) GetCheckCancellable() bool {
	retC := C.g_task_get_check_cancellable((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function g_task_get_context.
func (recv *Task) GetContext() *glib.MainContext {
	retC := C.g_task_get_context((*C.GTask)(recv.native))
	retGo := glib.MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPriority is a wrapper around the C function g_task_get_priority.
func (recv *Task) GetPriority() int32 {
	retC := C.g_task_get_priority((*C.GTask)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReturnOnCancel is a wrapper around the C function g_task_get_return_on_cancel.
func (recv *Task) GetReturnOnCancel() bool {
	retC := C.g_task_get_return_on_cancel((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSourceObject is a wrapper around the C function g_task_get_source_object.
func (recv *Task) GetSourceObject() uintptr {
	retC := C.g_task_get_source_object((*C.GTask)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetSourceTag is a wrapper around the C function g_task_get_source_tag.
func (recv *Task) GetSourceTag() uintptr {
	retC := C.g_task_get_source_tag((*C.GTask)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetTaskData is a wrapper around the C function g_task_get_task_data.
func (recv *Task) GetTaskData() uintptr {
	retC := C.g_task_get_task_data((*C.GTask)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// HadError is a wrapper around the C function g_task_had_error.
func (recv *Task) HadError() bool {
	retC := C.g_task_had_error((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PropagateBoolean is a wrapper around the C function g_task_propagate_boolean.
func (recv *Task) PropagateBoolean() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_task_propagate_boolean((*C.GTask)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PropagateInt is a wrapper around the C function g_task_propagate_int.
func (recv *Task) PropagateInt() (int64, error) {
	var cThrowableError *C.GError

	retC := C.g_task_propagate_int((*C.GTask)(recv.native), &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PropagatePointer is a wrapper around the C function g_task_propagate_pointer.
func (recv *Task) PropagatePointer() (uintptr, error) {
	var cThrowableError *C.GError

	retC := C.g_task_propagate_pointer((*C.GTask)(recv.native), &cThrowableError)
	retGo := (uintptr)(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ReturnBoolean is a wrapper around the C function g_task_return_boolean.
func (recv *Task) ReturnBoolean(result bool) {
	c_result :=
		boolToGboolean(result)

	C.g_task_return_boolean((*C.GTask)(recv.native), c_result)

	return
}

// ReturnError is a wrapper around the C function g_task_return_error.
func (recv *Task) ReturnError(error *glib.Error) {
	c_error := (*C.GError)(error.ToC())

	C.g_task_return_error((*C.GTask)(recv.native), c_error)

	return
}

// ReturnErrorIfCancelled is a wrapper around the C function g_task_return_error_if_cancelled.
func (recv *Task) ReturnErrorIfCancelled() bool {
	retC := C.g_task_return_error_if_cancelled((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReturnInt is a wrapper around the C function g_task_return_int.
func (recv *Task) ReturnInt(result int64) {
	c_result := (C.gssize)(result)

	C.g_task_return_int((*C.GTask)(recv.native), c_result)

	return
}

// Unsupported : g_task_return_new_error : unsupported parameter ... : varargs

// Unsupported : g_task_return_pointer : unsupported parameter result_destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_task_run_in_thread : unsupported parameter task_func : no type generator for TaskThreadFunc, GTaskThreadFunc

// Unsupported : g_task_run_in_thread_sync : unsupported parameter task_func : no type generator for TaskThreadFunc, GTaskThreadFunc

// SetCheckCancellable is a wrapper around the C function g_task_set_check_cancellable.
func (recv *Task) SetCheckCancellable(checkCancellable bool) {
	c_check_cancellable :=
		boolToGboolean(checkCancellable)

	C.g_task_set_check_cancellable((*C.GTask)(recv.native), c_check_cancellable)

	return
}

// SetPriority is a wrapper around the C function g_task_set_priority.
func (recv *Task) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.g_task_set_priority((*C.GTask)(recv.native), c_priority)

	return
}

// SetReturnOnCancel is a wrapper around the C function g_task_set_return_on_cancel.
func (recv *Task) SetReturnOnCancel(returnOnCancel bool) bool {
	c_return_on_cancel :=
		boolToGboolean(returnOnCancel)

	retC := C.g_task_set_return_on_cancel((*C.GTask)(recv.native), c_return_on_cancel)
	retGo := retC == C.TRUE

	return retGo
}

// SetSourceTag is a wrapper around the C function g_task_set_source_tag.
func (recv *Task) SetSourceTag(sourceTag uintptr) {
	c_source_tag := (C.gpointer)(sourceTag)

	C.g_task_set_source_tag((*C.GTask)(recv.native), c_source_tag)

	return
}

// Unsupported : g_task_set_task_data : unsupported parameter task_data_destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

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
